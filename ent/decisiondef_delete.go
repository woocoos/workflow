// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/predicate"
)

// DecisionDefDelete is the builder for deleting a DecisionDef entity.
type DecisionDefDelete struct {
	config
	hooks    []Hook
	mutation *DecisionDefMutation
}

// Where appends a list predicates to the DecisionDefDelete builder.
func (ddd *DecisionDefDelete) Where(ps ...predicate.DecisionDef) *DecisionDefDelete {
	ddd.mutation.Where(ps...)
	return ddd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ddd *DecisionDefDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DecisionDefMutation](ctx, ddd.sqlExec, ddd.mutation, ddd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ddd *DecisionDefDelete) ExecX(ctx context.Context) int {
	n, err := ddd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ddd *DecisionDefDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(decisiondef.Table, sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt))
	if ps := ddd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ddd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ddd.mutation.done = true
	return affected, err
}

// DecisionDefDeleteOne is the builder for deleting a single DecisionDef entity.
type DecisionDefDeleteOne struct {
	ddd *DecisionDefDelete
}

// Where appends a list predicates to the DecisionDefDelete builder.
func (dddo *DecisionDefDeleteOne) Where(ps ...predicate.DecisionDef) *DecisionDefDeleteOne {
	dddo.ddd.mutation.Where(ps...)
	return dddo
}

// Exec executes the deletion query.
func (dddo *DecisionDefDeleteOne) Exec(ctx context.Context) error {
	n, err := dddo.ddd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{decisiondef.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dddo *DecisionDefDeleteOne) ExecX(ctx context.Context) {
	if err := dddo.Exec(ctx); err != nil {
		panic(err)
	}
}