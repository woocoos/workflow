// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
	"github.com/woocoos/workflow/ent/orgapp"
)

// OrgAppDelete is the builder for deleting a OrgApp entity.
type OrgAppDelete struct {
	config
	hooks    []Hook
	mutation *OrgAppMutation
}

// Where appends a list predicates to the OrgAppDelete builder.
func (oad *OrgAppDelete) Where(ps ...predicate.OrgApp) *OrgAppDelete {
	oad.mutation.Where(ps...)
	return oad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oad *OrgAppDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, oad.sqlExec, oad.mutation, oad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (oad *OrgAppDelete) ExecX(ctx context.Context) int {
	n, err := oad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oad *OrgAppDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgapp.Table, sqlgraph.NewFieldSpec(orgapp.FieldID, field.TypeInt))
	_spec.Node.Schema = oad.schemaConfig.OrgApp
	ctx = internal.NewSchemaConfigContext(ctx, oad.schemaConfig)
	if ps := oad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	oad.mutation.done = true
	return affected, err
}

// OrgAppDeleteOne is the builder for deleting a single OrgApp entity.
type OrgAppDeleteOne struct {
	oad *OrgAppDelete
}

// Where appends a list predicates to the OrgAppDelete builder.
func (oado *OrgAppDeleteOne) Where(ps ...predicate.OrgApp) *OrgAppDeleteOne {
	oado.oad.mutation.Where(ps...)
	return oado
}

// Exec executes the deletion query.
func (oado *OrgAppDeleteOne) Exec(ctx context.Context) error {
	n, err := oado.oad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgapp.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oado *OrgAppDeleteOne) ExecX(ctx context.Context) {
	if err := oado.Exec(ctx); err != nil {
		panic(err)
	}
}
