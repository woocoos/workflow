package workflow

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/workflow/api/graphql/model"
	"github.com/woocoos/workflow/codegen/entgen/types"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/orgapp"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/pkg/api"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec"
	"github.com/woocoos/workflow/pkg/spec/vars"
	"go.temporal.io/sdk/client"
	"path/filepath"
	"strconv"
	"time"
)

type Service struct {
	Db           *ent.Client
	Client       client.Client
	TaskQueue    string
	WorkflowType string
	// 资源根目录
	ResourceDir string
}

func NewService(cnf *conf.AppConfiguration) (*Service, error) {
	co := client.Options{}
	err := cnf.Sub("temporal.clientOptions").Unmarshal(&co)
	if err != nil {
		return nil, fmt.Errorf("unable to build client option: %v", err)
	}
	c, err := client.Dial(co)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %v", err)
	}
	service := &Service{
		Client:      c,
		ResourceDir: cnf.String("workflow.resourceDir"),
		TaskQueue:   cnf.String("temporal.taskQueue"),
	}
	return service, nil
}

// BuildEngine 构建流程引擎
func (s *Service) BuildEngine(name string, path string, data []byte) (eg spec.Loader, err error) {
	switch filepath.Ext(name) {
	case ".bpmn", "bpm":
		eg = &engine.BpmnLoader{}
	case ".dmn":
		eg = &engine.DMNLoader{}
	default:
		return nil, fmt.Errorf("BuildEngine:not support file type")
	}

	if path != "" {
		if err = eg.LoadFromFile(path); err != nil {
			return nil, err
		}
		return
	}

	if len(data) == 0 {
		err = fmt.Errorf("BuildEngine:data is empty")
		return
	}
	if err = eg.Load(data); err != nil {
		return nil, err
	}
	return
}

func (s *Service) checkExists(ctx context.Context, input model.StartProcessInput, defID int) (*ent.ProcInst, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// 检查是否已经存在流程实例, TODO res ID是否存在可能需要应用端检查.
	pi, _ := s.Db.ProcInst.Query().Where(procinst.BusinessKey(input.BusinessKey),
		procinst.ProcDefID(defID), procinst.StatusIn(procinst.StatusReady, procinst.StatusActive),
	).Order(ent.Desc(procinst.FieldID)).First(ctx)
	if pi != nil {
		return pi, nil
	}
	// 检查应用是否在该组织下.
	if _, err := s.Db.OrgApp.Query().Where(
		orgapp.OrgID(tid), orgapp.AppID(pi.AppID)).Only(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

// CreateAndRunProcessInstance 启动流程实例,如果实例已经在执行过程则返回工作流ID,否则将创新流程.对于调用方需要控制状态.
func (s *Service) CreateAndRunProcessInstance(ctx context.Context, input model.StartProcessInput) (wfr *types.WorkflowRun, err error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	pd, err := s.Db.ProcDef.Query().Where(procdef.Key(input.ProcDefKey), procdef.TenantID(tid)).
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
		uid, err := strconv.Atoi(up.Name())
		if err != nil {
			return nil, err
		}
		pi, err = s.Db.ProcInst.Create().SetTenantID(pd.TenantID).SetProcDefID(pd.ID).
			SetAppID(pd.AppID).SetBusinessKey(input.BusinessKey).SetProcDefID(pd.ID).SetStartTime(time.Now()).
			SetStartUserID(uid).SetStatus(procinst.StatusReady).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	inst := api.InstanceRequest{
		ProcDefKey:  pd.Key,
		ProcInst:    pi,
		ResourceKey: pd.ResourceKey,
		Variables:   make(vars.Mapping),
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
