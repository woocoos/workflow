package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/api"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/pkg/spec/vars"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type (
	Option func(*BPMN)

	// BPMN is a workflow engine for bpmn
	BPMN struct {
		// worker id
		Id          string
		Exporter    api.Exporter
		Handlers    *HandlerFuncMap
		Activity    *BpmnActivity
		ResourceDir string
	}
)

// WithResourceDir 指定资源目录,在工作流中必须指定
func WithResourceDir(dir string) Option {
	return func(bp *BPMN) {
		bp.ResourceDir = dir
	}
}

// WithExporter 指定外部相关接口实现.
func WithExporter(exporter api.Exporter) Option {
	return func(bp *BPMN) {
		bp.Exporter = exporter
	}
}

func WithID(id string) Option {
	return func(bp *BPMN) {
		bp.Id = id
	}
}

func NewBPMN(opts ...Option) *BPMN {
	bp := &BPMN{}
	bp.Activity = &BpmnActivity{
		Engine: bp,
	}
	for _, opt := range opts {
		opt(bp)
	}
	bp.Handlers = NewHandlerFuncMap(bp.Id)
	if bp.Exporter == nil {

	}
	return bp
}

func DefaultActivityOption() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
	}
}

func (bp *BPMN) StartInstance(loader *BpmnLoader, pi *api.InstanceRequest) (*Iterator, error) {
	pd := loader.FindProcess(pi.ProcDefKey)
	if pd == nil {
		return nil, fmt.Errorf("[RunInstance]process %s not found", pi.ProcDefKey)
	}
	it := &Iterator{
		loader:  loader,
		process: pd,
		Request: pi,
	}
	for _, event := range it.process.StartEvents {
		it.queue = append(it.queue, event)
	}
	return it, nil
}

func (bp *BPMN) Start(ctx workflow.Context, pi *api.InstanceRequest, loader *BpmnLoader) (err error) {
	iterator, err := bp.StartInstance(loader, pi)
	if err != nil {
		return
	}
	iterator.CurrentHandler = func(it *Iterator, ele bpmn.Elementor) (err error) {
		switch et := ele.(type) {
		case *bpmn.StartEvent:
			// TODO message hanndler
		case *bpmn.UserTask:
			err = bp.handleUserTask(ctx, pi, et)
		case *bpmn.BusinessRuleTask:
			err = bp.handleBusinessRuleTask(ctx, pi, et)
		case *bpmn.ServiceTask:
			err = bp.handleServiceTask(ctx, pi, et)
		case *bpmn.EventBasedGateway:
		case *bpmn.IntermediateCatchEvent:
			err = bp.handleIntermediateCatchEvent(ctx, it, pi, et)
		case *bpmn.CallActivity:
			err = bp.handleCallActivity(ctx, pi, et)
		case *bpmn.EndEvent:
		default:
			// do nothing
		}
		return
	}
	iterator.QueueHandler = func(it *Iterator, subs []bpmn.Elementor) (next bpmn.Elementor, err error) {
		switch it.current.(type) {
		case *bpmn.EventBasedGateway:
			return bp.handleEventBasedGateway(ctx, it, pi, subs)
		}
		return nil, fmt.Errorf("not support queue handler: %T", it.current)
	}
	for iterator.Next() {
		err = iterator.Err()
		if err != nil {
			return
		}
	}
	err = iterator.Err()
	return err
}

func (bp *BPMN) handleEventBasedGateway(ctx workflow.Context, it *Iterator, pi *api.InstanceRequest, subs []bpmn.Elementor) (next bpmn.Elementor, err error) {
	//logger := workflow.GetLogger(ctx)
	evl := EventListen{
		loader: it.Loader(),
		eles:   subs,
		input:  pi.Variables,
	}
	workflow.Go(ctx, evl.MultiEventListen)

	err = workflow.Await(ctx, func() bool {
		if evl.err != nil {
			err = evl.err
			return true
		}
		return evl.isComplete
	})
	next = evl.next
	for k, v := range evl.output {
		pi.Variables[k] = v
	}
	return
}

func (bp *BPMN) handleIntermediateCatchEvent(ctx workflow.Context, it *Iterator, pi *api.InstanceRequest, ele *bpmn.IntermediateCatchEvent) (err error) {
	evl := EventListen{
		loader: it.Loader(),
		eles:   []bpmn.Elementor{ele},
		input:  pi.Variables,
	}
	workflow.Go(ctx, evl.MultiEventListen)
	err = workflow.Await(ctx, func() bool {
		if evl.err != nil {
			err = evl.err
			return true
		}
		return evl.isComplete
	})
	for k, v := range evl.output {
		pi.Variables[k] = v
	}
	return
}

