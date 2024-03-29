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
	"github.com/woocoos/workflow/ent/decisiondef"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
)

// DecisionReqDefUpdate is the builder for updating DecisionReqDef entities.
type DecisionReqDefUpdate struct {
	config
	hooks    []Hook
	mutation *DecisionReqDefMutation
}

// Where appends a list predicates to the DecisionReqDefUpdate builder.
func (drdu *DecisionReqDefUpdate) Where(ps ...predicate.DecisionReqDef) *DecisionReqDefUpdate {
	drdu.mutation.Where(ps...)
	return drdu
}

// SetUpdatedBy sets the "updated_by" field.
func (drdu *DecisionReqDefUpdate) SetUpdatedBy(i int) *DecisionReqDefUpdate {
	drdu.mutation.ResetUpdatedBy()
	drdu.mutation.SetUpdatedBy(i)
	return drdu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableUpdatedBy(i *int) *DecisionReqDefUpdate {
	if i != nil {
		drdu.SetUpdatedBy(*i)
	}
	return drdu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (drdu *DecisionReqDefUpdate) AddUpdatedBy(i int) *DecisionReqDefUpdate {
	drdu.mutation.AddUpdatedBy(i)
	return drdu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (drdu *DecisionReqDefUpdate) ClearUpdatedBy() *DecisionReqDefUpdate {
	drdu.mutation.ClearUpdatedBy()
	return drdu
}

// SetUpdatedAt sets the "updated_at" field.
func (drdu *DecisionReqDefUpdate) SetUpdatedAt(t time.Time) *DecisionReqDefUpdate {
	drdu.mutation.SetUpdatedAt(t)
	return drdu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableUpdatedAt(t *time.Time) *DecisionReqDefUpdate {
	if t != nil {
		drdu.SetUpdatedAt(*t)
	}
	return drdu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (drdu *DecisionReqDefUpdate) ClearUpdatedAt() *DecisionReqDefUpdate {
	drdu.mutation.ClearUpdatedAt()
	return drdu
}

// SetCategory sets the "category" field.
func (drdu *DecisionReqDefUpdate) SetCategory(s string) *DecisionReqDefUpdate {
	drdu.mutation.SetCategory(s)
	return drdu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableCategory(s *string) *DecisionReqDefUpdate {
	if s != nil {
		drdu.SetCategory(*s)
	}
	return drdu
}

// ClearCategory clears the value of the "category" field.
func (drdu *DecisionReqDefUpdate) ClearCategory() *DecisionReqDefUpdate {
	drdu.mutation.ClearCategory()
	return drdu
}

// SetName sets the "name" field.
func (drdu *DecisionReqDefUpdate) SetName(s string) *DecisionReqDefUpdate {
	drdu.mutation.SetName(s)
	return drdu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableName(s *string) *DecisionReqDefUpdate {
	if s != nil {
		drdu.SetName(*s)
	}
	return drdu
}

// ClearName clears the value of the "name" field.
func (drdu *DecisionReqDefUpdate) ClearName() *DecisionReqDefUpdate {
	drdu.mutation.ClearName()
	return drdu
}

// SetKey sets the "key" field.
func (drdu *DecisionReqDefUpdate) SetKey(s string) *DecisionReqDefUpdate {
	drdu.mutation.SetKey(s)
	return drdu
}

// SetVersion sets the "version" field.
func (drdu *DecisionReqDefUpdate) SetVersion(i int32) *DecisionReqDefUpdate {
	drdu.mutation.ResetVersion()
	drdu.mutation.SetVersion(i)
	return drdu
}

// AddVersion adds i to the "version" field.
func (drdu *DecisionReqDefUpdate) AddVersion(i int32) *DecisionReqDefUpdate {
	drdu.mutation.AddVersion(i)
	return drdu
}

// SetRevision sets the "revision" field.
func (drdu *DecisionReqDefUpdate) SetRevision(i int32) *DecisionReqDefUpdate {
	drdu.mutation.ResetRevision()
	drdu.mutation.SetRevision(i)
	return drdu
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableRevision(i *int32) *DecisionReqDefUpdate {
	if i != nil {
		drdu.SetRevision(*i)
	}
	return drdu
}

// AddRevision adds i to the "revision" field.
func (drdu *DecisionReqDefUpdate) AddRevision(i int32) *DecisionReqDefUpdate {
	drdu.mutation.AddRevision(i)
	return drdu
}

// ClearRevision clears the value of the "revision" field.
func (drdu *DecisionReqDefUpdate) ClearRevision() *DecisionReqDefUpdate {
	drdu.mutation.ClearRevision()
	return drdu
}

// SetResourceKey sets the "resource_key" field.
func (drdu *DecisionReqDefUpdate) SetResourceKey(s string) *DecisionReqDefUpdate {
	drdu.mutation.SetResourceKey(s)
	return drdu
}

// SetNillableResourceKey sets the "resource_key" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableResourceKey(s *string) *DecisionReqDefUpdate {
	if s != nil {
		drdu.SetResourceKey(*s)
	}
	return drdu
}

// ClearResourceKey clears the value of the "resource_key" field.
func (drdu *DecisionReqDefUpdate) ClearResourceKey() *DecisionReqDefUpdate {
	drdu.mutation.ClearResourceKey()
	return drdu
}

// SetResourceID sets the "resource_id" field.
func (drdu *DecisionReqDefUpdate) SetResourceID(i int) *DecisionReqDefUpdate {
	drdu.mutation.ResetResourceID()
	drdu.mutation.SetResourceID(i)
	return drdu
}

// SetNillableResourceID sets the "resource_id" field if the given value is not nil.
func (drdu *DecisionReqDefUpdate) SetNillableResourceID(i *int) *DecisionReqDefUpdate {
	if i != nil {
		drdu.SetResourceID(*i)
	}
	return drdu
}

// AddResourceID adds i to the "resource_id" field.
func (drdu *DecisionReqDefUpdate) AddResourceID(i int) *DecisionReqDefUpdate {
	drdu.mutation.AddResourceID(i)
	return drdu
}

// ClearResourceID clears the value of the "resource_id" field.
func (drdu *DecisionReqDefUpdate) ClearResourceID() *DecisionReqDefUpdate {
	drdu.mutation.ClearResourceID()
	return drdu
}

// AddDecisionDefIDs adds the "decision_defs" edge to the DecisionDef entity by IDs.
func (drdu *DecisionReqDefUpdate) AddDecisionDefIDs(ids ...int) *DecisionReqDefUpdate {
	drdu.mutation.AddDecisionDefIDs(ids...)
	return drdu
}

// AddDecisionDefs adds the "decision_defs" edges to the DecisionDef entity.
func (drdu *DecisionReqDefUpdate) AddDecisionDefs(d ...*DecisionDef) *DecisionReqDefUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return drdu.AddDecisionDefIDs(ids...)
}

// Mutation returns the DecisionReqDefMutation object of the builder.
func (drdu *DecisionReqDefUpdate) Mutation() *DecisionReqDefMutation {
	return drdu.mutation
}

// ClearDecisionDefs clears all "decision_defs" edges to the DecisionDef entity.
func (drdu *DecisionReqDefUpdate) ClearDecisionDefs() *DecisionReqDefUpdate {
	drdu.mutation.ClearDecisionDefs()
	return drdu
}

// RemoveDecisionDefIDs removes the "decision_defs" edge to DecisionDef entities by IDs.
func (drdu *DecisionReqDefUpdate) RemoveDecisionDefIDs(ids ...int) *DecisionReqDefUpdate {
	drdu.mutation.RemoveDecisionDefIDs(ids...)
	return drdu
}

// RemoveDecisionDefs removes "decision_defs" edges to DecisionDef entities.
func (drdu *DecisionReqDefUpdate) RemoveDecisionDefs(d ...*DecisionDef) *DecisionReqDefUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return drdu.RemoveDecisionDefIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (drdu *DecisionReqDefUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, drdu.sqlSave, drdu.mutation, drdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (drdu *DecisionReqDefUpdate) SaveX(ctx context.Context) int {
	affected, err := drdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (drdu *DecisionReqDefUpdate) Exec(ctx context.Context) error {
	_, err := drdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drdu *DecisionReqDefUpdate) ExecX(ctx context.Context) {
	if err := drdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drdu *DecisionReqDefUpdate) check() error {
	if _, ok := drdu.mutation.DeploymentID(); drdu.mutation.DeploymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DecisionReqDef.deployment"`)
	}
	return nil
}

func (drdu *DecisionReqDefUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := drdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(decisionreqdef.Table, decisionreqdef.Columns, sqlgraph.NewFieldSpec(decisionreqdef.FieldID, field.TypeInt))
	if ps := drdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := drdu.mutation.UpdatedBy(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := drdu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(decisionreqdef.FieldUpdatedBy, field.TypeInt, value)
	}
	if drdu.mutation.UpdatedByCleared() {
		_spec.ClearField(decisionreqdef.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := drdu.mutation.UpdatedAt(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedAt, field.TypeTime, value)
	}
	if drdu.mutation.UpdatedAtCleared() {
		_spec.ClearField(decisionreqdef.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := drdu.mutation.Category(); ok {
		_spec.SetField(decisionreqdef.FieldCategory, field.TypeString, value)
	}
	if drdu.mutation.CategoryCleared() {
		_spec.ClearField(decisionreqdef.FieldCategory, field.TypeString)
	}
	if value, ok := drdu.mutation.Name(); ok {
		_spec.SetField(decisionreqdef.FieldName, field.TypeString, value)
	}
	if drdu.mutation.NameCleared() {
		_spec.ClearField(decisionreqdef.FieldName, field.TypeString)
	}
	if value, ok := drdu.mutation.Key(); ok {
		_spec.SetField(decisionreqdef.FieldKey, field.TypeString, value)
	}
	if value, ok := drdu.mutation.Version(); ok {
		_spec.SetField(decisionreqdef.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := drdu.mutation.AddedVersion(); ok {
		_spec.AddField(decisionreqdef.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := drdu.mutation.Revision(); ok {
		_spec.SetField(decisionreqdef.FieldRevision, field.TypeInt32, value)
	}
	if value, ok := drdu.mutation.AddedRevision(); ok {
		_spec.AddField(decisionreqdef.FieldRevision, field.TypeInt32, value)
	}
	if drdu.mutation.RevisionCleared() {
		_spec.ClearField(decisionreqdef.FieldRevision, field.TypeInt32)
	}
	if value, ok := drdu.mutation.ResourceKey(); ok {
		_spec.SetField(decisionreqdef.FieldResourceKey, field.TypeString, value)
	}
	if drdu.mutation.ResourceKeyCleared() {
		_spec.ClearField(decisionreqdef.FieldResourceKey, field.TypeString)
	}
	if value, ok := drdu.mutation.ResourceID(); ok {
		_spec.SetField(decisionreqdef.FieldResourceID, field.TypeInt, value)
	}
	if value, ok := drdu.mutation.AddedResourceID(); ok {
		_spec.AddField(decisionreqdef.FieldResourceID, field.TypeInt, value)
	}
	if drdu.mutation.ResourceIDCleared() {
		_spec.ClearField(decisionreqdef.FieldResourceID, field.TypeInt)
	}
	if drdu.mutation.DecisionDefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drdu.schemaConfig.DecisionDef
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drdu.mutation.RemovedDecisionDefsIDs(); len(nodes) > 0 && !drdu.mutation.DecisionDefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drdu.schemaConfig.DecisionDef
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drdu.mutation.DecisionDefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drdu.schemaConfig.DecisionDef
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = drdu.schemaConfig.DecisionReqDef
	ctx = internal.NewSchemaConfigContext(ctx, drdu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, drdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{decisionreqdef.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	drdu.mutation.done = true
	return n, nil
}

// DecisionReqDefUpdateOne is the builder for updating a single DecisionReqDef entity.
type DecisionReqDefUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DecisionReqDefMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (drduo *DecisionReqDefUpdateOne) SetUpdatedBy(i int) *DecisionReqDefUpdateOne {
	drduo.mutation.ResetUpdatedBy()
	drduo.mutation.SetUpdatedBy(i)
	return drduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableUpdatedBy(i *int) *DecisionReqDefUpdateOne {
	if i != nil {
		drduo.SetUpdatedBy(*i)
	}
	return drduo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (drduo *DecisionReqDefUpdateOne) AddUpdatedBy(i int) *DecisionReqDefUpdateOne {
	drduo.mutation.AddUpdatedBy(i)
	return drduo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (drduo *DecisionReqDefUpdateOne) ClearUpdatedBy() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearUpdatedBy()
	return drduo
}

// SetUpdatedAt sets the "updated_at" field.
func (drduo *DecisionReqDefUpdateOne) SetUpdatedAt(t time.Time) *DecisionReqDefUpdateOne {
	drduo.mutation.SetUpdatedAt(t)
	return drduo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableUpdatedAt(t *time.Time) *DecisionReqDefUpdateOne {
	if t != nil {
		drduo.SetUpdatedAt(*t)
	}
	return drduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (drduo *DecisionReqDefUpdateOne) ClearUpdatedAt() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearUpdatedAt()
	return drduo
}

// SetCategory sets the "category" field.
func (drduo *DecisionReqDefUpdateOne) SetCategory(s string) *DecisionReqDefUpdateOne {
	drduo.mutation.SetCategory(s)
	return drduo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableCategory(s *string) *DecisionReqDefUpdateOne {
	if s != nil {
		drduo.SetCategory(*s)
	}
	return drduo
}

// ClearCategory clears the value of the "category" field.
func (drduo *DecisionReqDefUpdateOne) ClearCategory() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearCategory()
	return drduo
}

// SetName sets the "name" field.
func (drduo *DecisionReqDefUpdateOne) SetName(s string) *DecisionReqDefUpdateOne {
	drduo.mutation.SetName(s)
	return drduo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableName(s *string) *DecisionReqDefUpdateOne {
	if s != nil {
		drduo.SetName(*s)
	}
	return drduo
}

// ClearName clears the value of the "name" field.
func (drduo *DecisionReqDefUpdateOne) ClearName() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearName()
	return drduo
}

// SetKey sets the "key" field.
func (drduo *DecisionReqDefUpdateOne) SetKey(s string) *DecisionReqDefUpdateOne {
	drduo.mutation.SetKey(s)
	return drduo
}

// SetVersion sets the "version" field.
func (drduo *DecisionReqDefUpdateOne) SetVersion(i int32) *DecisionReqDefUpdateOne {
	drduo.mutation.ResetVersion()
	drduo.mutation.SetVersion(i)
	return drduo
}

// AddVersion adds i to the "version" field.
func (drduo *DecisionReqDefUpdateOne) AddVersion(i int32) *DecisionReqDefUpdateOne {
	drduo.mutation.AddVersion(i)
	return drduo
}

// SetRevision sets the "revision" field.
func (drduo *DecisionReqDefUpdateOne) SetRevision(i int32) *DecisionReqDefUpdateOne {
	drduo.mutation.ResetRevision()
	drduo.mutation.SetRevision(i)
	return drduo
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableRevision(i *int32) *DecisionReqDefUpdateOne {
	if i != nil {
		drduo.SetRevision(*i)
	}
	return drduo
}

// AddRevision adds i to the "revision" field.
func (drduo *DecisionReqDefUpdateOne) AddRevision(i int32) *DecisionReqDefUpdateOne {
	drduo.mutation.AddRevision(i)
	return drduo
}

// ClearRevision clears the value of the "revision" field.
func (drduo *DecisionReqDefUpdateOne) ClearRevision() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearRevision()
	return drduo
}

// SetResourceKey sets the "resource_key" field.
func (drduo *DecisionReqDefUpdateOne) SetResourceKey(s string) *DecisionReqDefUpdateOne {
	drduo.mutation.SetResourceKey(s)
	return drduo
}

// SetNillableResourceKey sets the "resource_key" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableResourceKey(s *string) *DecisionReqDefUpdateOne {
	if s != nil {
		drduo.SetResourceKey(*s)
	}
	return drduo
}

// ClearResourceKey clears the value of the "resource_key" field.
func (drduo *DecisionReqDefUpdateOne) ClearResourceKey() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearResourceKey()
	return drduo
}

// SetResourceID sets the "resource_id" field.
func (drduo *DecisionReqDefUpdateOne) SetResourceID(i int) *DecisionReqDefUpdateOne {
	drduo.mutation.ResetResourceID()
	drduo.mutation.SetResourceID(i)
	return drduo
}

// SetNillableResourceID sets the "resource_id" field if the given value is not nil.
func (drduo *DecisionReqDefUpdateOne) SetNillableResourceID(i *int) *DecisionReqDefUpdateOne {
	if i != nil {
		drduo.SetResourceID(*i)
	}
	return drduo
}

// AddResourceID adds i to the "resource_id" field.
func (drduo *DecisionReqDefUpdateOne) AddResourceID(i int) *DecisionReqDefUpdateOne {
	drduo.mutation.AddResourceID(i)
	return drduo
}

// ClearResourceID clears the value of the "resource_id" field.
func (drduo *DecisionReqDefUpdateOne) ClearResourceID() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearResourceID()
	return drduo
}

// AddDecisionDefIDs adds the "decision_defs" edge to the DecisionDef entity by IDs.
func (drduo *DecisionReqDefUpdateOne) AddDecisionDefIDs(ids ...int) *DecisionReqDefUpdateOne {
	drduo.mutation.AddDecisionDefIDs(ids...)
	return drduo
}

// AddDecisionDefs adds the "decision_defs" edges to the DecisionDef entity.
func (drduo *DecisionReqDefUpdateOne) AddDecisionDefs(d ...*DecisionDef) *DecisionReqDefUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return drduo.AddDecisionDefIDs(ids...)
}

// Mutation returns the DecisionReqDefMutation object of the builder.
func (drduo *DecisionReqDefUpdateOne) Mutation() *DecisionReqDefMutation {
	return drduo.mutation
}

// ClearDecisionDefs clears all "decision_defs" edges to the DecisionDef entity.
func (drduo *DecisionReqDefUpdateOne) ClearDecisionDefs() *DecisionReqDefUpdateOne {
	drduo.mutation.ClearDecisionDefs()
	return drduo
}

// RemoveDecisionDefIDs removes the "decision_defs" edge to DecisionDef entities by IDs.
func (drduo *DecisionReqDefUpdateOne) RemoveDecisionDefIDs(ids ...int) *DecisionReqDefUpdateOne {
	drduo.mutation.RemoveDecisionDefIDs(ids...)
	return drduo
}

// RemoveDecisionDefs removes "decision_defs" edges to DecisionDef entities.
func (drduo *DecisionReqDefUpdateOne) RemoveDecisionDefs(d ...*DecisionDef) *DecisionReqDefUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return drduo.RemoveDecisionDefIDs(ids...)
}

// Where appends a list predicates to the DecisionReqDefUpdate builder.
func (drduo *DecisionReqDefUpdateOne) Where(ps ...predicate.DecisionReqDef) *DecisionReqDefUpdateOne {
	drduo.mutation.Where(ps...)
	return drduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (drduo *DecisionReqDefUpdateOne) Select(field string, fields ...string) *DecisionReqDefUpdateOne {
	drduo.fields = append([]string{field}, fields...)
	return drduo
}

// Save executes the query and returns the updated DecisionReqDef entity.
func (drduo *DecisionReqDefUpdateOne) Save(ctx context.Context) (*DecisionReqDef, error) {
	return withHooks(ctx, drduo.sqlSave, drduo.mutation, drduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (drduo *DecisionReqDefUpdateOne) SaveX(ctx context.Context) *DecisionReqDef {
	node, err := drduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (drduo *DecisionReqDefUpdateOne) Exec(ctx context.Context) error {
	_, err := drduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drduo *DecisionReqDefUpdateOne) ExecX(ctx context.Context) {
	if err := drduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drduo *DecisionReqDefUpdateOne) check() error {
	if _, ok := drduo.mutation.DeploymentID(); drduo.mutation.DeploymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "DecisionReqDef.deployment"`)
	}
	return nil
}

func (drduo *DecisionReqDefUpdateOne) sqlSave(ctx context.Context) (_node *DecisionReqDef, err error) {
	if err := drduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(decisionreqdef.Table, decisionreqdef.Columns, sqlgraph.NewFieldSpec(decisionreqdef.FieldID, field.TypeInt))
	id, ok := drduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DecisionReqDef.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := drduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, decisionreqdef.FieldID)
		for _, f := range fields {
			if !decisionreqdef.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != decisionreqdef.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := drduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := drduo.mutation.UpdatedBy(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := drduo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(decisionreqdef.FieldUpdatedBy, field.TypeInt, value)
	}
	if drduo.mutation.UpdatedByCleared() {
		_spec.ClearField(decisionreqdef.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := drduo.mutation.UpdatedAt(); ok {
		_spec.SetField(decisionreqdef.FieldUpdatedAt, field.TypeTime, value)
	}
	if drduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(decisionreqdef.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := drduo.mutation.Category(); ok {
		_spec.SetField(decisionreqdef.FieldCategory, field.TypeString, value)
	}
	if drduo.mutation.CategoryCleared() {
		_spec.ClearField(decisionreqdef.FieldCategory, field.TypeString)
	}
	if value, ok := drduo.mutation.Name(); ok {
		_spec.SetField(decisionreqdef.FieldName, field.TypeString, value)
	}
	if drduo.mutation.NameCleared() {
		_spec.ClearField(decisionreqdef.FieldName, field.TypeString)
	}
	if value, ok := drduo.mutation.Key(); ok {
		_spec.SetField(decisionreqdef.FieldKey, field.TypeString, value)
	}
	if value, ok := drduo.mutation.Version(); ok {
		_spec.SetField(decisionreqdef.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := drduo.mutation.AddedVersion(); ok {
		_spec.AddField(decisionreqdef.FieldVersion, field.TypeInt32, value)
	}
	if value, ok := drduo.mutation.Revision(); ok {
		_spec.SetField(decisionreqdef.FieldRevision, field.TypeInt32, value)
	}
	if value, ok := drduo.mutation.AddedRevision(); ok {
		_spec.AddField(decisionreqdef.FieldRevision, field.TypeInt32, value)
	}
	if drduo.mutation.RevisionCleared() {
		_spec.ClearField(decisionreqdef.FieldRevision, field.TypeInt32)
	}
	if value, ok := drduo.mutation.ResourceKey(); ok {
		_spec.SetField(decisionreqdef.FieldResourceKey, field.TypeString, value)
	}
	if drduo.mutation.ResourceKeyCleared() {
		_spec.ClearField(decisionreqdef.FieldResourceKey, field.TypeString)
	}
	if value, ok := drduo.mutation.ResourceID(); ok {
		_spec.SetField(decisionreqdef.FieldResourceID, field.TypeInt, value)
	}
	if value, ok := drduo.mutation.AddedResourceID(); ok {
		_spec.AddField(decisionreqdef.FieldResourceID, field.TypeInt, value)
	}
	if drduo.mutation.ResourceIDCleared() {
		_spec.ClearField(decisionreqdef.FieldResourceID, field.TypeInt)
	}
	if drduo.mutation.DecisionDefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drduo.schemaConfig.DecisionDef
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drduo.mutation.RemovedDecisionDefsIDs(); len(nodes) > 0 && !drduo.mutation.DecisionDefsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drduo.schemaConfig.DecisionDef
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := drduo.mutation.DecisionDefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   decisionreqdef.DecisionDefsTable,
			Columns: []string{decisionreqdef.DecisionDefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(decisiondef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = drduo.schemaConfig.DecisionDef
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = drduo.schemaConfig.DecisionReqDef
	ctx = internal.NewSchemaConfigContext(ctx, drduo.schemaConfig)
	_node = &DecisionReqDef{config: drduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, drduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{decisionreqdef.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	drduo.mutation.done = true
	return _node, nil
}
