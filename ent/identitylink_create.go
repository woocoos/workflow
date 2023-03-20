// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/workflow/ent/identitylink"
	"github.com/woocoos/workflow/ent/task"
)

// IdentityLinkCreate is the builder for creating a IdentityLink entity.
type IdentityLinkCreate struct {
	config
	mutation *IdentityLinkMutation
	hooks    []Hook
}

// SetTaskID sets the "task_id" field.
func (ilc *IdentityLinkCreate) SetTaskID(i int) *IdentityLinkCreate {
	ilc.mutation.SetTaskID(i)
	return ilc
}

// SetProcDefID sets the "proc_def_id" field.
func (ilc *IdentityLinkCreate) SetProcDefID(i int) *IdentityLinkCreate {
	ilc.mutation.SetProcDefID(i)
	return ilc
}

// SetGroupID sets the "group_id" field.
func (ilc *IdentityLinkCreate) SetGroupID(i int) *IdentityLinkCreate {
	ilc.mutation.SetGroupID(i)
	return ilc
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableGroupID(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetGroupID(*i)
	}
	return ilc
}

// SetUserID sets the "user_id" field.
func (ilc *IdentityLinkCreate) SetUserID(i int) *IdentityLinkCreate {
	ilc.mutation.SetUserID(i)
	return ilc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableUserID(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetUserID(*i)
	}
	return ilc
}

// SetAssignerID sets the "assigner_id" field.
func (ilc *IdentityLinkCreate) SetAssignerID(i int) *IdentityLinkCreate {
	ilc.mutation.SetAssignerID(i)
	return ilc
}

// SetNillableAssignerID sets the "assigner_id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableAssignerID(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetAssignerID(*i)
	}
	return ilc
}

// SetLinkType sets the "link_type" field.
func (ilc *IdentityLinkCreate) SetLinkType(it identitylink.LinkType) *IdentityLinkCreate {
	ilc.mutation.SetLinkType(it)
	return ilc
}

// SetOrgID sets the "org_id" field.
func (ilc *IdentityLinkCreate) SetOrgID(i int) *IdentityLinkCreate {
	ilc.mutation.SetOrgID(i)
	return ilc
}

// SetOperationType sets the "operation_type" field.
func (ilc *IdentityLinkCreate) SetOperationType(it identitylink.OperationType) *IdentityLinkCreate {
	ilc.mutation.SetOperationType(it)
	return ilc
}

// SetComments sets the "comments" field.
func (ilc *IdentityLinkCreate) SetComments(s string) *IdentityLinkCreate {
	ilc.mutation.SetComments(s)
	return ilc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableComments(s *string) *IdentityLinkCreate {
	if s != nil {
		ilc.SetComments(*s)
	}
	return ilc
}

// SetID sets the "id" field.
func (ilc *IdentityLinkCreate) SetID(i int) *IdentityLinkCreate {
	ilc.mutation.SetID(i)
	return ilc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ilc *IdentityLinkCreate) SetNillableID(i *int) *IdentityLinkCreate {
	if i != nil {
		ilc.SetID(*i)
	}
	return ilc
}

// SetTask sets the "task" edge to the Task entity.
func (ilc *IdentityLinkCreate) SetTask(t *Task) *IdentityLinkCreate {
	return ilc.SetTaskID(t.ID)
}

// Mutation returns the IdentityLinkMutation object of the builder.
func (ilc *IdentityLinkCreate) Mutation() *IdentityLinkMutation {
	return ilc.mutation
}

// Save creates the IdentityLink in the database.
func (ilc *IdentityLinkCreate) Save(ctx context.Context) (*IdentityLink, error) {
	ilc.defaults()
	return withHooks[*IdentityLink, IdentityLinkMutation](ctx, ilc.sqlSave, ilc.mutation, ilc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ilc *IdentityLinkCreate) SaveX(ctx context.Context) *IdentityLink {
	v, err := ilc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilc *IdentityLinkCreate) Exec(ctx context.Context) error {
	_, err := ilc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilc *IdentityLinkCreate) ExecX(ctx context.Context) {
	if err := ilc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ilc *IdentityLinkCreate) defaults() {
	if _, ok := ilc.mutation.ID(); !ok {
		v := identitylink.DefaultID()
		ilc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ilc *IdentityLinkCreate) check() error {
	if _, ok := ilc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`ent: missing required field "IdentityLink.task_id"`)}
	}
	if _, ok := ilc.mutation.ProcDefID(); !ok {
		return &ValidationError{Name: "proc_def_id", err: errors.New(`ent: missing required field "IdentityLink.proc_def_id"`)}
	}
	if _, ok := ilc.mutation.LinkType(); !ok {
		return &ValidationError{Name: "link_type", err: errors.New(`ent: missing required field "IdentityLink.link_type"`)}
	}
	if v, ok := ilc.mutation.LinkType(); ok {
		if err := identitylink.LinkTypeValidator(v); err != nil {
			return &ValidationError{Name: "link_type", err: fmt.Errorf(`ent: validator failed for field "IdentityLink.link_type": %w`, err)}
		}
	}
	if _, ok := ilc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "IdentityLink.org_id"`)}
	}
	if _, ok := ilc.mutation.OperationType(); !ok {
		return &ValidationError{Name: "operation_type", err: errors.New(`ent: missing required field "IdentityLink.operation_type"`)}
	}
	if v, ok := ilc.mutation.OperationType(); ok {
		if err := identitylink.OperationTypeValidator(v); err != nil {
			return &ValidationError{Name: "operation_type", err: fmt.Errorf(`ent: validator failed for field "IdentityLink.operation_type": %w`, err)}
		}
	}
	if _, ok := ilc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "IdentityLink.task"`)}
	}
	return nil
}

func (ilc *IdentityLinkCreate) sqlSave(ctx context.Context) (*IdentityLink, error) {
	if err := ilc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ilc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ilc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ilc.mutation.id = &_node.ID
	ilc.mutation.done = true
	return _node, nil
}

func (ilc *IdentityLinkCreate) createSpec() (*IdentityLink, *sqlgraph.CreateSpec) {
	var (
		_node = &IdentityLink{config: ilc.config}
		_spec = sqlgraph.NewCreateSpec(identitylink.Table, sqlgraph.NewFieldSpec(identitylink.FieldID, field.TypeInt))
	)
	if id, ok := ilc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ilc.mutation.ProcDefID(); ok {
		_spec.SetField(identitylink.FieldProcDefID, field.TypeInt, value)
		_node.ProcDefID = value
	}
	if value, ok := ilc.mutation.GroupID(); ok {
		_spec.SetField(identitylink.FieldGroupID, field.TypeInt, value)
		_node.GroupID = value
	}
	if value, ok := ilc.mutation.UserID(); ok {
		_spec.SetField(identitylink.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := ilc.mutation.AssignerID(); ok {
		_spec.SetField(identitylink.FieldAssignerID, field.TypeInt, value)
		_node.AssignerID = value
	}
	if value, ok := ilc.mutation.LinkType(); ok {
		_spec.SetField(identitylink.FieldLinkType, field.TypeEnum, value)
		_node.LinkType = value
	}
	if value, ok := ilc.mutation.OrgID(); ok {
		_spec.SetField(identitylink.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
	}
	if value, ok := ilc.mutation.OperationType(); ok {
		_spec.SetField(identitylink.FieldOperationType, field.TypeEnum, value)
		_node.OperationType = value
	}
	if value, ok := ilc.mutation.Comments(); ok {
		_spec.SetField(identitylink.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := ilc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identitylink.TaskTable,
			Columns: []string{identitylink.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TaskID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IdentityLinkCreateBulk is the builder for creating many IdentityLink entities in bulk.
type IdentityLinkCreateBulk struct {
	config
	builders []*IdentityLinkCreate
}

// Save creates the IdentityLink entities in the database.
func (ilcb *IdentityLinkCreateBulk) Save(ctx context.Context) ([]*IdentityLink, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ilcb.builders))
	nodes := make([]*IdentityLink, len(ilcb.builders))
	mutators := make([]Mutator, len(ilcb.builders))
	for i := range ilcb.builders {
		func(i int, root context.Context) {
			builder := ilcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentityLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, ilcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ilcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ilcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ilcb *IdentityLinkCreateBulk) SaveX(ctx context.Context) []*IdentityLink {
	v, err := ilcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ilcb *IdentityLinkCreateBulk) Exec(ctx context.Context) error {
	_, err := ilcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ilcb *IdentityLinkCreateBulk) ExecX(ctx context.Context) {
	if err := ilcb.Exec(ctx); err != nil {
		panic(err)
	}
}