package workflow

import (
	"context"
	"errors"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/orgrole"
	"github.com/woocoos/workflow/ent/orgroleuser"
	"github.com/woocoos/workflow/ent/orguser"
	"github.com/woocoos/workflow/ent/procdef"
	entask "github.com/woocoos/workflow/ent/task"
	"github.com/woocoos/workflow/pkg/api"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/activity"
	"go.uber.org/zap"
)

var (
	_ api.Exporter = (*DbExport)(nil)

	ErrTenantIDRequire  = errors.New("tenant id is required")
	ErrUserIDRequire    = errors.New("user id is required")
	ErrTaskIDRequire    = errors.New("task id is required")
	ErrIdentityNotFound = errors.New("identity link not found")
)

type DbExport struct {
	Service *Service
}

func AlternateSchema() ent.Option {
	return ent.AlternateSchema(ent.SchemaConfig{
		OrgUser:     "portal",
		OrgRole:     "portal",
		OrgApp:      "portal",
		OrgRoleUser: "portal",
	})
}

func (e *DbExport) GetProcDef(ctx context.Context, req *api.GetProcDefRequest) (*ent.ProcDef, error) {
	return e.Service.Db.ProcDef.Query().Where(procdef.Key(req.ProcDefKey), procdef.TenantID(req.TenantID)).First(ctx)
}

func (e *DbExport) GetDecisionReqDef(ctx context.Context, req *api.GetDecisionReqDefRequest) (*ent.DecisionReqDef, error) {
	def, err := e.Service.Db.DecisionDef.Query().Where(decisiondef.Key(req.DecisionDefKey), decisiondef.TenantID(req.OrgID)).
		Order(ent.Desc(decisiondef.FieldVersion)).WithReqDef().First(ctx)
	if err != nil {
		return nil, err
	}
	return def.Edges.ReqDefOrErr()
}

