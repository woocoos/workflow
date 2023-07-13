// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/orgapp"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
)

// OrgAppUpdate is the builder for updating OrgApp entities.
type OrgAppUpdate struct {
	config
	hooks    []Hook
	mutation *OrgAppMutation
}

// Where appends a list predicates to the OrgAppUpdate builder.
func (oau *OrgAppUpdate) Where(ps ...predicate.OrgApp) *OrgAppUpdate {
	oau.mutation.Where(ps...)
	return oau
}

// SetOrgID sets the "org_id" field.
func (oau *OrgAppUpdate) SetOrgID(i int) *OrgAppUpdate {
	oau.mutation.ResetOrgID()
	oau.mutation.SetOrgID(i)
	return oau
}

// AddOrgID adds i to the "org_id" field.
func (oau *OrgAppUpdate) AddOrgID(i int) *OrgAppUpdate {
	oau.mutation.AddOrgID(i)
	return oau
}

// SetAppID sets the "app_id" field.
func (oau *OrgAppUpdate) SetAppID(i int) *OrgAppUpdate {
	oau.mutation.ResetAppID()
	oau.mutation.SetAppID(i)
	return oau
}

// AddAppID adds i to the "app_id" field.
func (oau *OrgAppUpdate) AddAppID(i int) *OrgAppUpdate {
	oau.mutation.AddAppID(i)
	return oau
}

// Mutation returns the OrgAppMutation object of the builder.
func (oau *OrgAppUpdate) Mutation() *OrgAppMutation {
	return oau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oau *OrgAppUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oau.sqlSave, oau.mutation, oau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oau *OrgAppUpdate) SaveX(ctx context.Context) int {
	affected, err := oau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oau *OrgAppUpdate) Exec(ctx context.Context) error {
	_, err := oau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oau *OrgAppUpdate) ExecX(ctx context.Context) {
	if err := oau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oau *OrgAppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orgapp.Table, orgapp.Columns, sqlgraph.NewFieldSpec(orgapp.FieldID, field.TypeInt))
	if ps := oau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oau.mutation.OrgID(); ok {
		_spec.SetField(orgapp.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oau.mutation.AddedOrgID(); ok {
		_spec.AddField(orgapp.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oau.mutation.AppID(); ok {
		_spec.SetField(orgapp.FieldAppID, field.TypeInt, value)
	}
	if value, ok := oau.mutation.AddedAppID(); ok {
		_spec.AddField(orgapp.FieldAppID, field.TypeInt, value)
	}
	_spec.Node.Schema = oau.schemaConfig.OrgApp
	ctx = internal.NewSchemaConfigContext(ctx, oau.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, oau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oau.mutation.done = true
	return n, nil
}

// OrgAppUpdateOne is the builder for updating a single OrgApp entity.
type OrgAppUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgAppMutation
}

// SetOrgID sets the "org_id" field.
func (oauo *OrgAppUpdateOne) SetOrgID(i int) *OrgAppUpdateOne {
	oauo.mutation.ResetOrgID()
	oauo.mutation.SetOrgID(i)
	return oauo
}

// AddOrgID adds i to the "org_id" field.
func (oauo *OrgAppUpdateOne) AddOrgID(i int) *OrgAppUpdateOne {
	oauo.mutation.AddOrgID(i)
	return oauo
}

// SetAppID sets the "app_id" field.
func (oauo *OrgAppUpdateOne) SetAppID(i int) *OrgAppUpdateOne {
	oauo.mutation.ResetAppID()
	oauo.mutation.SetAppID(i)
	return oauo
}

// AddAppID adds i to the "app_id" field.
func (oauo *OrgAppUpdateOne) AddAppID(i int) *OrgAppUpdateOne {
	oauo.mutation.AddAppID(i)
	return oauo
}

// Mutation returns the OrgAppMutation object of the builder.
func (oauo *OrgAppUpdateOne) Mutation() *OrgAppMutation {
	return oauo.mutation
}

// Where appends a list predicates to the OrgAppUpdate builder.
func (oauo *OrgAppUpdateOne) Where(ps ...predicate.OrgApp) *OrgAppUpdateOne {
	oauo.mutation.Where(ps...)
	return oauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oauo *OrgAppUpdateOne) Select(field string, fields ...string) *OrgAppUpdateOne {
	oauo.fields = append([]string{field}, fields...)
	return oauo
}

// Save executes the query and returns the updated OrgApp entity.
func (oauo *OrgAppUpdateOne) Save(ctx context.Context) (*OrgApp, error) {
	return withHooks(ctx, oauo.sqlSave, oauo.mutation, oauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oauo *OrgAppUpdateOne) SaveX(ctx context.Context) *OrgApp {
	node, err := oauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oauo *OrgAppUpdateOne) Exec(ctx context.Context) error {
	_, err := oauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oauo *OrgAppUpdateOne) ExecX(ctx context.Context) {
	if err := oauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oauo *OrgAppUpdateOne) sqlSave(ctx context.Context) (_node *OrgApp, err error) {
	_spec := sqlgraph.NewUpdateSpec(orgapp.Table, orgapp.Columns, sqlgraph.NewFieldSpec(orgapp.FieldID, field.TypeInt))
	id, ok := oauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgApp.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgapp.FieldID)
		for _, f := range fields {
			if !orgapp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orgapp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oauo.mutation.OrgID(); ok {
		_spec.SetField(orgapp.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oauo.mutation.AddedOrgID(); ok {
		_spec.AddField(orgapp.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oauo.mutation.AppID(); ok {
		_spec.SetField(orgapp.FieldAppID, field.TypeInt, value)
	}
	if value, ok := oauo.mutation.AddedAppID(); ok {
		_spec.AddField(orgapp.FieldAppID, field.TypeInt, value)
	}
	_spec.Node.Schema = oauo.schemaConfig.OrgApp
	ctx = internal.NewSchemaConfigContext(ctx, oauo.schemaConfig)
	_node = &OrgApp{config: oauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgapp.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oauo.mutation.done = true
	return _node, nil
}
