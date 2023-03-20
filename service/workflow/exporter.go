package workflow

import (
	"context"
	"errors"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/procdef"
	entask "github.com/woocoos/workflow/ent/task"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/activity"
)

var (
	_ engine.Exporter = (*DbExport)(nil)

	ErrUserIDRequire = errors.New("user id is required")
	ErrTaskIDRequire = errors.New("task id is required")
)

type DbExport struct {
	Service *Service
}

func (e *DbExport) GetProcDef(ctx context.Context, req *engine.GetProcDefRequest) (*ent.ProcDef, error) {
	return e.Service.WFDB.ProcDef.Query().Where(procdef.Key(req.ProcDefKey), procdef.OrgID(req.OrgID)).First(ctx)
}

func (e *DbExport) GetDecisionReqDef(ctx context.Context, req *engine.GetDecisionReqDefRequest) (*ent.DecisionReqDef, error) {
	def, err := e.Service.WFDB.DecisionDef.Query().Where(decisiondef.Key(req.DecisionDefKey), decisiondef.OrgID(req.OrgID)).
		Order(ent.Desc(decisiondef.FieldVersion)).WithReqDef().First(ctx)
	if err != nil {
		return nil, err
	}
	return def.Edges.ReqDefOrErr()
}

func (e *DbExport) GetUserIDs(ctx context.Context, req *engine.GetUserIDsRequest) (*engine.GetUserIDResponse, error) {
	ids, err := e.Service.PortalDB.OrgUser.Query().Where(orguser.OrgID(req.OrgID),
		orguser.DisplayNameIn(req.Names...)).Select(orguser.FieldUserID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	return &engine.GetUserIDResponse{UserIDs: ids}, nil
}

func (e *DbExport) GetGroupIDs(ctx context.Context, req *engine.GetGroupIDsRequest) (*engine.GetGroupIDResponse, error) {
	ids, err := e.Service.PortalDB.OrgRole.Query().Where(orgrole.OrgID(req.OrgID),
		orgrole.NameIn(req.Names...), orgrole.KindEQ(orgrole.KindGroup)).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return &engine.GetGroupIDResponse{GroupIDs: ids}, nil
}

func (e *DbExport) CreateUserTask(ctx context.Context, pi *engine.InstanceRequest, taskEle *bpmn.UserTask) (tk *ent.Task, err error) {
	exec := activity.GetInfo(ctx).WorkflowExecution
	if taskEle.Assignment == nil {
		return nil, nil
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.WFDB.Tx(ctx)
	}, func(vtx ecx.Transactor) error {
		tx := vtx.(*ent.Tx)
		cnt := int32(taskEle.Assignment.CountAssignees())
		taskbuilder := tx.Task.Create().SetProcInstID(pi.ID).SetOrgID(pi.OrgID).SetProcDefID(pi.ProcDefID).
			SetTaskDefKey(taskEle.Id).SetComments(taskEle.GetDocumentation()).
			SetExecutionID(exec.ID).SetRunID(exec.RunID).
			SetMemberCount(cnt).SetUnfinishedCount(cnt).SetStatus(entask.StatusRunning)
		if taskEle.Assignment.Assignee != "" {
			assignee, err := taskEle.Assignment.GetAssignee(pi.Variables)
			if err != nil {
				return err
			}
			taskbuilder.SetAssignee(assignee)
		}
		if taskEle.MultiInstance != nil {
			taskbuilder.SetSequential(taskEle.MultiInstance.IsSequential)
			taskbuilder.SetKind(entask.KindAND)
		}
		tk, err = taskbuilder.Save(ctx)
		if err != nil {
			return err
		}

		idbuilders := make([]*ent.IdentityLinkCreate, 0)
		if tk.Assignee != "" {
			resp, err := e.GetUserIDs(ctx, &engine.GetUserIDsRequest{Names: []string{tk.Assignee}, OrgID: pi.OrgID})
			if err != nil {
				return err
			}
			ids := resp.UserIDs
			idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetOrgID(pi.OrgID).
				SetLinkType(identitylink.LinkTypeAssignee).SetOperationType(identitylink.OperationTypeAdd).
				SetUserID(ids[0])
			idbuilders = append(idbuilders, idb)
		}
		if taskEle.Assignment.CandidateGroups != "" {
			cgs, err := taskEle.Assignment.GetCandidateGroups(pi.Variables)
			if err != nil {
				return err
			}
			resp, err := e.GetGroupIDs(ctx, &engine.GetGroupIDsRequest{Names: cgs, OrgID: pi.OrgID})
			for _, gid := range resp.GroupIDs {
				idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetOrgID(pi.OrgID).
					SetLinkType(identitylink.LinkTypeCandidate).SetOperationType(identitylink.OperationTypeAdd).
					SetGroupID(gid)
				idbuilders = append(idbuilders, idb)
			}
		}
		if taskEle.Assignment.CandidateUsers != "" {
			cus, err := taskEle.Assignment.GetCandidateUsers(pi.Variables)
			if err != nil {
				return err
			}
			resp, err := e.GetUserIDs(ctx, &engine.GetUserIDsRequest{Names: cus, OrgID: pi.OrgID})
			if err != nil {
				return err
			}
			for _, uid := range resp.UserIDs {
				idb := tx.IdentityLink.Create().SetTaskID(tk.ID).SetProcDefID(pi.ProcDefID).SetOrgID(pi.OrgID).
					SetLinkType(identitylink.LinkTypeCandidate).SetOperationType(identitylink.OperationTypeAdd).
					SetUserID(int(uid))
				idbuilders = append(idbuilders, idb)
			}
		}
		_, err = tx.IdentityLink.CreateBulk(idbuilders...).Save(ctx)

		return err
	})
	return
}

func (e *DbExport) ClaimUserTask(ctx context.Context, input *ent.CreateIdentityLinkInput) (err error) {
	if input.UserID == nil || *input.UserID == 0 {
		return ErrUserIDRequire
	}
	ol, _ := e.Service.WFDB.IdentityLink.Query().Where(identitylink.TaskID(input.TaskID), identitylink.UserID(*input.UserID),
		identitylink.LinkTypeEQ(identitylink.LinkTypeAssignee)).OnlyID(ctx)
	if ol != 0 { // 已经被委派
		return nil
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.WFDB.Tx(ctx)
	}, func(vtx ecx.Transactor) error {
		tx := vtx.(*ent.Tx)
		err = tx.IdentityLink.Create().SetInput(*input).SetTaskID(input.TaskID).
			SetLinkType(identitylink.LinkTypeAssignee).SetOrgID(input.OrgID).
			SetOperationType(identitylink.OperationTypeAdd).
			Exec(ctx)
		if err != nil {
			return err
		}
		return tx.Task.Update().AddMemberCount(1).AddUnfinishedCount(1).Exec(ctx)
	})
	return err
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
	taskID := *input.TaskID
	lk, err := e.Service.WFDB.IdentityLink.Query().Where(identitylink.TaskID(taskID), identitylink.UserID(*input.UserID),
		identitylink.LinkTypeEQ(identitylink.LinkTypeAssignee)).Only(ctx)
	if err != nil {
		return err
	}
	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return e.Service.WFDB.Tx(ctx)
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
