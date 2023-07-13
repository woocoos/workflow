package deployment

import (
	"context"
	"errors"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/workflow/api/graphql/model"
	"github.com/woocoos/workflow/codegen/entgen/types"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/task"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/service/workflow"
	"os"
	"path/filepath"
)

var (
	ErrTaskHasClaim = errors.New("task has claim")
)

type Service struct {
	WFDB   *ent.Client
	Engine *workflow.Service
}

func (s *Service) createFileName(id, key, name string) string {
	fn := key + ":" + id + filepath.Ext(name)
	return filepath.Join(conf.Global().String("workflow.deploymentDir"), fn)
}

// CreateDeployment 创建部署
func (s *Service) CreateDeployment(ctx context.Context, input model.DeployDiagramInput) (dp *ent.Deployment, err error) {
	client := ent.FromContext(ctx)
	did := int(snowflake.New().Int64())
	pdBuiler := make([]*ent.ProcDefCreate, 0)
	ddBuiler := make([]*ent.DecisionDefCreate, 0)
	drdBuilder := make([]*ent.DecisionReqDefCreate, 0)
	// save deployment
	versionTag := ""
	for i, fid := range input.ResourceID {
		fb, err := os.ReadFile(filepath.Join(s.Engine.ResourceDir, input.ResourceKey[i]))
		if err != nil {
			return nil, err
		}
		eg, err := s.Engine.BuildEngine(input.ResourceKey[i], "", fb)
		if err != nil {
			return nil, err
		}

		switch eg := eg.(type) {
		case *engine.BpmnLoader:
			// 一般情况下，bpmn对应的流程图只有一个process
			for _, process := range eg.Spec.Process {
				if process.VersionTag != "" {
					versionTag = process.VersionTag
				}
				cr := client.ProcDef.Create().SetDeploymentID(did).SetAppID(input.AppID).
					SetKey(process.Id).SetName(process.Name).SetResourceID(fid).SetResourceKey(input.ResourceKey[i]).
					SetVersionTag(process.VersionTag)
				pdBuiler = append(pdBuiler, cr)
			}
		case *engine.DMNLoader:
			dreqid := int(snowflake.New().Int64())
			// 多个决策定义在一个文件中
			drdBuilder = append(drdBuilder, client.DecisionReqDef.Create().SetID(dreqid).
				SetDeploymentID(did).SetAppID(input.AppID).
				SetKey(eg.GetId()).SetName(eg.GetName()).SetResourceID(fid).SetResourceKey(input.ResourceKey[i]),
			)
			for _, decision := range eg.Spec.Decisions {
				cr := client.DecisionDef.Create().SetDeploymentID(did).SetReqDefID(dreqid).SetAppID(input.AppID).
					SetReqDefKey(eg.GetId()).
					SetKey(decision.Id).SetName(decision.Name).SetVersionTag(versionTag)
				ddBuiler = append(ddBuiler, cr)
			}

		}
	}
	if len(pdBuiler) != 0 {
		pds, err := client.ProcDef.CreateBulk(pdBuiler...).Save(ctx)
		if err != nil {
			return nil, err
		}
		// 以主流程的版本号为准
		for _, create := range ddBuiler {
			create.SetVersion(pds[0].Version)
		}
		for _, create := range drdBuilder {
			create.SetVersion(pds[0].Version)
		}
	}
	if len(ddBuiler) != 0 {
		if err = client.DecisionReqDef.CreateBulk(drdBuilder...).Exec(ctx); err != nil {
			return nil, err
		}
		if err = client.DecisionDef.CreateBulk(ddBuiler...).Exec(ctx); err != nil {
			return nil, err
		}
	}
	return client.Deployment.Create().SetID(did).SetAppID(input.AppID).
		SetName(input.Name).Save(ctx)
}

// StartProcessInstance 启动流程实例
func (s *Service) StartProcessInstance(ctx context.Context, input model.StartProcessInput) (*types.WorkflowRun, error) {
	wf, err := s.Engine.CreateAndRunProcessInstance(ctx, input)
	return wf, err
}

// ClaimTask 认领任务
func (s *Service) ClaimTask(ctx context.Context, taskID int) (*ent.Task, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	trow, err := client.Task.Query().Where(task.ID(taskID), task.TenantID(tid)).
		WithTaskIdentities(func(query *ent.IdentityLinkQuery) {
			query.Where(identitylink.LinkTypeEQ(identitylink.LinkTypeAssignee), identitylink.UserID(uid))
		}).Only(ctx)
	if err != nil {
		return nil, err
	}
	// 任务已经被认领
	if len(trow.Edges.TaskIdentities) != 0 {
		return nil, ErrTaskHasClaim
	}
	err = s.Engine.Client.SignalWorkflow(ctx, trow.ExecutionID, trow.RunID, engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
		TaskID:    trow.ID,
		LinkType:  identitylink.LinkTypeAssignee,
		UserID:    &uid,
		ProcDefID: trow.ProcDefID,
	})
	return trow, err
}
