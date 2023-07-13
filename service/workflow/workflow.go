package workflow

import (
	"context"
	"fmt"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/pkg/engine"
	"go.temporal.io/sdk/workflow"
	"path/filepath"
	"strconv"
)

// Definition temporal workflow definition.
type Definition struct {
	engine.Exporter
	Service *Service
}

// BPMNWorkflowDef is entry point of workflow.
func (def *Definition) BPMNWorkflowDef(ctx workflow.Context, pi engine.InstanceRequest) error {
	if pi.ResourceKey == "" {
		dr, err := def.Service.Db.ProcDef.Query().Where(procdef.ID(pi.ProcDefID)).
			Only(schemax.SkipTenantPrivacy(context.Background()))
		if err != nil {
			return err
		}
		pi.ResourceKey = dr.ResourceKey
	}
	pi.ResourcePath = filepath.Join(def.Service.ResourceDir, pi.ResourceKey)
	eg, err := def.Service.BuildEngine(pi.ResourceKey, pi.ResourcePath, nil)
	if err != nil {
		return err
	}

	egn, ok := eg.(*engine.BpmnLoader)
	if !ok {
		err = fmt.Errorf("not support engine type")
		return err
	}
	id := strconv.Itoa(pi.TenantID) + ":" + pi.ProcDefKey
	eng := engine.NewBPMN(engine.WithID(id), engine.WithExporter(def.Exporter),
		engine.WithResourceDir(def.Service.ResourceDir))
	return eng.Start(ctx, &pi, egn)
}
