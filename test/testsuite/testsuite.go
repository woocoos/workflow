package testsuite

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/migrate"
	"github.com/woocoos/workflow/ent/orgrole"
	"github.com/woocoos/workflow/pkg/conv"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/test"
	"go.temporal.io/sdk/client"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//wfdns = "root:@tcp(localhost:3306)/workflow?parseTime=true&loc=Local"
	driverName = "sqlite3"
	dsn        = "file:workflow?mode=memory&cache=shared&_fk=1"
)

type BaseSuite struct {
	suite.Suite
	Db     *ent.Client
	Client client.Client
}

func (s *BaseSuite) SetupSuite() {
	file := filepath.Join(test.BaseDir(), "testsuite", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()
	app := woocoo.New(woocoo.WithAppConfiguration(cnf))

	err = snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		panic(err)
	}

	clientOpts := client.Options{}
	err = cnf.Sub("temporal.clientOptions").Unmarshal(&clientOpts)
	if err != nil {
		panic(err)
	}
	s.Client, err = client.Dial(clientOpts)

	scfg := ent.AlternateSchema(ent.SchemaConfig{
		OrgUser:     "portal",
		OrgRole:     "portal",
		OrgApp:      "portal",
		OrgRoleUser: "portal",
	})
	s.Db, err = ent.Open(driverName, dsn, ent.Debug(), scfg)
	if err != nil {
		panic(err)
	}
	s.initData()

	bpmn.Convert = conv.NewDefaultConverter()
}

func (s *BaseSuite) initData() {
	db, err := ent.Open(driverName, dsn)
	s.Require().NoError(err)
	err = db.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	)
	s.Require().NoError(err)

	ctx := schemax.SkipTenantPrivacy(context.Background())
	oub := make([]*ent.OrgUserCreate, 3)
	orb := make([]*ent.OrgRoleCreate, 3)
	orbu := make([]*ent.OrgRoleUserCreate, 3)
	for i := 1; i < 4; i++ {
		idx := i - 1
		oub[idx] = db.OrgUser.Create().SetID(i + 10).SetUserID(i).SetDisplayName(fmt.Sprintf("user%d", i)).SetOrgID(1)
		orb[idx] = db.OrgRole.Create().SetID(i + 10).SetName(fmt.Sprintf("role%d", i)).SetOrgID(1).SetKind(orgrole.KindGroup)
		orbu[idx] = db.OrgRoleUser.Create().SetOrgRoleID(i + 10).SetOrgUserID(i + 10).SetUserID(i).SetOrgID(i)
	}
	db.OrgApp.Create().SetOrgID(1).SetAppID(10).SaveX(ctx)
	db.OrgUser.CreateBulk(oub...).SaveX(ctx)
	db.OrgRole.CreateBulk(orb...).SaveX(ctx)
	db.OrgRoleUser.CreateBulk(orbu...).SaveX(ctx)

	db.ProcDef.Create().SetID(1633283735837216769).SetCreatedBy(1).SetAppID(10).SetTenantID(1).
		SetDeploymentID(1).SetKey("invoice").SetResourceKey("case/invoice.v2.bpmn").SetName("Invoice Receipt").
		SetStatus(typex.SimpleStatusActive).SaveX(ctx)
	db.ProcDef.Create().SetID(1000).SetCreatedBy(1).SetAppID(10).SetTenantID(1).
		SetDeploymentID(1000).SetKey("simple-task").SetResourceKey("case/simple-task.bpmn").SetName("").
		SetStatus(typex.SimpleStatusActive).SaveX(ctx)
	db.ProcDef.Create().SetID(1001).SetCreatedBy(1).SetAppID(10).SetTenantID(1).
		SetDeploymentID(1000).SetKey("simple-user-task").SetResourceKey("case/simple-user-task.bpmn").SetName("").
		SetStatus(typex.SimpleStatusActive).SaveX(ctx)
}
