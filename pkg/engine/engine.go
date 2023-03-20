package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"strconv"
	"time"
)

type (
	// BPMN is a workflow engine for bpmn
	BPMN struct {
		// worker id
		Id       string
		Exporter Exporter
		Handlers *HandlerFuncMap
	}
)

func NewBPMN(id string, exporter Exporter) *BPMN {
	return &BPMN{
		Id:       id,
		Exporter: exporter,
		Handlers: NewHandlerFuncMap(id),
	}
}

func DefaultActivityOption() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
	}
}

func (bp *BPMN) StartInstance(loader *BpmnLoader, pi *InstanceRequest) (*Iterator, error) {
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

func (bp *BPMN) Start(ctx workflow.Context, pi *InstanceRequest, loader *BpmnLoader) (err error) {
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

func (bp *BPMN) handleEventBasedGateway(ctx workflow.Context, it *Iterator, pi *InstanceRequest, subs []bpmn.Elementor) (next bpmn.Elementor, err error) {
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

func (bp *BPMN) handleIntermediateCatchEvent(ctx workflow.Context, it *Iterator, pi *InstanceRequest, ele *bpmn.IntermediateCatchEvent) (err error) {
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

func (bp *BPMN) handleServiceTask(ctx workflow.Context, pi *InstanceRequest, taskEle *bpmn.ServiceTask) error {
	ao := DefaultActivityOption()
	_ = taskEle.ExtensionProperties.Decode(&ao) // ignore error
	if taskEle.TaskDefinition != nil {
		v, err := taskEle.TaskDefinition.GetRetries()
		if err != nil {
			return err
		}
		ao.RetryPolicy = &temporal.RetryPolicy{
			MaximumAttempts: int32(v),
		}
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var variables bpmn.Mappings
	err := workflow.ExecuteActivity(ctx, bp.ServiceTaskActivity, ServiceTaskActivityRequest{
		InstanceRequest: pi,
		Element:         taskEle,
	}).Get(ctx, &variables)
	if err != nil {
		return err
	}
	for k, v := range variables {
		pi.Variables[k] = v
	}
	return nil
}

func (bp *BPMN) handleUserTask(ctx workflow.Context, pi *InstanceRequest, taskEle *bpmn.UserTask) error {
	logger := workflow.GetLogger(ctx)
	if taskEle.Assignment == nil {
		logger.Warn("user task assignment is nil", "task", taskEle.Id)
		return nil
	}
	ao := DefaultActivityOption()
	_ = taskEle.ExtensionProperties.Decode(&ao) // ignore error
	ctx = workflow.WithActivityOptions(ctx, ao)

	var task ent.Task
	err := workflow.ExecuteActivity(ctx, bp.CreateUserTaskActivity, UserTaskActivityRequest{
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

func (bp *BPMN) ServiceTaskActivity(ctx context.Context, req ServiceTaskActivityRequest) (bpmn.Mappings, error) {
	typename := ""
	if req.Element.TaskDefinition != nil {
		typename = req.Element.TaskDefinition.TypeName
	}
	output := make(bpmn.Mappings)
	hout, err := bp.Handlers.RunHandler(ctx, typename, req.InstanceRequest.Variables)
	if err != nil {
		return nil, err
	}
	for k, v := range hout {
		req.InstanceRequest.Variables[k] = v
	}
	if len(req.Element.Outputs) > 0 {
		for _, ele := range req.Element.Outputs {
			err = ele.OutputTarget(req.InstanceRequest.Variables, output)
			if err != nil {
				return nil, err
			}
		}
	} else {
		output = hout
	}
	return output, nil
}

func (bp *BPMN) CreateUserTaskActivity(ctx context.Context, req UserTaskActivityRequest) (*ent.Task, error) {
	return bp.Exporter.CreateUserTask(ctx, req.InstanceRequest, req.Element)
}

func (bp *BPMN) handleCallActivity(ctx workflow.Context, pi *InstanceRequest, el *bpmn.CallActivity) error {
	logger := workflow.GetLogger(ctx)
	cdf, err := bp.Exporter.GetProcDef(context.Background(), &GetProcDefRequest{
		ProcDefKey: el.CalledElement.ProcessId,
		OrgID:      pi.OrgID,
	})
	cir := InstanceRequest{
		ProcInst:     pi.ProcInst,
		ProcDefKey:   cdf.Key,
		ResourceName: cdf.ResourceName,
		ResourceData: cdf.ResourceData,
	}
	if el.CalledElement.PropagateAllChildVariables {
		cir.Variables = pi.Variables
	} else {
		cir.Variables = make(bpmn.Mappings)
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

	var result bpmn.Mappings
	err = workflow.ExecuteChildWorkflow(ctx, workflow.GetInfo(ctx).WorkflowType.Name, cir).Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return err
	}
	return err
}

func loadDMN(cdf *ent.DecisionReqDef) (*DMNLoader, error) {
	dmn := NewDMNLoader(strconv.Itoa(cdf.ID))
	if len(cdf.ResourceData) != 0 {
		err := dmn.Load(cdf.ResourceData)
		if err != nil {
			return nil, err
		}
	}
	// TODO load from other ,cache
	return dmn, nil
}

func (bp *BPMN) handleBusinessRuleTask(ctx workflow.Context, pi *InstanceRequest, el *bpmn.BusinessRuleTask) error {

	ao := DefaultActivityOption()
	_ = el.ExtensionProperties.Decode(&ao) // ignore error
	ctx = workflow.WithActivityOptions(ctx, ao)

	var variables bpmn.Mappings
	err := workflow.ExecuteActivity(ctx, bp.BusinessRuleTaskActivity, BusinessRuleTaskActivityRequest{
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

func (bp *BPMN) BusinessRuleTaskActivity(ctx context.Context, req BusinessRuleTaskActivityRequest) (bpmn.Mappings, error) {
	//logger := workflow.GetLogger(ctx)
	if req.Element.CalledDecision == nil {
		return nil, fmt.Errorf("called decision is nil: %s", req.Element.Id)
	}

	cdf, err := bp.Exporter.GetDecisionReqDef(context.Background(), &GetDecisionReqDefRequest{
		DecisionDefKey: req.Element.CalledDecision.DecisionId,
		OrgID:          req.InstanceRequest.OrgID,
	})
	loader, err := loadDMN(cdf)
	if err != nil {
		return nil, err
	}
	dmn := DMN{
		Id: strconv.Itoa(cdf.ID),
	}
	ret, err := dmn.StartInstance(loader, req.Element, req.InstanceRequest)
	return ret, err
}
