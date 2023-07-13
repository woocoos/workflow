package main

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/workflow/api/graphql"
	"github.com/woocoos/workflow/ent"
	_ "github.com/woocoos/workflow/ent/runtime"
	"github.com/woocoos/workflow/service/deployment"
	"github.com/woocoos/workflow/service/workflow"
)

var (
	dbClient *ent.Client
)

func main() {
	app := woocoo.New()
	cnf := app.AppConfiguration()
	dbClient = buildEntClient(cnf)
	defer dbClient.Close()

	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Fatal(err)
	}

	wf, err := workflow.NewService(cnf)
	if err != nil {
		log.Fatal(err)
	}
	wf.Db = dbClient
	ds := &deployment.Service{
		WFDB:   dbClient,
		Engine: wf,
	}
	webSrv := buildWebServer(app.AppConfiguration(), ds)
	app.RegisterServer(webSrv)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildWebServer(cnf *conf.AppConfiguration, service *deployment.Service) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		web.RegisterMiddleware(authz.New()),
		identity.RegistryTenantIDMiddleware(),
	)
	// 如果需要设置gqlserver的中间件,可以使用第一个返回值
	gqlsrv := handler.NewDefaultServer(
		graphql.NewSchema(
			graphql.WithEntClient(dbClient),
			graphql.WithDeploymentService(service),
		),
	)
	gqlsrv.AroundResponses(gqlx.ContextCache())
	gqlsrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlsrv.Use(entgql.Transactioner{TxOpener: dbClient})
	if err := gql.RegisterGraphqlServer(webSrv, gqlsrv); err != nil {
		log.Fatal(err)
	}
	return webSrv
}

func buildEntClient(cnf *conf.AppConfiguration) *ent.Client {
	pd := oteldriver.BuildOTELDriver(cnf, "store.workflow")
	pd = ecx.BuildEntCacheDriver(cnf, pd)
	scfg := workflow.AlternateSchema()
	if cnf.Development {
		dbClient = ent.NewClient(ent.Driver(pd), ent.Debug(), scfg)
	} else {
		dbClient = ent.NewClient(ent.Driver(pd), scfg)
	}
	return dbClient
}
