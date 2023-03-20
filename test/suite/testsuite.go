package suite

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/pkg/snowflake"
	portal "github.com/woocoos/knockout/ent"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/test"
	"go.temporal.io/sdk/client"
	"os"
	"path/filepath"
)

var (
	wfdns = "root:@tcp(localhost:3306)/workflow?parseTime=true&loc=Local"
	ptdns = "root:@tcp(localhost:3306)/portal?parseTime=true&loc=Local"
)

type WFSuite struct {
	WFDB     *ent.Client
	PortalDB *portal.Client
	Client   client.Client
}

func (s *WFSuite) SetupSuite() {
	file := filepath.Join(test.BaseDir(), "suite", "etc", "app.yaml")
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
	s.WFDB, err = ent.Open("mysql", wfdns, ent.Debug())
	if err != nil {
		panic(err)
	}
	s.PortalDB, err = portal.Open("mysql", ptdns, portal.Debug())
	if err != nil {
		panic(err)
	}
	co := client.Options{}
	err = conf.Global().Sub("temporal.clientOptions").Unmarshal(&co)
	if err != nil {
		panic(err)
	}
	s.Client, err = client.Dial(co)
}
