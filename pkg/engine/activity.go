package engine

import (
	"context"
	"fmt"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"strconv"
)

type BpmnActivity struct {
	Engine *BPMN
}

func (a *BpmnActivity) ServiceTaskActivity(ctx context.Context, req ServiceTaskActivityRequest) (bpmn.Mappings, error) {
	typename := ""
	if req.Element.TaskDefinition != nil {
		typename = req.Element.TaskDefinition.TypeName
	}
	output := make(bpmn.Mappings)
	hout, err := a.Engine.Handlers.RunHandler(ctx, typename, req.InstanceRequest.Variables)
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

// CreateUserTaskActivity 创建用户任务,将任务信息返回到工作流引擎中.
func (a *BpmnActivity) CreateUserTaskActivity(ctx context.Context, req UserTaskActivityRequest) (*ent.Task, error) {
	return a.Engine.Exporter.CreateUserTask(ctx, req.InstanceRequest, req.Element)
}

func (a *BpmnActivity) BusinessRuleTaskActivity(ctx context.Context, req BusinessRuleTaskActivityRequest) (bpmn.Mappings, error) {
	//logger := workflow.GetLogger(ctx)
	if req.Element.CalledDecision == nil {
		return nil, fmt.Errorf("called decision is nil: %s", req.Element.Id)
	}

	cdf, err := a.Engine.Exporter.GetDecisionReqDef(context.Background(), &GetDecisionReqDefRequest{
		DecisionDefKey: req.Element.CalledDecision.DecisionId,
		OrgID:          req.InstanceRequest.TenantID,
	})
	loader, err := loadDMN(a.Engine.ResourceDir, cdf)
	if err != nil {
		return nil, err
	}
	dmn := DMN{
		Id: strconv.Itoa(cdf.ID),
	}
	ret, err := dmn.StartInstance(loader, req.Element, req.InstanceRequest)
	return ret, err
}
