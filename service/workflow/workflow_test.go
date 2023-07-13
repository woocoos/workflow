package workflow

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/test"
	wfsuite "github.com/woocoos/workflow/test/testsuite"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/workflow/ent/runtime"
)

func SetInstanceRequestDef(ir *engine.InstanceRequest, def *ent.ProcDef) {
	if def.ID != 0 {
		ir.ProcDefID = def.ID
	}
	ir.ProcDefKey = def.Key
	ir.ResourceKey = def.ResourceKey
}

type TestSuite struct {
	wfsuite.BaseSuite
	Service *Service
	def     Definition
	testsuite.WorkflowTestSuite

	//env *testsuite.TestWorkflowEnvironment
}

func (ts *TestSuite) SetupTest() {
	//ts.env = ts.NewTestWorkflowEnvironment()
}

func (ts *TestSuite) SetupSuite() {
	ts.BaseSuite.SetupSuite()
	ts.Service = &Service{
		Db:           ts.BaseSuite.Db,
		TaskQueue:    "approval",
		Client:       ts.BaseSuite.Client,
		WorkflowType: "BPMNWorkflowDef",
		ResourceDir:  test.BaseDir(),
	}
	ts.def = Definition{
		Service: ts.Service,
		Exporter: &FileExport{
			DbExport: DbExport{
				Service: ts.Service,
			},
		},
	}
}

func (ts *TestSuite) TearDownSuite() {
	if ts.Client != nil {
		ts.Client.Close()
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (ts *TestSuite) newBPMN(id string, exporter engine.Exporter) *engine.BPMN {
	return engine.NewBPMN(engine.WithID(id), engine.WithExporter(exporter),
		engine.WithResourceDir(ts.Service.ResourceDir))
}

func (ts *TestSuite) Test_SimpleTask() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			TenantID: 1,
			Status:   procinst.StatusReady,
		},
		Variables: bpmn.Mappings{},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "simple-task", TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := ts.newBPMN(wf.ProcDefKey, ts.def)
	a.Handlers.RegistryHandler("Sleep", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		time.Sleep(time.Second * 2)
		return nil, nil
	})
	env.RegisterActivity(a.Activity)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

func (ts *TestSuite) Test_SimpleUserTask_WrongUser() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1001,
			ProcDefID: 1,
			TenantID:  1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"assignee": "user1",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "simple-user-task", TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			TenantID: 1,
			UserID:   gds.Ptr(198640048702208),
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(198640048702208),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)

	a := ts.newBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a.Activity)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.Error(env.GetWorkflowError())
}

func (ts *TestSuite) Test_SimpleUserTask() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1001,
			ProcDefID: 1,
			TenantID:  1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"assignee": "user1",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "simple-user-task", TenantID: wf.TenantID})
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

	a := ts.newBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "parallel-gateway", TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := ts.newBPMN(wf.ProcDefKey, ts.def)

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
	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "exclusive-gateway", TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	a := ts.newBPMN(wf.ProcDefKey, &ts.def)

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
	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "service-task-input-output",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := ts.newBPMN(wf.ProcDefKey, &ts.def)

	a.Handlers.RegistryHandler("input-task-1", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		activity.GetLogger(ctx).Info("step1", "vars", vars)
		return bpmn.Mappings{
			"step1": "step1",
		}, nil
	})
	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-timer",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := ts.newBPMN(wf.ProcDefKey, &ts.def)

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
	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-timer",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)

	SetInstanceRequestDef(&wf, pd)
	a := ts.newBPMN(wf.ProcDefKey, ts.def)

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

	env.RegisterActivity(a.Activity)
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
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "message-loop",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := ts.newBPMN(wf.ProcDefKey, ts.def)
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
	env.RegisterActivity(a.Activity)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

// TODO 目前由于子流程与主流程为同一个实例,在测试环境中,无法监控到子流程的执行,而导致测试失败
func (ts *TestSuite) Test_CallActivity() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			TenantID:  1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"assignee": "user1",
			"price":    1,
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "call-activity",
		TenantID: wf.TenantID})
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

	a := ts.newBPMN(wf.ProcDefKey, ts.def)

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
	env.RegisterActivity(a.Activity)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())
}

// MultiUserTask is a multi candidate users task. User1 is not allow user.
func (ts *TestSuite) Test_MultiUserTask() {
	env := ts.NewTestWorkflowEnvironment()
	wf := engine.InstanceRequest{
		ProcInst: &ent.ProcInst{
			ID:        1,
			ProcDefID: 1,
			TenantID:  1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "multi-user-task",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID:   gds.Ptr(1),
			TenantID: 1,
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
			UserID:   gds.Ptr(2),
			TenantID: 1,
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(2),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskClaimChannel, ent.CreateIdentityLinkInput{
			UserID:   gds.Ptr(3),
			TenantID: 1,
		})
	}, time.Hour)
	env.RegisterDelayedCallback(func() {
		env.SignalWorkflow(engine.UserTaskReviewChannel, ent.UpdateIdentityLinkInput{
			UserID:        gds.Ptr(3),
			OperationType: gds.Ptr(identitylink.OperationTypePass),
			Comments:      gds.Ptr("pass"),
		})
	}, time.Hour+time.Second)
	a := ts.newBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a.Activity)
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
			TenantID:  1,
			Status:    procinst.StatusReady,
		},
		Variables: bpmn.Mappings{
			"amount":          100,
			"invoiceCategory": "Misc",
		},
	}
	pd, err := ts.def.GetProcDef(context.Background(), &engine.GetProcDefRequest{ProcDefKey: "business-rule",
		TenantID: wf.TenantID})
	ts.Require().NoError(err)
	SetInstanceRequestDef(&wf, pd)

	a := ts.newBPMN(wf.ProcDefKey, ts.def)
	env.RegisterActivity(a.Activity)
	env.ExecuteWorkflow(ts.def.BPMNWorkflowDef, wf)
	ts.True(env.IsWorkflowCompleted())
	ts.NoError(env.GetWorkflowError())

}
