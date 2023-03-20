package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Deployments is the resolver for the deployments field.
func (r *queryResolver) Deployments(ctx context.Context) ([]*ent.Deployment, error) {
	panic(fmt.Errorf("not implemented: Deployments - deployments"))
}

// ProcInsts is the resolver for the procInsts field.
func (r *queryResolver) ProcInsts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.ProcInstOrder, where *ent.ProcInstWhereInput) (*ent.ProcInstConnection, error) {
	panic(fmt.Errorf("not implemented: ProcInsts - procInsts"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.TaskWhereInput) (*ent.TaskConnection, error) {
	return r.DB.Task.Query().Paginate(ctx, after, first, before, last, ent.WithTaskFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }