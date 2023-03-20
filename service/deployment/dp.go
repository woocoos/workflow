package deployment

import (
	"context"
	"errors"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/pkg/snowflake"
	portal "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/security"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/task"
	"github.com/woocoos/workflow/graph/entgen/types"
	"github.com/woocoos/workflow/graph/model"
	"github.com/woocoos/workflow/identity"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/service/util"
	"github.com/woocoos/workflow/service/workflow"
	"io"
	"path/filepath"
	"strconv"
)

var (
	ErrTaskHasClaim = errors.New("task has claim")
)

type Service struct {
	WFDB     *ent.Client
	PortalDB *portal.Client
	Engine   workflow.Service
}

func (s *Service) createFileName(id, key, name string) string {
	fn := key + ":" + id + filepath.Ext(name)
	return filepath.Join(conf.Global().String("workflow.deploymentDir"), fn)
}

// CreateDeployment 创建部署
func (s *Service) CreateDeployment(ctx context.Context, input model.DeployDiagramInput) (dp *ent.Deployment, err error) {
	client := ent.FromContext(ctx)
	did := snowflake.New()
	pdBuiler := make([]*ent.ProcDefCreate, 0)
	ddBuiler := make([]*ent.DecisionDefCreate, 0)
	drdBuilder := make([]*ent.DecisionReqDefCreate, 0)
	// save deployment
	versionTag := ""
	for _, file := range input.Files {
		fb, err := io.ReadAll(file.File.File)
		if err != nil {
			return nil, err
		}
		eg, err := s.Engine.BuildEngine(file.File.Filename, "", fb)
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
				cr := client.ProcDef.Create().SetDeploymentID(int(did.Int64())).SetOrgID(input.OrgID).SetAppID(input.AppID).
					SetKey(process.Id).SetResourceName(file.File.Filename).SetVersionTag(process.VersionTag).
					SetResourceData(fb)
				pdBuiler = append(pdBuiler, cr)
			}
		case *engine.DMNLoader:
			dreqid := snowflake.New()
			// 多个决策定义在一个文件中
			drdBuilder = append(drdBuilder, client.DecisionReqDef.Create().SetID(int(dreqid.Int64())).
				SetDeploymentID(int(did.Int64())).SetOrgID(input.OrgID).SetAppID(input.AppID).
				SetKey(eg.GetId()).SetName(eg.GetName()).SetResourceData(fb),
			)
			for _, decision := range eg.Spec.Decisions {
				cr := client.DecisionDef.Create().SetDeploymentID(int(did.Int64())).SetReqDefID(int(dreqid.Int64())).
					SetOrgID(input.OrgID).SetAppID(input.AppID).SetReqDefKey(eg.GetId()).
					SetKey(decision.Id).SetResourceName(file.File.Filename).SetVersionTag(versionTag)
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

		if conf.Global().IsSet("workflow.deploymentDir") {
			for _, pd := range pds {
				err = util.SaveResource(strconv.Itoa(pd.ID), pd.Key, pd.ResourceName, pd.ResourceData)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	if len(ddBuiler) != 0 {
		dds, err := client.DecisionReqDef.CreateBulk(drdBuilder...).Save(ctx)
		if err != nil {
			return nil, err
		}
		_, err = client.DecisionDef.CreateBulk(ddBuiler...).Save(ctx)
		if err != nil {
			return nil, err
		}
		if conf.Global().IsSet("workflow.deploymentDir") {
			for _, dd := range dds {
				err = util.SaveResource(strconv.Itoa(dd.ID), dd.Key, dd.ResourceName, dd.ResourceData)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return client.Deployment.Create().SetID(int(did.Int64())).SetOrgID(input.OrgID).SetAppID(input.AppID).
		SetName(input.Name).Save(ctx)
}

// StartProcessInstance 启动流程实例
func (s *Service) StartProcessInstance(ctx context.Context, input model.StartProcessInput) (*types.WorkflowRun, error) {
	wf, err := s.Engine.CreateAndRunProcessInstance(ctx, input)
	return wf, err
}

// ClaimTask 认领任务
func (s *Service) ClaimTask(ctx context.Context, taskID int) (*ent.Task, error) {
	uid := security.UserIDFromContext(ctx)
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromWebContext(ctx)
	if err != nil {
		return nil, err
	}
	trow, err := client.Task.Query().Where(task.ID(taskID), task.OrgID(tid)).
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
