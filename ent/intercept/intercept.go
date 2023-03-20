// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/predicate"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...ent.OrderFunc)
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The DecisionDefFunc type is an adapter to allow the use of ordinary function as a Querier.
type DecisionDefFunc func(context.Context, *ent.DecisionDefQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DecisionDefFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DecisionDefQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DecisionDefQuery", q)
}

// The TraverseDecisionDef type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDecisionDef func(context.Context, *ent.DecisionDefQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDecisionDef) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDecisionDef) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DecisionDefQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DecisionDefQuery", q)
}

// The DecisionReqDefFunc type is an adapter to allow the use of ordinary function as a Querier.
type DecisionReqDefFunc func(context.Context, *ent.DecisionReqDefQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DecisionReqDefFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DecisionReqDefQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DecisionReqDefQuery", q)
}

// The TraverseDecisionReqDef type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDecisionReqDef func(context.Context, *ent.DecisionReqDefQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDecisionReqDef) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDecisionReqDef) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DecisionReqDefQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DecisionReqDefQuery", q)
}

// The DeploymentFunc type is an adapter to allow the use of ordinary function as a Querier.
type DeploymentFunc func(context.Context, *ent.DeploymentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DeploymentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DeploymentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DeploymentQuery", q)
}

// The TraverseDeployment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDeployment func(context.Context, *ent.DeploymentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDeployment) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDeployment) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeploymentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DeploymentQuery", q)
}

// The IdentityLinkFunc type is an adapter to allow the use of ordinary function as a Querier.
type IdentityLinkFunc func(context.Context, *ent.IdentityLinkQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f IdentityLinkFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.IdentityLinkQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.IdentityLinkQuery", q)
}

// The TraverseIdentityLink type is an adapter to allow the use of ordinary function as Traverser.
type TraverseIdentityLink func(context.Context, *ent.IdentityLinkQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseIdentityLink) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseIdentityLink) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.IdentityLinkQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.IdentityLinkQuery", q)
}

// The ProcDefFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProcDefFunc func(context.Context, *ent.ProcDefQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ProcDefFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ProcDefQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ProcDefQuery", q)
}

// The TraverseProcDef type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProcDef func(context.Context, *ent.ProcDefQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProcDef) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProcDef) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProcDefQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ProcDefQuery", q)
}

// The ProcInstFunc type is an adapter to allow the use of ordinary function as a Querier.
type ProcInstFunc func(context.Context, *ent.ProcInstQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ProcInstFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ProcInstQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ProcInstQuery", q)
}

// The TraverseProcInst type is an adapter to allow the use of ordinary function as Traverser.
type TraverseProcInst func(context.Context, *ent.ProcInstQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseProcInst) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseProcInst) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProcInstQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ProcInstQuery", q)
}

// The TaskFunc type is an adapter to allow the use of ordinary function as a Querier.
type TaskFunc func(context.Context, *ent.TaskQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TaskFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TaskQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TaskQuery", q)
}

// The TraverseTask type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTask func(context.Context, *ent.TaskQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTask) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTask) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TaskQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TaskQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.DecisionDefQuery:
		return &query[*ent.DecisionDefQuery, predicate.DecisionDef]{typ: ent.TypeDecisionDef, tq: q}, nil
	case *ent.DecisionReqDefQuery:
		return &query[*ent.DecisionReqDefQuery, predicate.DecisionReqDef]{typ: ent.TypeDecisionReqDef, tq: q}, nil
	case *ent.DeploymentQuery:
		return &query[*ent.DeploymentQuery, predicate.Deployment]{typ: ent.TypeDeployment, tq: q}, nil
	case *ent.IdentityLinkQuery:
		return &query[*ent.IdentityLinkQuery, predicate.IdentityLink]{typ: ent.TypeIdentityLink, tq: q}, nil
	case *ent.ProcDefQuery:
		return &query[*ent.ProcDefQuery, predicate.ProcDef]{typ: ent.TypeProcDef, tq: q}, nil
	case *ent.ProcInstQuery:
		return &query[*ent.ProcInstQuery, predicate.ProcInst]{typ: ent.TypeProcInst, tq: q}, nil
	case *ent.TaskQuery:
		return &query[*ent.TaskQuery, predicate.Task]{typ: ent.TypeTask, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...ent.OrderFunc) T
		Where(...P) T
	}
}

func (q query[T, P]) Type() string {
	return q.typ
}

func (q query[T, P]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P]) Order(orders ...ent.OrderFunc) {
	q.tq.Order(orders...)
}

func (q query[T, P]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
