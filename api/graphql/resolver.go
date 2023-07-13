package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/workflow/api/graphql/generated"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/service/deployment"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client     *ent.Client
	deployment *deployment.Service
}

type Option func(*Resolver)

// WithEntClient sets the ent client on the resolver.
func WithEntClient(client *ent.Client) Option {
	return func(r *Resolver) {
		r.client = client
	}
}

// WithDeploymentService sets the deployment service on the resolver.
func WithDeploymentService(deployment *deployment.Service) Option {
	return func(r *Resolver) {
		r.deployment = deployment
	}
}

// NewResolver creates a new resolver with the given options.
func NewResolver(opts ...Option) *Resolver {
	r := &Resolver{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...Option) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: NewResolver(opts...),
	})
}