// GetUserIDs 根据用户名称获取用户ID
func (e *DbExport) GetUserIDs(ctx context.Context, req *api.GetUserIDsRequest) (*api.GetUserIDResponse, error) {
	ids, err := e.Service.Db.OrgUser.Query().Where(orguser.OrgID(req.OrgID),
		orguser.DisplayNameIn(req.Names...)).Select(orguser.FieldUserID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	return &api.GetUserIDResponse{UserIDs: ids}, nil
}

func (e *DbExport) GetGroupIDs(ctx context.Context, req *api.GetGroupIDsRequest) (*api.GetGroupIDResponse, error) {
	ids, err := e.Service.Db.OrgRole.Query().Where(orgrole.OrgID(req.OrgID),
		orgrole.NameIn(req.Names...), orgrole.KindEQ(orgrole.KindGroup)).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return &api.GetGroupIDResponse{GroupIDs: ids}, nil
}

func (e *DbExport) GetGroupUsers(ctx context.Context, req *api.GetGroupUserRequest) (*api.GetGroupUserResponse, error) {
	ids, err := e.Service.Db.OrgRoleUser.Query().Where(orgroleuser.OrgID(req.TenantID),
		orgroleuser.HasOrgRoleWith(orgrole.NameEQ(req.Name))).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return &api.GetGroupUserResponse{UserIDs: ids}, nil
}

func (e *DbExport) getAssignee(ctx context.Context, pi *api.InstanceRequest, taskEle *bpmn.UserTask) (
	name string, id int, err error) {
	if taskEle.Assignment.Assignee == "" {
		return
	}
	name, err = taskEle.Assignment.GetAssignee(pi.Variables)
	if err != nil {
		return
	}
	resp, err := e.GetUserIDs(ctx, &api.GetUserIDsRequest{Names: []string{name}, OrgID: pi.TenantID})
	if err != nil {
		return
	}
	if len(resp.UserIDs) != 1 {
		err = ErrUserIDRequire
		return
	}
	id = resp.UserIDs[0]
	return
}

func (e *DbExport) getCandidateGroups(ctx context.Context, pi *api.InstanceRequest, taskEle *bpmn.UserTask) (
	ids []int, err error) {
	logger := activity.GetLogger(ctx)
	if taskEle.Assignment.CandidateGroups == "" {
		return
	}
	cgs, err := taskEle.Assignment.GetCandidateGroups(pi.Variables)
	if err != nil {
		return nil, err
	}
	resp, err := e.GetGroupIDs(ctx, &api.GetGroupIDsRequest{Names: cgs, OrgID: pi.TenantID})
	if err != nil {
		return nil, err
	}
	if len(resp.GroupIDs) != len(cgs) {
		logger.Error("candidate group not found", zap.Strings("groups", cgs),
			zap.Ints("found", resp.GroupIDs))
	}
	ids = resp.GroupIDs
	return
}

func (e *DbExport) getCandidateUsers(ctx context.Context, pi *api.InstanceRequest, taskEle *bpmn.UserTask) (
	ids []int, err error) {
	logger := activity.GetLogger(ctx)
	if taskEle.Assignment.CandidateUsers == "" {
		return
	}
	cus, err := taskEle.Assignment.GetCandidateUsers(pi.Variables)
	if err != nil {
		return nil, err
	}
	resp, err := e.GetUserIDs(ctx, &api.GetUserIDsRequest{Names: cus, OrgID: pi.TenantID})
	if err != nil {
		return nil, err
	}
	if len(resp.UserIDs) != len(cus) {
		logger.Error("candidate group not found", zap.Strings("users", cus),
			zap.Ints("found", resp.UserIDs))
	}
	ids = resp.UserIDs
	return
}

// CreateUserTask 创建用户任务,并更新任务统计后,将任务数据返回给工作流引擎.
//
// 任务统计:
//
//	1.如果任务已经明确指了Assignee,则将成员数+1;
//	2.如果只设定了候选组,则成员数+1;
//	3.如果是指定了候选人且任务为多实例,则成员数及未完成数初始化总数;
func (e *DbExport) CreateUserTask(ctx context.Context, pi *api.InstanceRequest, taskEle *bpmn.UserTask) (tk *ent.Task, err error) {
	exec := activity.GetInfo(ctx).WorkflowExecution
	if taskEle.Assignment == nil {
		return nil, nil
	}
	ctx = schemax.SkipTenantPrivacy(ctx)

	assignee, assigneeId, err := e.getAssignee(ctx, pi, taskEle)
	if err != nil {
		return nil, err
	}

	candidateGroups, err := e.getCandidateGroups(ctx, pi, taskEle)
	if err != nil {
		return nil, err
	}

	candidateUsers, err := e.getCandidateUsers(ctx, pi, taskEle)
	if err != nil {
		return nil, err
	}

	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.Db.Tx(ctx)
	}, func(vtx ecx.Transactor) error {
		tx := vtx.(*ent.Tx)
		cnt := int32(0)
		taskbuilder := tx.Task.Create().SetProcInstID(pi.ID).SetTenantID(pi.TenantID).SetProcDefID(pi.ProcDefID).
			SetTaskDefKey(taskEle.Id).SetComments(taskEle.GetDocumentation()).
			SetExecutionID(exec.ID).SetRunID(exec.RunID).SetStatus(entask.StatusRunning)
		if assignee != "" {
			taskbuilder.SetAssignee(assignee)
			cnt++
		} else if len(candidateGroups) > 0 {
			taskbuilder.SetKind(entask.KindOR)
			cnt++
		} else if taskEle.MultiInstance != nil {
			taskbuilder.SetSequential(taskEle.MultiInstance.IsSequential)
			taskbuilder.SetKind(entask.KindAND)
			cnt += int32(len(candidateUsers))
		}
		taskbuilder.SetMemberCount(cnt).SetUnfinishedCount(cnt)
		tk, err = taskbuilder.Save(ctx)
		if err != nil {
			return err
		}

		// identity link
		idbuilders := make([]*ent.IdentityLinkCreate, 0)
		if assignee != "" {
			idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetTenantID(pi.TenantID).
				SetLinkType(identitylink.LinkTypeAssignee).SetOperationType(identitylink.OperationTypeInit).
				SetUserID(assigneeId)
			idbuilders = append(idbuilders, idb)
		}
		for _, gid := range candidateGroups {
			idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetTenantID(pi.TenantID).
				SetLinkType(identitylink.LinkTypeCandidate).SetOperationType(identitylink.OperationTypeInit).
				SetGroupID(gid)
			idbuilders = append(idbuilders, idb)
		}
		for _, uid := range candidateUsers {
			idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetTenantID(pi.TenantID).
				SetLinkType(identitylink.LinkTypeCandidate).SetOperationType(identitylink.OperationTypeInit).
				SetUserID(int(uid))
			idbuilders = append(idbuilders, idb)
		}
		_, err = tx.IdentityLink.CreateBulk(idbuilders...).Save(ctx)

		return err
	})
	return
}

