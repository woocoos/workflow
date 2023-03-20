package graph

import (
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/service/deployment"
)

var (
	Err = false
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB         *ent.Client
	Deployment *deployment.Service
}
