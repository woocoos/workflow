// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

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

// SetOrgID sets the "org_id" field.
func (dc *DeploymentCreate) SetOrgID(i int) *DeploymentCreate {
	dc.mutation.SetOrgID(i)
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
	return withHooks[*Deployment, DeploymentMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
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
	if _, ok := dc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "Deployment.org_id"`)}
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
	if value, ok := dc.mutation.OrgID(); ok {
		_spec.SetField(deployment.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: procdef.FieldID,
				},
			},
		}
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: decisionreqdef.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	builders []*DeploymentCreate
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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