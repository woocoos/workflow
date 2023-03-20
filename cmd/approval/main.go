package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/snowflake"
	portalent "github.com/woocoos/knockout/ent"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/service/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
	_ "github.com/woocoos/workflow/ent/runtime"
)

func main() {
	app := woocoo.New()
	if err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake")); err != nil {
		log.Fatal(err)
	}

	co := client.Options{}
	err := app.AppConfiguration().Sub("temporal.clientOptions").Unmarshal(&co)
	if err != nil {
		log.Fatalf("Unable to build client option:%v", err)
	}

	c, err := client.Dial(co)
	if err != nil {
		log.Fatalf("Unable to create client:%v", err)
	}
	defer c.Close()

	service := workflow.Service{
		Client: c,
	}

	drv := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.workflow")
	drv = ecx.BuildEntCacheDriver(app.AppConfiguration(), drv)

	drvp := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.portal")
	drvp = ecx.BuildEntCacheDriver(app.AppConfiguration(), drvp)

	if app.AppConfiguration().Development {
		service.WFDB = ent.NewClient(ent.Driver(drv), ent.Debug())
		service.PortalDB = portalent.NewClient(portalent.Driver(drvp), portalent.Debug())
	} else {
		service.WFDB = ent.NewClient(ent.Driver(drv))
		service.PortalDB = portalent.NewClient(portalent.Driver(drvp))
	}

	var wopts worker.Options
	if err := app.AppConfiguration().Sub("temporal.workerOptions").Unmarshal(&wopts); err != nil {
		log.Fatal(err)
	}

	def := workflow.Definition{
		Service:   &service,
		TaskQueue: app.AppConfiguration().String("temporal.taskQueue"),
		Exporter: &workflow.DbExport{
			Service: &service,
		},
	}
	service.TaskQueue = def.TaskQueue
	service.WorkflowType = app.AppConfiguration().String("temporal.type")
	w := worker.New(c, def.TaskQueue, wopts)
	w.RegisterWorkflow(def.BPMNWorkflowDef)
	eg := engine.NewBPMN("", def.Exporter)
	w.RegisterActivity(eg)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Unable to start worker:%v", err)
	}
}
