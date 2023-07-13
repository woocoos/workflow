// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/orgrole"
	"github.com/woocoos/workflow/ent/orgroleuser"
	"github.com/woocoos/workflow/ent/orguser"
)

// OrgUserCreate is the builder for creating a OrgUser entity.
type OrgUserCreate struct {
	config
	mutation *OrgUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetOrgID sets the "org_id" field.
func (ouc *OrgUserCreate) SetOrgID(i int) *OrgUserCreate {
	ouc.mutation.SetOrgID(i)
	return ouc
}

// SetUserID sets the "user_id" field.
func (ouc *OrgUserCreate) SetUserID(i int) *OrgUserCreate {
	ouc.mutation.SetUserID(i)
	return ouc
}

// SetJoinedAt sets the "joined_at" field.
func (ouc *OrgUserCreate) SetJoinedAt(t time.Time) *OrgUserCreate {
	ouc.mutation.SetJoinedAt(t)
	return ouc
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableJoinedAt(t *time.Time) *OrgUserCreate {
	if t != nil {
		ouc.SetJoinedAt(*t)
	}
	return ouc
}

// SetDisplayName sets the "display_name" field.
func (ouc *OrgUserCreate) SetDisplayName(s string) *OrgUserCreate {
	ouc.mutation.SetDisplayName(s)
	return ouc
}

// SetID sets the "id" field.
func (ouc *OrgUserCreate) SetID(i int) *OrgUserCreate {
	ouc.mutation.SetID(i)
	return ouc
}

// AddOrgRoleIDs adds the "org_roles" edge to the OrgRole entity by IDs.
func (ouc *OrgUserCreate) AddOrgRoleIDs(ids ...int) *OrgUserCreate {
	ouc.mutation.AddOrgRoleIDs(ids...)
	return ouc
}

// AddOrgRoles adds the "org_roles" edges to the OrgRole entity.
func (ouc *OrgUserCreate) AddOrgRoles(o ...*OrgRole) *OrgUserCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddOrgRoleIDs(ids...)
}

// AddOrgRoleUserIDs adds the "org_role_user" edge to the OrgRoleUser entity by IDs.
func (ouc *OrgUserCreate) AddOrgRoleUserIDs(ids ...int) *OrgUserCreate {
	ouc.mutation.AddOrgRoleUserIDs(ids...)
	return ouc
}

// AddOrgRoleUser adds the "org_role_user" edges to the OrgRoleUser entity.
func (ouc *OrgUserCreate) AddOrgRoleUser(o ...*OrgRoleUser) *OrgUserCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddOrgRoleUserIDs(ids...)
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouc *OrgUserCreate) Mutation() *OrgUserMutation {
	return ouc.mutation
}

