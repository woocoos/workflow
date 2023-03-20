package workflow

import (
	"context"
	"fmt"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/pkg/conv"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/workflow"
	"strconv"
)

type Definition struct {
	engine.Exporter
	Service   *Service
	TaskQueue string
}

// BPMNWorkflowDef temporal workflow definition
func (def *Definition) BPMNWorkflowDef(ctx workflow.Context, pi engine.InstanceRequest) (err error) {
	rd := pi.ResourceData
	if len(pi.ResourceData) == 0 && pi.ResourcePath == "" {
		def, err := def.Service.WFDB.ProcDef.Query().Where(procdef.ID(pi.ProcDefID)).Only(context.Background())
		if err != nil {
			return err
		}
		rd = def.ResourceData
	}
	eg, err := def.Service.BuildEngine(pi.ResourceName, pi.ResourcePath, rd)
	if err != nil {
		return
	}

	egn, ok := eg.(*engine.BpmnLoader)
	if !ok {
		err = fmt.Errorf("not support engine type")
		return
	}
	id := strconv.Itoa(pi.OrgID) + ":" + pi.ProcDefKey
	bpmn.Convert = conv.NewDefaultConverter()
	eng := &engine.BPMN{
		Id:       id,
		Exporter: def.Exporter,
		Handlers: engine.NewHandlerFuncMap(id),
	}
	return eng.Start(ctx, &pi, egn)
}
