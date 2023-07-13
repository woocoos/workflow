// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/orgrole"
	"github.com/woocoos/workflow/ent/orgroleuser"
	"github.com/woocoos/workflow/ent/orguser"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
)

// OrgRoleUserUpdate is the builder for updating OrgRoleUser entities.
type OrgRoleUserUpdate struct {
	config
	hooks    []Hook
	mutation *OrgRoleUserMutation
}

// Where appends a list predicates to the OrgRoleUserUpdate builder.
func (oruu *OrgRoleUserUpdate) Where(ps ...predicate.OrgRoleUser) *OrgRoleUserUpdate {
	oruu.mutation.Where(ps...)
	return oruu
}

// SetOrgRoleID sets the "org_role_id" field.
func (oruu *OrgRoleUserUpdate) SetOrgRoleID(i int) *OrgRoleUserUpdate {
	oruu.mutation.SetOrgRoleID(i)
	return oruu
}

// SetOrgUserID sets the "org_user_id" field.
func (oruu *OrgRoleUserUpdate) SetOrgUserID(i int) *OrgRoleUserUpdate {
	oruu.mutation.SetOrgUserID(i)
	return oruu
}

// SetUserID sets the "user_id" field.
func (oruu *OrgRoleUserUpdate) SetUserID(i int) *OrgRoleUserUpdate {
	oruu.mutation.ResetUserID()
	oruu.mutation.SetUserID(i)
	return oruu
}

// AddUserID adds i to the "user_id" field.
func (oruu *OrgRoleUserUpdate) AddUserID(i int) *OrgRoleUserUpdate {
	oruu.mutation.AddUserID(i)
	return oruu
}

// SetOrgID sets the "org_id" field.
func (oruu *OrgRoleUserUpdate) SetOrgID(i int) *OrgRoleUserUpdate {
	oruu.mutation.ResetOrgID()
	oruu.mutation.SetOrgID(i)
	return oruu
}

// AddOrgID adds i to the "org_id" field.
func (oruu *OrgRoleUserUpdate) AddOrgID(i int) *OrgRoleUserUpdate {
	oruu.mutation.AddOrgID(i)
	return oruu
}

// SetOrgRole sets the "org_role" edge to the OrgRole entity.
func (oruu *OrgRoleUserUpdate) SetOrgRole(o *OrgRole) *OrgRoleUserUpdate {
	return oruu.SetOrgRoleID(o.ID)
}

// SetOrgUser sets the "org_user" edge to the OrgUser entity.
func (oruu *OrgRoleUserUpdate) SetOrgUser(o *OrgUser) *OrgRoleUserUpdate {
	return oruu.SetOrgUserID(o.ID)
}

// Mutation returns the OrgRoleUserMutation object of the builder.
func (oruu *OrgRoleUserUpdate) Mutation() *OrgRoleUserMutation {
	return oruu.mutation
}

// ClearOrgRole clears the "org_role" edge to the OrgRole entity.
func (oruu *OrgRoleUserUpdate) ClearOrgRole() *OrgRoleUserUpdate {
	oruu.mutation.ClearOrgRole()
	return oruu
}

// ClearOrgUser clears the "org_user" edge to the OrgUser entity.
func (oruu *OrgRoleUserUpdate) ClearOrgUser() *OrgRoleUserUpdate {
	oruu.mutation.ClearOrgUser()
	return oruu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oruu *OrgRoleUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oruu.sqlSave, oruu.mutation, oruu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oruu *OrgRoleUserUpdate) SaveX(ctx context.Context) int {
	affected, err := oruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oruu *OrgRoleUserUpdate) Exec(ctx context.Context) error {
	_, err := oruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruu *OrgRoleUserUpdate) ExecX(ctx context.Context) {
	if err := oruu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oruu *OrgRoleUserUpdate) check() error {
	if _, ok := oruu.mutation.OrgRoleID(); oruu.mutation.OrgRoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgRoleUser.org_role"`)
	}
	if _, ok := oruu.mutation.OrgUserID(); oruu.mutation.OrgUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgRoleUser.org_user"`)
	}
	return nil
}

func (oruu *OrgRoleUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oruu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgroleuser.Table, orgroleuser.Columns, sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt))
	if ps := oruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oruu.mutation.UserID(); ok {
		_spec.SetField(orgroleuser.FieldUserID, field.TypeInt, value)
	}
	if value, ok := oruu.mutation.AddedUserID(); ok {
		_spec.AddField(orgroleuser.FieldUserID, field.TypeInt, value)
	}
	if value, ok := oruu.mutation.OrgID(); ok {
		_spec.SetField(orgroleuser.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oruu.mutation.AddedOrgID(); ok {
		_spec.AddField(orgroleuser.FieldOrgID, field.TypeInt, value)
	}
	if oruu.mutation.OrgRoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgRoleTable,
			Columns: []string{orgroleuser.OrgRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruu.schemaConfig.OrgRoleUser
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruu.mutation.OrgRoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgRoleTable,
			Columns: []string{orgroleuser.OrgRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruu.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oruu.mutation.OrgUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgUserTable,
			Columns: []string{orgroleuser.OrgUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruu.schemaConfig.OrgRoleUser
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruu.mutation.OrgUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgUserTable,
			Columns: []string{orgroleuser.OrgUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruu.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = oruu.schemaConfig.OrgRoleUser
	ctx = internal.NewSchemaConfigContext(ctx, oruu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, oruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgroleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oruu.mutation.done = true
	return n, nil
}

// OrgRoleUserUpdateOne is the builder for updating a single OrgRoleUser entity.
type OrgRoleUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgRoleUserMutation
}

// SetOrgRoleID sets the "org_role_id" field.
func (oruuo *OrgRoleUserUpdateOne) SetOrgRoleID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.SetOrgRoleID(i)
	return oruuo
}

// SetOrgUserID sets the "org_user_id" field.
func (oruuo *OrgRoleUserUpdateOne) SetOrgUserID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.SetOrgUserID(i)
	return oruuo
}

// SetUserID sets the "user_id" field.
func (oruuo *OrgRoleUserUpdateOne) SetUserID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.ResetUserID()
	oruuo.mutation.SetUserID(i)
	return oruuo
}

// AddUserID adds i to the "user_id" field.
func (oruuo *OrgRoleUserUpdateOne) AddUserID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.AddUserID(i)
	return oruuo
}

// SetOrgID sets the "org_id" field.
func (oruuo *OrgRoleUserUpdateOne) SetOrgID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.ResetOrgID()
	oruuo.mutation.SetOrgID(i)
	return oruuo
}

