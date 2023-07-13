package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/conv"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/service/workflow"
	"go.temporal.io/sdk/worker"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/workflow/ent/runtime"
)

func main() {
	app := woocoo.New()
	if err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake")); err != nil {
		log.Fatal(err)
	}

	service, err := workflow.NewService(app.AppConfiguration())
	if err != nil {
		log.Fatal(err)
	}
	defer service.Client.Close()

	service.Db = buildEntClient(app.AppConfiguration())

	var wopts worker.Options
	if err := app.AppConfiguration().Sub("temporal.workerOptions").Unmarshal(&wopts); err != nil {
		log.Fatal(err)
	}

	bpmn.Convert = conv.NewDefaultConverter()
	def := workflow.Definition{
		Service: service,
		Exporter: &workflow.DbExport{
			Service: service,
		},
	}
	service.WorkflowType = app.AppConfiguration().String("temporal.type")
	w := worker.New(service.Client, service.TaskQueue, wopts)
	w.RegisterWorkflow(def.BPMNWorkflowDef)
	eg := engine.NewBPMN(engine.WithExporter(def.Exporter))
	w.RegisterActivity(eg.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Unable to start worker:%v", err)
	}
}

func buildEntClient(cnf *conf.AppConfiguration) (client *ent.Client) {
	pd := oteldriver.BuildOTELDriver(cnf, "store.workflow")
	pd = ecx.BuildEntCacheDriver(cnf, pd)
	scfg := workflow.AlternateSchema()
	if cnf.Development {
		client = ent.NewClient(ent.Driver(pd), ent.Debug(), scfg)
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd), casbinent.Debug())
	} else {
		client = ent.NewClient(ent.Driver(pd), scfg)
		//casbinClient = casbinent.NewClient(casbinent.Driver(pd))
	}
	return
}
