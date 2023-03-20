package workflow

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/security"
	portalent "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/graph/entgen/types"
	"github.com/woocoos/workflow/graph/model"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/client"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	WFDB         *ent.Client
	PortalDB     *portalent.Client
	Client       client.Client
	TaskQueue    string
	WorkflowType string
}

func (s *Service) BuildEngine(name string, path string, data []byte) (eg spec.Loader, err error) {
	switch filepath.Ext(name) {
	case ".bpmn", "bpm":
		eg = &engine.BpmnLoader{}
	case ".dmn":
		eg = &engine.DMNLoader{}
	default:
		return nil, fmt.Errorf("[BuildEngine]not support file type")
	}

	if path != "" {
		if err = eg.LoadFromFile(path); err != nil {
			return nil, err
		}
		return
	}

	if len(data) == 0 {
		return
	}
	if err = eg.Load(data); err != nil {
		return nil, err
	}
	return
}

func (s *Service) checkExists(ctx context.Context, input model.StartProcessInput, defID int) (*ent.ProcInst, error) {
	// 检查是否已经存在流程实例, TODO res ID是否存在可能需要应用端检查.
	pi, _ := s.WFDB.ProcInst.Query().Where(procinst.BusinessKey(input.ResID),
		procinst.ProcDefID(defID), procinst.StatusIn(procinst.StatusReady, procinst.StatusActive),
	).Order(ent.Desc(procinst.FieldID)).First(ctx)
	if pi != nil {
		return pi, nil
	}
	// 检查应用是否在该组织下.
	res := strings.Split(input.ResID, ":")
	appCode := res[0]
	if _, err := s.PortalDB.OrgApp.Query().WithApp(func(query *portalent.AppQuery) {
		query.Where(app.Code(appCode))
	}).Where(orgapp.OrgID(input.OrgID)).Only(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

// CreateAndRunProcessInstance 启动流程实例,如果实例已经在执行过程则返回工作流ID,否则将创新流程.对于调用方需要控制状态.
func (s *Service) CreateAndRunProcessInstance(ctx context.Context, input model.StartProcessInput) (wfr *types.WorkflowRun, err error) {
	pd, err := s.WFDB.ProcDef.Query().Where(procdef.Key(input.ProcDefKey), procdef.OrgID(input.OrgID)).
		Order(ent.Desc(procdef.FieldVersion, procdef.FieldRevision)).First(ctx)
	if err != nil {
		return nil, err
	}
	pi, err := s.checkExists(ctx, input, pd.ID)
	if err != nil {
		return nil, err
	}
	if pi == nil {
		up := security.GenericIdentityFromContext(ctx)
		pi, err = s.WFDB.ProcInst.Create().SetOrgID(pd.OrgID).SetProcDefID(pd.ID).
			SetAppID(pd.AppID).SetBusinessKey(input.ResID).SetProcDefID(pd.ID).SetStartTime(time.Now()).
			SetStartUserID(up.NameIntX()).SetStatus(procinst.StatusReady).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	inst := engine.InstanceRequest{
		ProcDefKey:   pd.Key,
		ProcInst:     pi,
		ResourceName: pd.ResourceName,
		Variables:    make(bpmn.Mappings),
	}
	for _, variable := range input.Variables {
		inst.Variables[variable.Name] = variable.Value
	}
	wf, err := s.Client.ExecuteWorkflow(ctx,
		client.StartWorkflowOptions{
			ID:        inst.ProcDefKey + strconv.Itoa(pi.ID),
			TaskQueue: s.TaskQueue,
		},
		s.WorkflowType, inst)
	if err != nil {
		return nil, err
	}
	wfr = &types.WorkflowRun{
		ID:    wf.GetID(),
		RunID: wf.GetRunID(),
	}
	return
}
