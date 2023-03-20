package workflow

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	portal "github.com/woocoos/knockout/ent"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	wfsuite "github.com/woocoos/workflow/test/suite"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
	"testing"
	"time"

	_ "github.com/woocoos/knockout/ent/runtime"
	_ "github.com/woocoos/workflow/ent/runtime"
)

func SetInstanceRequestDef(ir *engine.InstanceRequest, def *ent.ProcDef) {
	ir.ProcDefID = def.ID
	ir.ProcDefKey = def.Key
	ir.ResourceName = def.ResourceName
	ir.ResourceData = def.ResourceData
}

type TestSuite struct {
	wfsuite.WFSuite
	Service *Service
	def     Definition
	suite.Suite
	testsuite.WorkflowTestSuite
}

func (ts *TestSuite) SetupSuite() {
	builder := make([]*portal.OrgUserCreate, 3)
	for i := 0; i < 3; i++ {
		builder[i] = ts.Service.PortalDB.OrgUser.Create().SetID(i + 1).SetUserID(i + 1).SetOrgID(1)
	}
	ts.Service.PortalDB.OrgUser.CreateBulk(builder...)
}

func TestTestSuite(t *testing.T) {
	wfs := wfsuite.WFSuite{}
	wfs.SetupSuite()
	svc := &Service{
		WFDB:         wfs.WFDB,
		PortalDB:     wfs.PortalDB,
		TaskQueue:    "approval",
		Client:       wfs.Client,
		WorkflowType: "BPMNWorkflowDef",
	}
	defer wfs.Client.Close()
	suite.Run(t, &TestSuite{
		Service: svc,
		def: Definition{
			Service:   svc,
			TaskQueue: "approval",
			Exporter: &FileExport{
				DbExport: DbExport{
					Service: svc,
				},
			},
		},
	})
}

func (ts *TestSuite) Test_BuildWorkflow() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1633283735837216769,
			OrgID:     1,
			AppID:     1,
			Status:    procinst.StatusReady,
		},
		ProcDefKey:   "invoice",
		ResourceName: "invoice.bpmn",
	}
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_SimpleTask() {
	env := ts.NewTestWorkflowEnvironment()

	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "simple-task", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := engine.NewBPMN(wf.ProcDefKey, ts.def)
	a.Handlers.RegistryHandler("Sleep", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		time.Sleep(time.Second * 2)
		return nil, nil
	})
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_SimpleUserTask() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			OrgID:     1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"assignee": "admin",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "simple-user-task", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID: gds.Ptr(198640048702208),
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(198640048702208),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)

	a := engine.NewBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_ParallelGateway() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "parallel-gateway", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := engine.NewBPMN(wf.ProcDefKey, ts.def)

	a.Handlers.RegistryHandler("step1", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		return bpmn.Mappings{
			"step1": "step1",
		}, nil
	})
	a.Handlers.RegistryHandler("step2", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step2", "vars", vars)
		return bpmn.Mappings{
			"step2": "step2",
		}, nil
	})
	a.Handlers.RegistryHandler("step3", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step3", "vars", vars)
		return bpmn.Mappings{
			"step3": "step3",
		}, nil
	})
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_ExclusiveGateway() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"price": 1.5,
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "exclusive-gateway", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := engine.NewBPMN(wf.ProcDefKey, &ts.def)

	a.Handlers.RegistryHandler("task-a", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		return bpmn.Mappings{
			"step1": "step1",
		}, nil
	})
	a.Handlers.RegistryHandler("task-b", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		assert.Fail(ts.T(), "should not be called")
		return nil, nil
	})
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_ServiceTaskInputOutput() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"price": 1.5,
			"city":  "beijing",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "service-task-input-output", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := engine.NewBPMN(wf.ProcDefKey, &ts.def)

	a.Handlers.RegistryHandler("input-task-1", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		return bpmn.Mappings{
			"step1": "step1",
		}, nil
	})
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_MessageTimer() {
	env := ts.NewTestWorkflowEnvironment()

	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"key": "chan1",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-timer", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := engine.NewBPMN(wf.ProcDefKey, &ts.def)

	a.Handlers.RegistryHandler("ask", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("ask", "vars", vars)
		return nil, nil
	})
	a.Handlers.RegistryHandler("win", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		return nil, fmt.Errorf("should not be called")
	})
	checkLose := false
	a.Handlers.RegistryHandler("lose", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		checkLose = true
		return nil, nil
	})
	//env.OnActivity(a.ServiceTaskActivity, mock.Anything, mock.Anything).Return(engine.Mappings{}, nil)
	//env.OnActivity(a.CreateUserTaskActivity(),mock.Anything, mock.Anything).Return(nil, nil)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
	ts.True(checkLose)
}