// AddOrgID adds i to the "org_id" field.
func (oruuo *OrgRoleUserUpdateOne) AddOrgID(i int) *OrgRoleUserUpdateOne {
	oruuo.mutation.AddOrgID(i)
	return oruuo
}

// SetOrgRole sets the "org_role" edge to the OrgRole entity.
func (oruuo *OrgRoleUserUpdateOne) SetOrgRole(o *OrgRole) *OrgRoleUserUpdateOne {
	return oruuo.SetOrgRoleID(o.ID)
}

// SetOrgUser sets the "org_user" edge to the OrgUser entity.
func (oruuo *OrgRoleUserUpdateOne) SetOrgUser(o *OrgUser) *OrgRoleUserUpdateOne {
	return oruuo.SetOrgUserID(o.ID)
}

// Mutation returns the OrgRoleUserMutation object of the builder.
func (oruuo *OrgRoleUserUpdateOne) Mutation() *OrgRoleUserMutation {
	return oruuo.mutation
}

// ClearOrgRole clears the "org_role" edge to the OrgRole entity.
func (oruuo *OrgRoleUserUpdateOne) ClearOrgRole() *OrgRoleUserUpdateOne {
	oruuo.mutation.ClearOrgRole()
	return oruuo
}

// ClearOrgUser clears the "org_user" edge to the OrgUser entity.
func (oruuo *OrgRoleUserUpdateOne) ClearOrgUser() *OrgRoleUserUpdateOne {
	oruuo.mutation.ClearOrgUser()
	return oruuo
}

// Where appends a list predicates to the OrgRoleUserUpdate builder.
func (oruuo *OrgRoleUserUpdateOne) Where(ps ...predicate.OrgRoleUser) *OrgRoleUserUpdateOne {
	oruuo.mutation.Where(ps...)
	return oruuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oruuo *OrgRoleUserUpdateOne) Select(field string, fields ...string) *OrgRoleUserUpdateOne {
	oruuo.fields = append([]string{field}, fields...)
	return oruuo
}

// Save executes the query and returns the updated OrgRoleUser entity.
func (oruuo *OrgRoleUserUpdateOne) Save(ctx context.Context) (*OrgRoleUser, error) {
	return withHooks(ctx, oruuo.sqlSave, oruuo.mutation, oruuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oruuo *OrgRoleUserUpdateOne) SaveX(ctx context.Context) *OrgRoleUser {
	node, err := oruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oruuo *OrgRoleUserUpdateOne) Exec(ctx context.Context) error {
	_, err := oruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruuo *OrgRoleUserUpdateOne) ExecX(ctx context.Context) {
	if err := oruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oruuo *OrgRoleUserUpdateOne) check() error {
	if _, ok := oruuo.mutation.OrgRoleID(); oruuo.mutation.OrgRoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgRoleUser.org_role"`)
	}
	if _, ok := oruuo.mutation.OrgUserID(); oruuo.mutation.OrgUserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgRoleUser.org_user"`)
	}
	return nil
}

func (oruuo *OrgRoleUserUpdateOne) sqlSave(ctx context.Context) (_node *OrgRoleUser, err error) {
	if err := oruuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgroleuser.Table, orgroleuser.Columns, sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt))
	id, ok := oruuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgRoleUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgroleuser.FieldID)
		for _, f := range fields {
			if !orgroleuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orgroleuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oruuo.mutation.UserID(); ok {
		_spec.SetField(orgroleuser.FieldUserID, field.TypeInt, value)
	}
	if value, ok := oruuo.mutation.AddedUserID(); ok {
		_spec.AddField(orgroleuser.FieldUserID, field.TypeInt, value)
	}
	if value, ok := oruuo.mutation.OrgID(); ok {
		_spec.SetField(orgroleuser.FieldOrgID, field.TypeInt, value)
	}
	if value, ok := oruuo.mutation.AddedOrgID(); ok {
		_spec.AddField(orgroleuser.FieldOrgID, field.TypeInt, value)
	}
	if oruuo.mutation.OrgRoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgRoleTable,
			Columns: []string{orgroleuser.OrgRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruuo.schemaConfig.OrgRoleUser
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruuo.mutation.OrgRoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgRoleTable,
			Columns: []string{orgroleuser.OrgRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruuo.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oruuo.mutation.OrgUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgUserTable,
			Columns: []string{orgroleuser.OrgUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruuo.schemaConfig.OrgRoleUser
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruuo.mutation.OrgUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgUserTable,
			Columns: []string{orgroleuser.OrgUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt),
			},
		}
		edge.Schema = oruuo.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = oruuo.schemaConfig.OrgRoleUser
	ctx = internal.NewSchemaConfigContext(ctx, oruuo.schemaConfig)
	_node = &OrgRoleUser{config: oruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgroleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oruuo.mutation.done = true
	return _node, nil
}