// Save creates the OrgUser in the database.
func (ouc *OrgUserCreate) Save(ctx context.Context) (*OrgUser, error) {
	ouc.defaults()
	return withHooks(ctx, ouc.sqlSave, ouc.mutation, ouc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ouc *OrgUserCreate) SaveX(ctx context.Context) *OrgUser {
	v, err := ouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ouc *OrgUserCreate) Exec(ctx context.Context) error {
	_, err := ouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouc *OrgUserCreate) ExecX(ctx context.Context) {
	if err := ouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouc *OrgUserCreate) defaults() {
	if _, ok := ouc.mutation.JoinedAt(); !ok {
		v := orguser.DefaultJoinedAt()
		ouc.mutation.SetJoinedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouc *OrgUserCreate) check() error {
	if _, ok := ouc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgUser.org_id"`)}
	}
	if _, ok := ouc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OrgUser.user_id"`)}
	}
	if _, ok := ouc.mutation.JoinedAt(); !ok {
		return &ValidationError{Name: "joined_at", err: errors.New(`ent: missing required field "OrgUser.joined_at"`)}
	}
	if _, ok := ouc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "OrgUser.display_name"`)}
	}
	return nil
}

func (ouc *OrgUserCreate) sqlSave(ctx context.Context) (*OrgUser, error) {
	if err := ouc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ouc.mutation.id = &_node.ID
	ouc.mutation.done = true
	return _node, nil
}

func (ouc *OrgUserCreate) createSpec() (*OrgUser, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgUser{config: ouc.config}
		_spec = sqlgraph.NewCreateSpec(orguser.Table, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	)
	_spec.Schema = ouc.schemaConfig.OrgUser
	_spec.OnConflict = ouc.conflict
	if id, ok := ouc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ouc.mutation.OrgID(); ok {
		_spec.SetField(orguser.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
	}
	if value, ok := ouc.mutation.UserID(); ok {
		_spec.SetField(orguser.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := ouc.mutation.JoinedAt(); ok {
		_spec.SetField(orguser.FieldJoinedAt, field.TypeTime, value)
		_node.JoinedAt = value
	}
	if value, ok := ouc.mutation.DisplayName(); ok {
		_spec.SetField(orguser.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if nodes := ouc.mutation.OrgRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		edge.Schema = ouc.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.OrgRoleUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		edge.Schema = ouc.schemaConfig.OrgRoleUser
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgUser.Create().
//		SetOrgID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUserUpsert) {
//			SetOrgID(v+v).
//		}).
//		Exec(ctx)
func (ouc *OrgUserCreate) OnConflict(opts ...sql.ConflictOption) *OrgUserUpsertOne {
	ouc.conflict = opts
	return &OrgUserUpsertOne{
		create: ouc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ouc *OrgUserCreate) OnConflictColumns(columns ...string) *OrgUserUpsertOne {
	ouc.conflict = append(ouc.conflict, sql.ConflictColumns(columns...))
	return &OrgUserUpsertOne{
		create: ouc,
	}
}

type (
	// OrgUserUpsertOne is the builder for "upsert"-ing
	//  one OrgUser node.
	OrgUserUpsertOne struct {
		create *OrgUserCreate
	}

	// OrgUserUpsert is the "OnConflict" setter.
	OrgUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetOrgID sets the "org_id" field.
func (u *OrgUserUpsert) SetOrgID(v int) *OrgUserUpsert {
	u.Set(orguser.FieldOrgID, v)
	return u
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgUserUpsert) UpdateOrgID() *OrgUserUpsert {
	u.SetExcluded(orguser.FieldOrgID)
	return u
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgUserUpsert) AddOrgID(v int) *OrgUserUpsert {
	u.Add(orguser.FieldOrgID, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrgUserUpsert) SetUserID(v int) *OrgUserUpsert {
	u.Set(orguser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserUpsert) UpdateUserID() *OrgUserUpsert {
	u.SetExcluded(orguser.FieldUserID)
	return u
}

// AddUserID adds v to the "user_id" field.
func (u *OrgUserUpsert) AddUserID(v int) *OrgUserUpsert {
	u.Add(orguser.FieldUserID, v)
	return u
}

// SetJoinedAt sets the "joined_at" field.
func (u *OrgUserUpsert) SetJoinedAt(v time.Time) *OrgUserUpsert {
	u.Set(orguser.FieldJoinedAt, v)
	return u
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *OrgUserUpsert) UpdateJoinedAt() *OrgUserUpsert {
	u.SetExcluded(orguser.FieldJoinedAt)
	return u
}

// SetDisplayName sets the "display_name" field.
func (u *OrgUserUpsert) SetDisplayName(v string) *OrgUserUpsert {
	u.Set(orguser.FieldDisplayName, v)
	return u
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *OrgUserUpsert) UpdateDisplayName() *OrgUserUpsert {
	u.SetExcluded(orguser.FieldDisplayName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orguser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgUserUpsertOne) UpdateNewValues() *OrgUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orguser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrgUserUpsertOne) Ignore() *OrgUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUserUpsertOne) DoNothing() *OrgUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUserCreate.OnConflict
// documentation for more info.
func (u *OrgUserUpsertOne) Update(set func(*OrgUserUpsert)) *OrgUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrgID sets the "org_id" field.
func (u *OrgUserUpsertOne) SetOrgID(v int) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetOrgID(v)
	})
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgUserUpsertOne) AddOrgID(v int) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.AddOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgUserUpsertOne) UpdateOrgID() *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateOrgID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgUserUpsertOne) SetUserID(v int) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrgUserUpsertOne) AddUserID(v int) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserUpsertOne) UpdateUserID() *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateUserID()
	})
}

// SetJoinedAt sets the "joined_at" field.
func (u *OrgUserUpsertOne) SetJoinedAt(v time.Time) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetJoinedAt(v)
	})
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *OrgUserUpsertOne) UpdateJoinedAt() *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateJoinedAt()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *OrgUserUpsertOne) SetDisplayName(v string) *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *OrgUserUpsertOne) UpdateDisplayName() *OrgUserUpsertOne {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateDisplayName()
	})
}

// Exec executes the query.
func (u *OrgUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgUserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgUserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgUserCreateBulk is the builder for creating many OrgUser entities in bulk.
type OrgUserCreateBulk struct {
	config
	builders []*OrgUserCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgUser entities in the database.
func (oucb *OrgUserCreateBulk) Save(ctx context.Context) ([]*OrgUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oucb.builders))
	nodes := make([]*OrgUser, len(oucb.builders))
	mutators := make([]Mutator, len(oucb.builders))
	for i := range oucb.builders {
		func(i int, root context.Context) {
			builder := oucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) SaveX(ctx context.Context) []*OrgUser {
	v, err := oucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oucb *OrgUserCreateBulk) Exec(ctx context.Context) error {
	_, err := oucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) ExecX(ctx context.Context) {
	if err := oucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUserUpsert) {
//			SetOrgID(v+v).
//		}).
//		Exec(ctx)
func (oucb *OrgUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgUserUpsertBulk {
	oucb.conflict = opts
	return &OrgUserUpsertBulk{
		create: oucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oucb *OrgUserCreateBulk) OnConflictColumns(columns ...string) *OrgUserUpsertBulk {
	oucb.conflict = append(oucb.conflict, sql.ConflictColumns(columns...))
	return &OrgUserUpsertBulk{
		create: oucb,
	}
}

// OrgUserUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgUser nodes.
type OrgUserUpsertBulk struct {
	create *OrgUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orguser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgUserUpsertBulk) UpdateNewValues() *OrgUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orguser.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrgUserUpsertBulk) Ignore() *OrgUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUserUpsertBulk) DoNothing() *OrgUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUserCreateBulk.OnConflict
// documentation for more info.
func (u *OrgUserUpsertBulk) Update(set func(*OrgUserUpsert)) *OrgUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetOrgID sets the "org_id" field.
func (u *OrgUserUpsertBulk) SetOrgID(v int) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetOrgID(v)
	})
}

// AddOrgID adds v to the "org_id" field.
func (u *OrgUserUpsertBulk) AddOrgID(v int) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.AddOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgUserUpsertBulk) UpdateOrgID() *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateOrgID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgUserUpsertBulk) SetUserID(v int) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *OrgUserUpsertBulk) AddUserID(v int) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserUpsertBulk) UpdateUserID() *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateUserID()
	})
}

// SetJoinedAt sets the "joined_at" field.
func (u *OrgUserUpsertBulk) SetJoinedAt(v time.Time) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetJoinedAt(v)
	})
}

// UpdateJoinedAt sets the "joined_at" field to the value that was provided on create.
func (u *OrgUserUpsertBulk) UpdateJoinedAt() *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateJoinedAt()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *OrgUserUpsertBulk) SetDisplayName(v string) *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *OrgUserUpsertBulk) UpdateDisplayName() *OrgUserUpsertBulk {
	return u.Update(func(s *OrgUserUpsert) {
		s.UpdateDisplayName()
	})
}

// Exec executes the query.
func (u *OrgUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrgUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
