package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/snowflake"
	portal "github.com/woocoos/knockout/ent"
	"github.com/woocoos/workflow/ent"
	_ "github.com/woocoos/workflow/ent/runtime"
	"github.com/woocoos/workflow/graph"
	"github.com/woocoos/workflow/graph/generated"
	"github.com/woocoos/workflow/identity"
	"github.com/woocoos/workflow/service/deployment"
	"github.com/woocoos/workflow/service/workflow"
	"go.temporal.io/sdk/client"
)

var (
	wfClient     *ent.Client
	portalClient *portal.Client
)

var tenantContextKey = "github.com_woocoos_entco_tenant_id"

func main() {
	app := woocoo.New()
	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Fatal(err)
	}

	wfdri := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.workflow")
	wfdri = ecx.BuildEntCacheDriver(app.AppConfiguration(), wfdri)

	ptdri := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.portal")
	ptdri = ecx.BuildEntCacheDriver(app.AppConfiguration(), ptdri)
	if app.AppConfiguration().Development {
		wfClient = ent.NewClient(ent.Driver(wfdri), ent.Debug())
		portalClient = portal.NewClient(portal.Driver(ptdri), portal.Debug())
	} else {
		wfClient = ent.NewClient(ent.Driver(wfdri))
		portalClient = portal.NewClient(portal.Driver(ptdri))
	}

	buildCashbin(app.AppConfiguration(), ptdri)

	var topts client.Options
	if err := app.AppConfiguration().Sub("temporal.clientOptions").Unmarshal(&topts); err != nil {
		log.Fatal(err)
	}
	tc, err := client.Dial(topts)
	if err != nil {
		log.Fatalf("Unable to create client:%v", err)
	}
	gr := &graph.Resolver{
		DB: wfClient,
		Deployment: &deployment.Service{
			WFDB:     wfClient,
			PortalDB: portalClient,
			Engine: workflow.Service{
				WFDB:      wfClient,
				PortalDB:  portalClient,
				Client:    tc,
				TaskQueue: app.AppConfiguration().String("temporal.taskQueue"),
			},
		},
	}
	webSrv := newWebServer(app.AppConfiguration(), gr)
	app.RegisterServer(webSrv)

	defer func() {
		wfClient.Close()
		portalClient.Close()
	}()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func newWebServer(cnf *conf.AppConfiguration, resolver *graph.Resolver) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		web.RegisterMiddleware(authz.New()),
		identity.RegistryTenantIDMiddleware(),
	)
	// 如果需要设置gqlserver的中间件,可以使用第一个返回值
	gqlSrv, err := gql.RegisterSchema(webSrv,
		generated.NewExecutableSchema(generated.Config{
			Resolvers: resolver,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	// gqlserver的中间件处理
	for _, srv := range gqlSrv {
		if cnf.String("entcache.level") == "context" {
			// 启用针对http request的缓存
			useContextCache(srv)
		}
		// mutation事务
		srv.Use(entgql.Transactioner{TxOpener: wfClient})
	}
	return webSrv
}

func buildCashbin(cnf *conf.AppConfiguration, portalDriver dialect.Driver) {
	_, err := authorization.SetAuthorization(cnf.Sub("authz"), portalDriver)
	if err != nil {
		log.Fatal(err)
	}
}

func useContextCache(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})
}
