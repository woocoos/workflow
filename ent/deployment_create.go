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
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/deployment"
	"github.com/woocoos/workflow/ent/procdef"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (dc *DeploymentCreate) SetCreatedBy(i int) *DeploymentCreate {
	dc.mutation.SetCreatedBy(i)
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeploymentCreate) SetCreatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableCreatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedBy sets the "updated_by" field.
func (dc *DeploymentCreate) SetUpdatedBy(i int) *DeploymentCreate {
	dc.mutation.SetUpdatedBy(i)
	return dc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdatedBy(i *int) *DeploymentCreate {
	if i != nil {
		dc.SetUpdatedBy(*i)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeploymentCreate) SetUpdatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetTenantID sets the "tenant_id" field.
func (dc *DeploymentCreate) SetTenantID(i int) *DeploymentCreate {
	dc.mutation.SetTenantID(i)
	return dc
}

// SetAppID sets the "app_id" field.
func (dc *DeploymentCreate) SetAppID(i int) *DeploymentCreate {
	dc.mutation.SetAppID(i)
	return dc
}

// SetName sets the "name" field.
func (dc *DeploymentCreate) SetName(s string) *DeploymentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableName(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetSource sets the "source" field.
func (dc *DeploymentCreate) SetSource(s string) *DeploymentCreate {
	dc.mutation.SetSource(s)
	return dc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableSource(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetSource(*s)
	}
	return dc
}

// SetDeployTime sets the "deploy_time" field.
func (dc *DeploymentCreate) SetDeployTime(t time.Time) *DeploymentCreate {
	dc.mutation.SetDeployTime(t)
	return dc
}

// SetNillableDeployTime sets the "deploy_time" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableDeployTime(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetDeployTime(*t)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DeploymentCreate) SetID(i int) *DeploymentCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableID(i *int) *DeploymentCreate {
	if i != nil {
		dc.SetID(*i)
	}
	return dc
}

// AddProcDefIDs adds the "proc_defs" edge to the ProcDef entity by IDs.
func (dc *DeploymentCreate) AddProcDefIDs(ids ...int) *DeploymentCreate {
	dc.mutation.AddProcDefIDs(ids...)
	return dc
}

// AddProcDefs adds the "proc_defs" edges to the ProcDef entity.
func (dc *DeploymentCreate) AddProcDefs(p ...*ProcDef) *DeploymentCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddProcDefIDs(ids...)
}

// AddDecisionReqIDs adds the "decision_reqs" edge to the DecisionReqDef entity by IDs.
func (dc *DeploymentCreate) AddDecisionReqIDs(ids ...int) *DeploymentCreate {
	dc.mutation.AddDecisionReqIDs(ids...)
	return dc
}

// AddDecisionReqs adds the "decision_reqs" edges to the DecisionReqDef entity.
func (dc *DeploymentCreate) AddDecisionReqs(d ...*DecisionReqDef) *DeploymentCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDecisionReqIDs(ids...)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if deployment.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized deployment.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := deployment.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.DeployTime(); !ok {
		if deployment.DefaultDeployTime == nil {
			return fmt.Errorf("ent: uninitialized deployment.DefaultDeployTime (forgotten import ent/runtime?)")
		}
		v := deployment.DefaultDeployTime()
		dc.mutation.SetDeployTime(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		if deployment.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized deployment.DefaultID (forgotten import ent/runtime?)")
		}
		v := deployment.DefaultID()
		dc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Deployment.created_by"`)}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Deployment.created_at"`)}
	}
	if _, ok := dc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "Deployment.tenant_id"`)}
	}
	if _, ok := dc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Deployment.app_id"`)}
	}
	if _, ok := dc.mutation.DeployTime(); !ok {
		return &ValidationError{Name: "deploy_time", err: errors.New(`ent: missing required field "Deployment.deploy_time"`)}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(deployment.Table, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeInt))
	)
	_spec.Schema = dc.schemaConfig.Deployment
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedBy(); ok {
		_spec.SetField(deployment.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(deployment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedBy(); ok {
		_spec.SetField(deployment.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(deployment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.TenantID(); ok {
		_spec.SetField(deployment.FieldTenantID, field.TypeInt, value)
		_node.TenantID = value
	}
	if value, ok := dc.mutation.AppID(); ok {
		_spec.SetField(deployment.FieldAppID, field.TypeInt, value)
		_node.AppID = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Source(); ok {
		_spec.SetField(deployment.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := dc.mutation.DeployTime(); ok {
		_spec.SetField(deployment.FieldDeployTime, field.TypeTime, value)
		_node.DeployTime = value
	}
	if nodes := dc.mutation.ProcDefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.ProcDefsTable,
			Columns: []string{deployment.ProcDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = dc.schemaConfig.ProcDef
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DecisionReqsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deployment.DecisionReqsTable,
			Columns: []string{deployment.DecisionReqsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisionreqdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = dc.schemaConfig.DecisionReqDef
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
//	client.Deployment.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeploymentUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dc *DeploymentCreate) OnConflict(opts ...sql.ConflictOption) *DeploymentUpsertOne {
	dc.conflict = opts
	return &DeploymentUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DeploymentCreate) OnConflictColumns(columns ...string) *DeploymentUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DeploymentUpsertOne{
		create: dc,
	}
}

type (
	// DeploymentUpsertOne is the builder for "upsert"-ing
	//  one Deployment node.
	DeploymentUpsertOne struct {
		create *DeploymentCreate
	}

	// DeploymentUpsert is the "OnConflict" setter.
	DeploymentUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *DeploymentUpsert) SetUpdatedBy(v int) *DeploymentUpsert {
	u.Set(deployment.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateUpdatedBy() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeploymentUpsert) AddUpdatedBy(v int) *DeploymentUpsert {
	u.Add(deployment.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *DeploymentUpsert) ClearUpdatedBy() *DeploymentUpsert {
	u.SetNull(deployment.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeploymentUpsert) SetUpdatedAt(v time.Time) *DeploymentUpsert {
	u.Set(deployment.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateUpdatedAt() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeploymentUpsert) ClearUpdatedAt() *DeploymentUpsert {
	u.SetNull(deployment.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *DeploymentUpsert) SetName(v string) *DeploymentUpsert {
	u.Set(deployment.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateName() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *DeploymentUpsert) ClearName() *DeploymentUpsert {
	u.SetNull(deployment.FieldName)
	return u
}

// SetSource sets the "source" field.
func (u *DeploymentUpsert) SetSource(v string) *DeploymentUpsert {
	u.Set(deployment.FieldSource, v)
	return u
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *DeploymentUpsert) UpdateSource() *DeploymentUpsert {
	u.SetExcluded(deployment.FieldSource)
	return u
}

// ClearSource clears the value of the "source" field.
func (u *DeploymentUpsert) ClearSource() *DeploymentUpsert {
	u.SetNull(deployment.FieldSource)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deployment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeploymentUpsertOne) UpdateNewValues() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deployment.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(deployment.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(deployment.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.TenantID(); exists {
			s.SetIgnore(deployment.FieldTenantID)
		}
		if _, exists := u.create.mutation.AppID(); exists {
			s.SetIgnore(deployment.FieldAppID)
		}
		if _, exists := u.create.mutation.DeployTime(); exists {
			s.SetIgnore(deployment.FieldDeployTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeploymentUpsertOne) Ignore() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeploymentUpsertOne) DoNothing() *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeploymentCreate.OnConflict
// documentation for more info.
func (u *DeploymentUpsertOne) Update(set func(*DeploymentUpsert)) *DeploymentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeploymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeploymentUpsertOne) SetUpdatedBy(v int) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeploymentUpsertOne) AddUpdatedBy(v int) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateUpdatedBy() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *DeploymentUpsertOne) ClearUpdatedBy() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeploymentUpsertOne) SetUpdatedAt(v time.Time) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateUpdatedAt() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeploymentUpsertOne) ClearUpdatedAt() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *DeploymentUpsertOne) SetName(v string) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateName() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeploymentUpsertOne) ClearName() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearName()
	})
}

// SetSource sets the "source" field.
func (u *DeploymentUpsertOne) SetSource(v string) *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *DeploymentUpsertOne) UpdateSource() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateSource()
	})
}

// ClearSource clears the value of the "source" field.
func (u *DeploymentUpsertOne) ClearSource() *DeploymentUpsertOne {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearSource()
	})
}

// Exec executes the query.
func (u *DeploymentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeploymentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeploymentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeploymentUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeploymentUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	builders []*DeploymentCreate
	conflict []sql.ConflictOption
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Deployment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeploymentUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (dcb *DeploymentCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeploymentUpsertBulk {
	dcb.conflict = opts
	return &DeploymentUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DeploymentCreateBulk) OnConflictColumns(columns ...string) *DeploymentUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DeploymentUpsertBulk{
		create: dcb,
	}
}

// DeploymentUpsertBulk is the builder for "upsert"-ing
// a bulk of Deployment nodes.
type DeploymentUpsertBulk struct {
	create *DeploymentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deployment.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeploymentUpsertBulk) UpdateNewValues() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deployment.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(deployment.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(deployment.FieldCreatedAt)
			}
			if _, exists := b.mutation.TenantID(); exists {
				s.SetIgnore(deployment.FieldTenantID)
			}
			if _, exists := b.mutation.AppID(); exists {
				s.SetIgnore(deployment.FieldAppID)
			}
			if _, exists := b.mutation.DeployTime(); exists {
				s.SetIgnore(deployment.FieldDeployTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Deployment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeploymentUpsertBulk) Ignore() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeploymentUpsertBulk) DoNothing() *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeploymentCreateBulk.OnConflict
// documentation for more info.
func (u *DeploymentUpsertBulk) Update(set func(*DeploymentUpsert)) *DeploymentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeploymentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *DeploymentUpsertBulk) SetUpdatedBy(v int) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *DeploymentUpsertBulk) AddUpdatedBy(v int) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateUpdatedBy() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *DeploymentUpsertBulk) ClearUpdatedBy() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeploymentUpsertBulk) SetUpdatedAt(v time.Time) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateUpdatedAt() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *DeploymentUpsertBulk) ClearUpdatedAt() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *DeploymentUpsertBulk) SetName(v string) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateName() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeploymentUpsertBulk) ClearName() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearName()
	})
}

// SetSource sets the "source" field.
func (u *DeploymentUpsertBulk) SetSource(v string) *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.SetSource(v)
	})
}

// UpdateSource sets the "source" field to the value that was provided on create.
func (u *DeploymentUpsertBulk) UpdateSource() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.UpdateSource()
	})
}

// ClearSource clears the value of the "source" field.
func (u *DeploymentUpsertBulk) ClearSource() *DeploymentUpsertBulk {
	return u.Update(func(s *DeploymentUpsert) {
		s.ClearSource()
	})
}

// Exec executes the query.
func (u *DeploymentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeploymentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeploymentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeploymentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
