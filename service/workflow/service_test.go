package workflow

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/pkg/security"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	_ "github.com/woocoos/workflow/ent/runtime"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/test"
	"go.temporal.io/sdk/client"
	"os"
	"path/filepath"
)

func (ts *TestSuite) Test_RunSimpleTask() {
	file := filepath.Join(test.BaseDir(), "case", "simple-task.bpmn")
	inst := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		ResourcePath: file,
		ProcDefKey:   "Simple_Task_Process",
		ResourceName: "simple_task.bpmn",
		Variables:    bpmn.Mappings{},
	}
	wf, err := ts.Service.Client.ExecuteWorkflow(context.Background(),
		client.StartWorkflowOptions{
			ID:        inst.ProcDefKey,
			TaskQueue: ts.Service.TaskQueue,
		},
		ts.Service.WorkflowType, inst)
	ts.NoError(err)
	ts.NoError(wf.Get(context.Background(), nil))
}

func (ts *TestSuite) Test_RunSimpleUserTask() {
	file := filepath.Join(test.BaseDir(), "case", "simple-user-task.bpmn")
	inst := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		ResourcePath: file,
		ProcDefKey:   "simple-user-task",
		ResourceName: "simple_user_task.bpmn",
		Variables: bpmn.Mappings{
			"assignee": "admin",
		},
	}
	wf, err := ts.Service.Client.ExecuteWorkflow(context.Background(),
		client.StartWorkflowOptions{
			ID:        inst.ProcDefKey,
			TaskQueue: ts.Service.TaskQueue,
		},
		ts.Service.WorkflowType, inst)
	ts.NoError(err)
	ts.NotNil(wf.GetID())
	err = ts.Service.Client.SignalWorkflow(context.Background(), wf.GetID(), "", engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
		UserID: gds.Ptr(198640048702208),
	})
	ts.NoError(err)
	err = ts.Service.Client.SignalWorkflow(context.Background(), wf.GetID(), "", engine.UserTaskReviewChannel, ent.CreateIdentityLinkInput{
		UserID:        gds.Ptr(198640048702208),
		OperationType: identitylink.OperationTypePass,
		Comments:      gds.Ptr("pass the task"),
	})
	ts.NoError(err)
}

func (ts *TestSuite) Test_RunCallActivity() {
	callfile := filepath.Join(test.BaseDir(), "case", "call-activity.bpmn")
	chidlfiel := filepath.Join(test.BaseDir(), "case", "simple-user-task.bpmn")
	cbs, err := os.ReadFile(chidlfiel)
	ts.Require().NoError(err)
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{
		"sub": "1",
	}))
	has, err := ts.WFDB.ProcDef.Query().Where(procdef.OrgID(1), procdef.Key("simple-user-task")).Exist(ctx)
	ts.Require().NoError(err)
	if !has {
		_, err = ts.WFDB.ProcDef.Create().SetOrgID(1).
			SetKey("simple-user-task").
			SetResourceName("simple-user-task.bpmn").
			SetResourceData(cbs).
			SetDeploymentID(1).
			SetAppID(1).
			Save(ctx)
		ts.Require().NoError(err)
	}
	inst := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			OrgID:  1,
			Status: procinst.StatusReady,
		},
		ResourcePath: callfile,
		ProcDefKey:   "call-activity",
		ResourceName: "call-activity.bpmn",
		Variables: bpmn.Mappings{
			"assignee": "admin",
			"price":    1,
		},
	}
	wf, err := ts.Service.Client.ExecuteWorkflow(context.Background(),
		client.StartWorkflowOptions{
			ID:        "9ea1ef92-c981-45a7-b489-5ef1c37ccfe7",
			TaskQueue: ts.Service.TaskQueue,
		},
		ts.Service.WorkflowType, inst)
	ts.NoError(err)
	err = ts.Service.Client.SignalWorkflow(context.Background(), wf.GetID(), "", engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
		UserID: gds.Ptr(198640048702208),
	})
	ts.NoError(err)
	err = ts.Service.Client.SignalWorkflow(context.Background(), wf.GetID(), "", engine.UserTaskReviewChannel, ent.CreateIdentityLinkInput{
		UserID:        gds.Ptr(198640048702208),
		OperationType: identitylink.OperationTypePass,
		Comments:      gds.Ptr("pass the task"),
	})
	ts.NoError(err)
}