func (bp *BPMN) handleServiceTask(ctx workflow.Context, pi *api.InstanceRequest, taskEle *bpmn.ServiceTask) error {
	if taskEle.TaskDefinition == nil {
		return fmt.Errorf("service task definition not found: %s", taskEle.Id)
	}
	typename, err := taskEle.TaskDefinition.GetType()
	if err != nil {
		return err
	}

	ao := DefaultActivityOption()
	_ = taskEle.ExtensionProperties.Decode(&ao) // ignore error

	v, err := taskEle.TaskDefinition.GetRetries()
	if err != nil {
		return err
	}
	if v > 0 {
		ao.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(v),
		}
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var variables vars.Mapping
	switch typename {
	case "http":
		err = workflow.ExecuteActivity(ctx, bp.Activity.HttpServiceTaskActivity, api.ServiceTaskActivityRequest{
			InstanceRequest: pi,
			Element:         taskEle,
		}).Get(ctx, &variables)
	default:
		err = workflow.ExecuteActivity(ctx, bp.Activity.ServiceTaskActivity, api.ServiceTaskActivityRequest{
			InstanceRequest: pi,
			Element:         taskEle,
		}).Get(ctx, &variables)
	}
	if err != nil {
		return err
	}
	for k, v := range variables {
		pi.Variables[k] = v
	}
	return nil
}

func (bp *BPMN) handleUserTask(ctx workflow.Context, pi *api.InstanceRequest, taskEle *bpmn.UserTask) error {
	logger := workflow.GetLogger(ctx)
	if taskEle.Assignment == nil {
		logger.Warn("user task assignment is nil", "task", taskEle.Id)
		return nil
	}
	ao := DefaultActivityOption()
	_ = taskEle.ExtensionProperties.Decode(&ao) // ignore error
	ctx = workflow.WithActivityOptions(ctx, ao)

	var task ent.Task
	err := workflow.ExecuteActivity(ctx, bp.Activity.CreateUserTaskActivity, api.UserTaskActivityRequest{
		InstanceRequest: pi,
		Element:         taskEle,
	}).Get(ctx, &task)
	if err != nil {
		return err
	}

	ul := UserTaskListen{
		task:     &task,
		exporter: bp.Exporter,
	}
	workflow.Go(ctx, ul.Listen)

	var iscomplete bool
	awErr := workflow.Await(ctx, func() bool {
		iscomplete, err = ul.CheckComplete(context.Background(), taskEle)
		if err != nil {
			return true
		}
		return iscomplete
	})
	if awErr != nil || err != nil {
		return errors.Join(awErr, err)
	}
	return nil
}

func (bp *BPMN) handleCallActivity(ctx workflow.Context, pi *api.InstanceRequest, el *bpmn.CallActivity) error {
	logger := workflow.GetLogger(ctx)
	cdf, err := bp.Exporter.GetProcDef(context.Background(), &api.GetProcDefRequest{
		ProcDefKey: el.CalledElement.ProcessId,
		TenantID:   pi.TenantID,
	})
	cir := api.InstanceRequest{
		ProcInst:    pi.ProcInst,
		ProcDefKey:  cdf.Key,
		ResourceKey: cdf.ResourceKey,
	}
	if el.CalledElement.PropagateAllChildVariables {
		cir.Variables = pi.Variables
	} else {
		cir.Variables = make(vars.Mapping)
		for _, ot := range el.Outputs {
			err = ot.OutputTarget(pi.Variables, cir.Variables)
			if err != nil {
				return err
			}
		}
	}

	cwo := workflow.ChildWorkflowOptions{
		ParentClosePolicy: enums.PARENT_CLOSE_POLICY_ABANDON,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var result vars.Mapping
	err = workflow.ExecuteChildWorkflow(ctx, workflow.GetInfo(ctx).WorkflowType.Name, cir).Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return err
	}
	return err
}

func loadDMN(resourceDir string, cdf *ent.DecisionReqDef) (*DMNLoader, error) {
	dmn := NewDMNLoader(strconv.Itoa(cdf.ID))
	bs, err := os.ReadFile(filepath.Join(resourceDir, cdf.ResourceKey))
	if err != nil {
		return nil, err
	}
	if err := dmn.Load(bs); err != nil {
		return nil, err
	}
	// TODO load from other ,cache
	return dmn, nil
}

func (bp *BPMN) handleBusinessRuleTask(ctx workflow.Context, pi *api.InstanceRequest, el *bpmn.BusinessRuleTask) error {

	ao := DefaultActivityOption()
	_ = el.ExtensionProperties.Decode(&ao) // ignore error
	ctx = workflow.WithActivityOptions(ctx, ao)

	var variables vars.Mapping
	err := workflow.ExecuteActivity(ctx, bp.Activity.BusinessRuleTaskActivity, api.BusinessRuleTaskActivityRequest{
		InstanceRequest: pi,
		Element:         el,
	}).Get(ctx, &variables)
	if err != nil {
		return err
	}
	for k, v := range variables {
		pi.Variables[k] = v
	}
	return nil
}