func (ts *TestSuite) Test_MessageTimerWin() {
	env := ts.NewTestWorkflowEnvironment()

	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"key": "key",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-timer", OrgID: wf.OrgID})
	ts.Require().NoError(err)

	SetInstanceRequestDef(&wf, pd)
	a := engine.NewBPMN(wf.ProcDefKey, ts.def)

	a.Handlers.RegistryHandler("ask", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("ask", "vars", vars)
		return nil, nil
	})
	checkWin := false
	a.Handlers.RegistryHandler("win", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("ask", "vars", vars)
		checkWin = true
		return nil, nil
	})
	a.Handlers.RegistryHandler("lose", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		return nil, fmt.Errorf("should not be called")
	})
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("key", nil)
	}, time.Second*2)

	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
	ts.True(checkWin)
}

func (ts *TestSuite) Test_MessageLoop() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			Status: procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"key": "key",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-loop", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := engine.NewBPMN(wf.ProcDefKey, ts.def)
	loop := 0
	a.Handlers.RegistryHandler("do-nothing", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		loop++
		activity.GetLogger(ctx).Info("do-nothing", "loop", loop)
		return nil, nil
	})
	a.Handlers.RegistryHandler("validate", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		if loop > 2 {
			return bpmn.Mappings{
				"engineValidationAttempts": true,
				"hasReachedMaxAttempts":    true,
			}, nil
		}
		return bpmn.Mappings{
			"engineValidationAttempts": false,
			"hasReachedMaxAttempts":    false,
		}, nil
	})
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("key", bpmn.Mappings{
			"engineValidationAttempts": true,
			"hasReachedMaxAttempts":    false,
		})
	}, time.Second*1)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("key", bpmn.Mappings{
			"engineValidationAttempts": true,
			"hasReachedMaxAttempts":    false,
		})
	}, time.Second*2)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow("key", bpmn.Mappings{
			"engineValidationAttempts": true,
			"hasReachedMaxAttempts":    false,
		})
	}, time.Second*3)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_CallActivity() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			OrgID:     1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"assignee": "admin",
			"price":    1,
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "call-activity", OrgID: wf.OrgID})
	ts.Require().NoError(err)

	SetInstanceRequestDef(&wf, pd)

	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID: gds.Ptr(198640048702208),
		})
	}, time.Hour)

	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(198640048702208),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)

	a := engine.NewBPMN(wf.ProcDefKey, ts.def)

	a.Handlers.RegistryHandler("task-a", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		return bpmn.Mappings{
			"step1": "step1",
		}, nil
	})
	a.Handlers.RegistryHandler("task-b", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		assert.Fail(ts.T(), "should not be called")
		return nil, nil
	})

	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_MultiUserTask() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			OrgID:     1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "multi-user-task", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID: gds.Ptr(1),
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(1),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID: gds.Ptr(2),
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(3),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID: gds.Ptr(3),
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(3),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)
	a := engine.NewBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_BusinessRule() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			OrgID:     1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"amount":          100,
			"invoiceCategory": "Misc",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "business-rule", OrgID: wf.OrgID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := engine.NewBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())

}
