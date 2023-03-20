// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/deployment"
)

// DecisionReqDefCreate is the builder for creating a DecisionReqDef entity.
type DecisionReqDefCreate struct {
	config
	mutation *DecisionReqDefMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (drdc *DecisionReqDefCreate) SetCreatedBy(i int) *DecisionReqDefCreate {
	drdc.mutation.SetCreatedBy(i)
	return drdc
}

// SetCreatedAt sets the "created_at" field.
func (drdc *DecisionReqDefCreate) SetCreatedAt(t time.Time) *DecisionReqDefCreate {
	drdc.mutation.SetCreatedAt(t)
	return drdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableCreatedAt(t *time.Time) *DecisionReqDefCreate {
	if t != nil {
		drdc.SetCreatedAt(*t)
	}
	return drdc
}

// SetUpdatedBy sets the "updated_by" field.
func (drdc *DecisionReqDefCreate) SetUpdatedBy(i int) *DecisionReqDefCreate {
	drdc.mutation.SetUpdatedBy(i)
	return drdc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableUpdatedBy(i *int) *DecisionReqDefCreate {
	if i != nil {
		drdc.SetUpdatedBy(*i)
	}
	return drdc
}

// SetUpdatedAt sets the "updated_at" field.
func (drdc *DecisionReqDefCreate) SetUpdatedAt(t time.Time) *DecisionReqDefCreate {
	drdc.mutation.SetUpdatedAt(t)
	return drdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableUpdatedAt(t *time.Time) *DecisionReqDefCreate {
	if t != nil {
		drdc.SetUpdatedAt(*t)
	}
	return drdc
}

// SetDeploymentID sets the "deployment_id" field.
func (drdc *DecisionReqDefCreate) SetDeploymentID(i int) *DecisionReqDefCreate {
	drdc.mutation.SetDeploymentID(i)
	return drdc
}

// SetOrgID sets the "org_id" field.
func (drdc *DecisionReqDefCreate) SetOrgID(i int) *DecisionReqDefCreate {
	drdc.mutation.SetOrgID(i)
	return drdc
}

// SetAppID sets the "app_id" field.
func (drdc *DecisionReqDefCreate) SetAppID(i int) *DecisionReqDefCreate {
	drdc.mutation.SetAppID(i)
	return drdc
}

// SetCategory sets the "category" field.
func (drdc *DecisionReqDefCreate) SetCategory(s string) *DecisionReqDefCreate {
	drdc.mutation.SetCategory(s)
	return drdc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableCategory(s *string) *DecisionReqDefCreate {
	if s != nil {
		drdc.SetCategory(*s)
	}
	return drdc
}

// SetName sets the "name" field.
func (drdc *DecisionReqDefCreate) SetName(s string) *DecisionReqDefCreate {
	drdc.mutation.SetName(s)
	return drdc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableName(s *string) *DecisionReqDefCreate {
	if s != nil {
		drdc.SetName(*s)
	}
	return drdc
}

// SetKey sets the "key" field.
func (drdc *DecisionReqDefCreate) SetKey(s string) *DecisionReqDefCreate {
	drdc.mutation.SetKey(s)
	return drdc
}

// SetVersion sets the "version" field.
func (drdc *DecisionReqDefCreate) SetVersion(i int32) *DecisionReqDefCreate {
	drdc.mutation.SetVersion(i)
	return drdc
}

// SetRevision sets the "revision" field.
func (drdc *DecisionReqDefCreate) SetRevision(i int32) *DecisionReqDefCreate {
	drdc.mutation.SetRevision(i)
	return drdc
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableRevision(i *int32) *DecisionReqDefCreate {
	if i != nil {
		drdc.SetRevision(*i)
	}
	return drdc
}

// SetResourceName sets the "resource_name" field.
func (drdc *DecisionReqDefCreate) SetResourceName(s string) *DecisionReqDefCreate {
	drdc.mutation.SetResourceName(s)
	return drdc
}

// SetNillableResourceName sets the "resource_name" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableResourceName(s *string) *DecisionReqDefCreate {
	if s != nil {
		drdc.SetResourceName(*s)
	}
	return drdc
}

// SetDgrmResourceName sets the "dgrm_resource_name" field.
func (drdc *DecisionReqDefCreate) SetDgrmResourceName(s string) *DecisionReqDefCreate {
	drdc.mutation.SetDgrmResourceName(s)
	return drdc
}

// SetNillableDgrmResourceName sets the "dgrm_resource_name" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableDgrmResourceName(s *string) *DecisionReqDefCreate {
	if s != nil {
		drdc.SetDgrmResourceName(*s)
	}
	return drdc
}

// SetResourceData sets the "resource_data" field.
func (drdc *DecisionReqDefCreate) SetResourceData(b []byte) *DecisionReqDefCreate {
	drdc.mutation.SetResourceData(b)
	return drdc
}

// SetID sets the "id" field.
func (drdc *DecisionReqDefCreate) SetID(i int) *DecisionReqDefCreate {
	drdc.mutation.SetID(i)
	return drdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (drdc *DecisionReqDefCreate) SetNillableID(i *int) *DecisionReqDefCreate {
	if i != nil {
		drdc.SetID(*i)
	}
	return drdc
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (drdc *DecisionReqDefCreate) SetDeployment(d *Deployment) *DecisionReqDefCreate {
	return drdc.SetDeploymentID(d.ID)
}

// AddDecisionDefIDs adds the "decision_defs" edge to the DecisionDef entity by IDs.
func (drdc *DecisionReqDefCreate) AddDecisionDefIDs(ids ...int) *DecisionReqDefCreate {
	drdc.mutation.AddDecisionDefIDs(ids...)
	return drdc
}

// AddDecisionDefs adds the "decision_defs" edges to the DecisionDef entity.
func (drdc *DecisionReqDefCreate) AddDecisionDefs(d ...*DecisionDef) *DecisionReqDefCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return drdc.AddDecisionDefIDs(ids...)
}

// Mutation returns the DecisionReqDefMutation object of the builder.
func (drdc *DecisionReqDefCreate) Mutation() *DecisionReqDefMutation {
	return drdc.mutation
}

// Save creates the DecisionReqDef in the database.
func (drdc *DecisionReqDefCreate) Save(ctx context.Context) (*DecisionReqDef, error) {
	if err := drdc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*DecisionReqDef, DecisionReqDefMutation](ctx, drdc.sqlSave, drdc.mutation, drdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (drdc *DecisionReqDefCreate) SaveX(ctx context.Context) *DecisionReqDef {
	v, err := drdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drdc *DecisionReqDefCreate) Exec(ctx context.Context) error {
	_, err := drdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drdc *DecisionReqDefCreate) ExecX(ctx context.Context) {
	if err := drdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (drdc *DecisionReqDefCreate) defaults() error {
	if _, ok := drdc.mutation.CreatedAt(); !ok {
		if decisionreqdef.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized decisionreqdef.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := decisionreqdef.DefaultCreatedAt()
		drdc.mutation.SetCreatedAt(v)
	}
	if _, ok := drdc.mutation.ID(); !ok {
		if decisionreqdef.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized decisionreqdef.DefaultID (forgotten import ent/runtime?)")
		}
		v := decisionreqdef.DefaultID()
		drdc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (drdc *DecisionReqDefCreate) check() error {
	if _, ok := drdc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "DecisionReqDef.created_by"`)}
	}
	if _, ok := drdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DecisionReqDef.created_at"`)}
	}
	if _, ok := drdc.mutation.DeploymentID(); !ok {
		return &ValidationError{Name: "deployment_id", err: errors.New(`ent: missing required field "DecisionReqDef.deployment_id"`)}
	}
	if _, ok := drdc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "DecisionReqDef.org_id"`)}
	}
	if _, ok := drdc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "DecisionReqDef.app_id"`)}
	}
	if _, ok := drdc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "DecisionReqDef.key"`)}
	}
	if _, ok := drdc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "DecisionReqDef.version"`)}
	}
	if _, ok := drdc.mutation.DeploymentID(); !ok {
		return &ValidationError{Name: "deployment", err: errors.New(`ent: missing required edge "DecisionReqDef.deployment"`)}
	}
	return nil
}