// ClaimUserTask 用户主动认领任务,并更新任务统计后,将任务数据返回给工作流引擎.
//
// 在创建用户任务时 CreateUserTask会默认的更新任务人数统计,因此在本方法执行时,只有候选组或参考人主动认领任务时,才会再次更新任务人数统计.
// 需要注意added返回值的处理.
func (e *DbExport) ClaimUserTask(ctx context.Context, input *ent.CreateIdentityLinkInput) (added bool, err error) {
	ctx = schemax.SkipTenantPrivacy(ctx)
	if input.UserID == nil || *input.UserID == 0 {
		return false, ErrUserIDRequire
	}
	if input.TenantID == 0 {
		return false, ErrTenantIDRequire
	}
	links, err := e.Service.Db.IdentityLink.Query().Where(
		identitylink.TaskID(input.TaskID),
		identitylink.TenantID(input.TenantID),
		identitylink.OperationTypeEQ(identitylink.OperationTypeInit),
		identitylink.HasTaskWith(entask.StatusEQ(entask.StatusRunning)),
	).WithTask(func(query *ent.TaskQuery) {
		query.Select(entask.FieldSequential)
	}).Order(identitylink.ByID()).All(ctx)
	if err != nil {
		return
	}
	if len(links) == 0 {
		return false, errors.New("task not found")
	}
	var (
		group  []int
		linked *ent.IdentityLink
	)
LOOP:
	for i, link := range links {
		switch link.LinkType {
		case identitylink.LinkTypeAssignee:
			if link.UserID != *input.UserID {
				return false, errors.New("task already claimed")
			}
			linked = link
			break LOOP
		case identitylink.LinkTypeCandidate:
			if link.GroupID != 0 {
				group = append(group, link.GroupID)
			}
			if link.UserID == *input.UserID {
				linked = link
				break LOOP
			}
		case identitylink.LinkTypeParticipant:
			if link.UserID != *input.UserID {
				continue
			}
			task, err := link.Task(ctx)
			if err != nil { // 不应该发生
				return false, err
			}
			if task.Sequential {
				if i > 1 {
					lop := links[i-1].OperationType
					if !(lop == identitylink.OperationTypePass || lop == identitylink.OperationTypeDelete) {
						// 上一个人 通过或删除的操作才能认领
						return false, errors.New("the previous user is not completed")
					}
				}
			}
			linked = link
			break LOOP
		}
	}
	if linked == nil && len(group) == 0 {
		return false, ErrIdentityNotFound
	}
	if linked == nil {
		// 该任务对应的用户组,进行用户角色关联判断
		has, err := e.Service.Db.OrgRoleUser.Query().Where(
			orgroleuser.OrgRoleIDIn(group...),
			orgroleuser.UserID(*input.UserID)).Exist(ctx)
		if err != nil {
			return false, err
		}
		if !has {
			return false, ErrIdentityNotFound
		}
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.Db.Tx(ctx)
	}, func(vtx ecx.Transactor) error {
		tx := vtx.(*ent.Tx)
		if linked != nil {
			err = tx.IdentityLink.UpdateOne(linked).SetOperationType(identitylink.OperationTypeClaim).Exec(ctx)
			if err != nil {
				return err
			}
			// 直接指定任务人的任务,不需要再更新任务统计
			if linked.LinkType == identitylink.LinkTypeAssignee {
				return nil
			}
		} else {
			err = tx.IdentityLink.Create().SetInput(*input).SetTaskID(input.TaskID).
				SetLinkType(identitylink.LinkTypeAssignee).SetTenantID(input.TenantID).
				SetOperationType(identitylink.OperationTypeClaim).
				Exec(ctx)
			if err != nil {
				return err
			}
			added = true
		}
		return tx.Task.Update().AddMemberCount(1).AddUnfinishedCount(1).Exec(ctx)
	})
	return added, err
}

func (e *DbExport) Review(ctx context.Context, input *ent.UpdateIdentityLinkInput, lastUndone int) error {
	if input.UserID == nil || *input.UserID == 0 {
		return ErrUserIDRequire
	}
	if input.TaskID == nil || *input.TaskID == 0 {
		return ErrTaskIDRequire
	}
	if input.Comments == nil {
		input.Comments = new(string)
	}
	ctx = schemax.SkipTenantPrivacy(ctx)
	taskID := *input.TaskID
	lk, err := e.Service.Db.IdentityLink.Query().Where(identitylink.TaskID(taskID), identitylink.UserID(*input.UserID),
		identitylink.OperationTypeEQ(identitylink.OperationTypeClaim)).Only(ctx)
	if err != nil {
		return err
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.Db.Tx(ctx)
	}, func(vtx ecx.Transactor) error {
		tx := vtx.(*ent.Tx)
		err = tx.IdentityLink.UpdateOneID(lk.ID).SetComments(*input.Comments).SetOperationType(*input.OperationType).Exec(ctx)
		if err != nil {
			return err
		}

		lastUndone--
		tku := tx.Task.UpdateOneID(taskID).AddUnfinishedCount(-1)
		switch *input.OperationType {
		case identitylink.OperationTypePass:
			tku.AddAgreeCount(1)
		case identitylink.OperationTypeReject:
		}
		if lastUndone == 0 {
			tku.SetStatus(entask.StatusFinished)
		}
		return tku.Exec(ctx)
	})
	return err
}