func (drdc *DecisionReqDefCreate) sqlSave(ctx context.Context) (*DecisionReqDef, error) {
	if err := drdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := drdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	drdc.mutation.id = &_node.ID
	drdc.mutation.done = true
	return _node, nil
}

func (drdc *DecisionReqDefCreate) createSpec() (*DecisionReqDef, *sqlgraph.CreateSpec) {
	var (
		_node = &DecisionReqDef{config: drdc.config}
		_spec = sqlgraph.NewCreateSpec(decisionreqdef.Table, sqlgraph.NewFieldSpec(decisionreqdef.FieldID, field.TypeInt))
	)
	if id, ok := drdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := drdc.mutation.CreatedBy(); ok {
		_spec.SetField(decisionreqdef.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := drdc.mutation.CreatedAt(); ok {
		_spec.SetField(decisionreqdef.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := drdc.mutation.UpdatedBy(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := drdc.mutation.UpdatedAt(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := drdc.mutation.OrgID(); ok {
		_spec.SetField(decisionreqdef.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
	}
	if value, ok := drdc.mutation.AppID(); ok {
		_spec.SetField(decisionreqdef.FieldAppID, field.TypeInt, value)
		_node.AppID = value
	}
	if value, ok := drdc.mutation.Category(); ok {
		_spec.SetField(decisionreqdef.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := drdc.mutation.Name(); ok {
		_spec.SetField(decisionreqdef.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := drdc.mutation.Key(); ok {
		_spec.SetField(decisionreqdef.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := drdc.mutation.Version(); ok {
		_spec.SetField(decisionreqdef.FieldVersion, field.TypeInt32, value)
		_node.Version = value
	}
	if value, ok := drdc.mutation.Revision(); ok {
		_spec.SetField(decisionreqdef.FieldRevision, field.TypeInt32, value)
		_node.Revision = value
	}
	if value, ok := drdc.mutation.ResourceName(); ok {
		_spec.SetField(decisionreqdef.FieldResourceName, field.TypeString, value)
		_node.ResourceName = value
	}
	if value, ok := drdc.mutation.DgrmResourceName(); ok {
		_spec.SetField(decisionreqdef.FieldDgrmResourceName, field.TypeString, value)
		_node.DgrmResourceName = value
	}
	if value, ok := drdc.mutation.ResourceData(); ok {
		_spec.SetField(decisionreqdef.FieldResourceData, field.TypeBytes, value)
		_node.ResourceData = value
	}
	if nodes := drdc.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   decisionreqdef.DeploymentTable,
			Columns: []string{decisionreqdef.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deployment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DeploymentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drdc.mutation.DecisionDefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: decisiondef.FieldID,
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

// DecisionReqDefCreateBulk is the builder for creating many DecisionReqDef entities in bulk.
type DecisionReqDefCreateBulk struct {
	config
	builders []*DecisionReqDefCreate
}

// Save creates the DecisionReqDef entities in the database.
func (drdcb *DecisionReqDefCreateBulk) Save(ctx context.Context) ([]*DecisionReqDef, error) {
	specs := make([]*sqlgraph.CreateSpec, len(drdcb.builders))
	nodes := make([]*DecisionReqDef, len(drdcb.builders))
	mutators := make([]Mutator, len(drdcb.builders))
	for i := range drdcb.builders {
		func(i int, root context.Context) {
			builder := drdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DecisionReqDefMutation)
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
					_, err = mutators[i+1].Mutate(root, drdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, drdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, drdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (drdcb *DecisionReqDefCreateBulk) SaveX(ctx context.Context) []*DecisionReqDef {
	v, err := drdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drdcb *DecisionReqDefCreateBulk) Exec(ctx context.Context) error {
	_, err := drdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drdcb *DecisionReqDefCreateBulk) ExecX(ctx context.Context) {
	if err := drdcb.Exec(ctx); err != nil {
		panic(err)
	}
}