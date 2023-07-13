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
	"github.com/woocoos/workflow/ent/predicate"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
	"github.com/woocoos/workflow/ent/task"

	"github.com/woocoos/workflow/ent/internal"
)

// ProcInstUpdate is the builder for updating ProcInst entities.
type ProcInstUpdate struct {
	config
	hooks    []Hook
	mutation *ProcInstMutation
}

// Where appends a list predicates to the ProcInstUpdate builder.
func (piu *ProcInstUpdate) Where(ps ...predicate.ProcInst) *ProcInstUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetUpdatedBy sets the "updated_by" field.
func (piu *ProcInstUpdate) SetUpdatedBy(i int) *ProcInstUpdate {
	piu.mutation.ResetUpdatedBy()
	piu.mutation.SetUpdatedBy(i)
	return piu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableUpdatedBy(i *int) *ProcInstUpdate {
	if i != nil {
		piu.SetUpdatedBy(*i)
	}
	return piu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (piu *ProcInstUpdate) AddUpdatedBy(i int) *ProcInstUpdate {
	piu.mutation.AddUpdatedBy(i)
	return piu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (piu *ProcInstUpdate) ClearUpdatedBy() *ProcInstUpdate {
	piu.mutation.ClearUpdatedBy()
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *ProcInstUpdate) SetUpdatedAt(t time.Time) *ProcInstUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableUpdatedAt(t *time.Time) *ProcInstUpdate {
	if t != nil {
		piu.SetUpdatedAt(*t)
	}
	return piu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piu *ProcInstUpdate) ClearUpdatedAt() *ProcInstUpdate {
	piu.mutation.ClearUpdatedAt()
	return piu
}

// SetProcDefID sets the "proc_def_id" field.
func (piu *ProcInstUpdate) SetProcDefID(i int) *ProcInstUpdate {
	piu.mutation.SetProcDefID(i)
	return piu
}

// SetBusinessKey sets the "business_key" field.
func (piu *ProcInstUpdate) SetBusinessKey(s string) *ProcInstUpdate {
	piu.mutation.SetBusinessKey(s)
	return piu
}

// SetStartTime sets the "start_time" field.
func (piu *ProcInstUpdate) SetStartTime(t time.Time) *ProcInstUpdate {
	piu.mutation.SetStartTime(t)
	return piu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableStartTime(t *time.Time) *ProcInstUpdate {
	if t != nil {
		piu.SetStartTime(*t)
	}
	return piu
}

// SetEndTime sets the "end_time" field.
func (piu *ProcInstUpdate) SetEndTime(t time.Time) *ProcInstUpdate {
	piu.mutation.SetEndTime(t)
	return piu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableEndTime(t *time.Time) *ProcInstUpdate {
	if t != nil {
		piu.SetEndTime(*t)
	}
	return piu
}

// ClearEndTime clears the value of the "end_time" field.
func (piu *ProcInstUpdate) ClearEndTime() *ProcInstUpdate {
	piu.mutation.ClearEndTime()
	return piu
}

// SetDuration sets the "duration" field.
func (piu *ProcInstUpdate) SetDuration(i int) *ProcInstUpdate {
	piu.mutation.ResetDuration()
	piu.mutation.SetDuration(i)
	return piu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableDuration(i *int) *ProcInstUpdate {
	if i != nil {
		piu.SetDuration(*i)
	}
	return piu
}

// AddDuration adds i to the "duration" field.
func (piu *ProcInstUpdate) AddDuration(i int) *ProcInstUpdate {
	piu.mutation.AddDuration(i)
	return piu
}

// ClearDuration clears the value of the "duration" field.
func (piu *ProcInstUpdate) ClearDuration() *ProcInstUpdate {
	piu.mutation.ClearDuration()
	return piu
}

// SetStartUserID sets the "start_user_id" field.
func (piu *ProcInstUpdate) SetStartUserID(i int) *ProcInstUpdate {
	piu.mutation.ResetStartUserID()
	piu.mutation.SetStartUserID(i)
	return piu
}

// AddStartUserID adds i to the "start_user_id" field.
func (piu *ProcInstUpdate) AddStartUserID(i int) *ProcInstUpdate {
	piu.mutation.AddStartUserID(i)
	return piu
}

// SetSupperInstanceID sets the "supper_instance_id" field.
func (piu *ProcInstUpdate) SetSupperInstanceID(i int) *ProcInstUpdate {
	piu.mutation.ResetSupperInstanceID()
	piu.mutation.SetSupperInstanceID(i)
	return piu
}

// SetNillableSupperInstanceID sets the "supper_instance_id" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableSupperInstanceID(i *int) *ProcInstUpdate {
	if i != nil {
		piu.SetSupperInstanceID(*i)
	}
	return piu
}

// AddSupperInstanceID adds i to the "supper_instance_id" field.
func (piu *ProcInstUpdate) AddSupperInstanceID(i int) *ProcInstUpdate {
	piu.mutation.AddSupperInstanceID(i)
	return piu
}

// ClearSupperInstanceID clears the value of the "supper_instance_id" field.
func (piu *ProcInstUpdate) ClearSupperInstanceID() *ProcInstUpdate {
	piu.mutation.ClearSupperInstanceID()
	return piu
}

// SetRootInstanceID sets the "root_instance_id" field.
func (piu *ProcInstUpdate) SetRootInstanceID(i int) *ProcInstUpdate {
	piu.mutation.ResetRootInstanceID()
	piu.mutation.SetRootInstanceID(i)
	return piu
}

// SetNillableRootInstanceID sets the "root_instance_id" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableRootInstanceID(i *int) *ProcInstUpdate {
	if i != nil {
		piu.SetRootInstanceID(*i)
	}
	return piu
}

// AddRootInstanceID adds i to the "root_instance_id" field.
func (piu *ProcInstUpdate) AddRootInstanceID(i int) *ProcInstUpdate {
	piu.mutation.AddRootInstanceID(i)
	return piu
}

// ClearRootInstanceID clears the value of the "root_instance_id" field.
func (piu *ProcInstUpdate) ClearRootInstanceID() *ProcInstUpdate {
	piu.mutation.ClearRootInstanceID()
	return piu
}

// SetDeletedTime sets the "deleted_time" field.
func (piu *ProcInstUpdate) SetDeletedTime(t time.Time) *ProcInstUpdate {
	piu.mutation.SetDeletedTime(t)
	return piu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableDeletedTime(t *time.Time) *ProcInstUpdate {
	if t != nil {
		piu.SetDeletedTime(*t)
	}
	return piu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (piu *ProcInstUpdate) ClearDeletedTime() *ProcInstUpdate {
	piu.mutation.ClearDeletedTime()
	return piu
}

// SetDeletedReason sets the "deleted_reason" field.
func (piu *ProcInstUpdate) SetDeletedReason(s string) *ProcInstUpdate {
	piu.mutation.SetDeletedReason(s)
	return piu
}

// SetNillableDeletedReason sets the "deleted_reason" field if the given value is not nil.
func (piu *ProcInstUpdate) SetNillableDeletedReason(s *string) *ProcInstUpdate {
	if s != nil {
		piu.SetDeletedReason(*s)
	}
	return piu
}

// ClearDeletedReason clears the value of the "deleted_reason" field.
func (piu *ProcInstUpdate) ClearDeletedReason() *ProcInstUpdate {
	piu.mutation.ClearDeletedReason()
	return piu
}

// SetStatus sets the "status" field.
func (piu *ProcInstUpdate) SetStatus(pr procinst.Status) *ProcInstUpdate {
	piu.mutation.SetStatus(pr)
	return piu
}

// SetProcDef sets the "proc_def" edge to the ProcDef entity.
func (piu *ProcInstUpdate) SetProcDef(p *ProcDef) *ProcInstUpdate {
	return piu.SetProcDefID(p.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (piu *ProcInstUpdate) AddTaskIDs(ids ...int) *ProcInstUpdate {
	piu.mutation.AddTaskIDs(ids...)
	return piu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (piu *ProcInstUpdate) AddTasks(t ...*Task) *ProcInstUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return piu.AddTaskIDs(ids...)
}

// Mutation returns the ProcInstMutation object of the builder.
func (piu *ProcInstUpdate) Mutation() *ProcInstMutation {
	return piu.mutation
}

// ClearProcDef clears the "proc_def" edge to the ProcDef entity.
func (piu *ProcInstUpdate) ClearProcDef() *ProcInstUpdate {
	piu.mutation.ClearProcDef()
	return piu
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (piu *ProcInstUpdate) ClearTasks() *ProcInstUpdate {
	piu.mutation.ClearTasks()
	return piu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (piu *ProcInstUpdate) RemoveTaskIDs(ids ...int) *ProcInstUpdate {
	piu.mutation.RemoveTaskIDs(ids...)
	return piu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (piu *ProcInstUpdate) RemoveTasks(t ...*Task) *ProcInstUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return piu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *ProcInstUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *ProcInstUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *ProcInstUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *ProcInstUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *ProcInstUpdate) check() error {
	if v, ok := piu.mutation.Status(); ok {
		if err := procinst.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProcInst.status": %w`, err)}
		}
	}
	if _, ok := piu.mutation.ProcDefID(); piu.mutation.ProcDefCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProcInst.proc_def"`)
	}
	return nil
}

func (piu *ProcInstUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(procinst.Table, procinst.Columns, sqlgraph.NewFieldSpec(procinst.FieldID, field.TypeInt))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.UpdatedBy(); ok {
		_spec.SetField(procinst.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(procinst.FieldUpdatedBy, field.TypeInt, value)
	}
	if piu.mutation.UpdatedByCleared() {
		_spec.ClearField(procinst.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(procinst.FieldUpdatedAt, field.TypeTime, value)
	}
	if piu.mutation.UpdatedAtCleared() {
		_spec.ClearField(procinst.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := piu.mutation.BusinessKey(); ok {
		_spec.SetField(procinst.FieldBusinessKey, field.TypeString, value)
	}
	if value, ok := piu.mutation.StartTime(); ok {
		_spec.SetField(procinst.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := piu.mutation.EndTime(); ok {
		_spec.SetField(procinst.FieldEndTime, field.TypeTime, value)
	}
	if piu.mutation.EndTimeCleared() {
		_spec.ClearField(procinst.FieldEndTime, field.TypeTime)
	}
	if value, ok := piu.mutation.Duration(); ok {
		_spec.SetField(procinst.FieldDuration, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedDuration(); ok {
		_spec.AddField(procinst.FieldDuration, field.TypeInt, value)
	}
	if piu.mutation.DurationCleared() {
		_spec.ClearField(procinst.FieldDuration, field.TypeInt)
	}
	if value, ok := piu.mutation.StartUserID(); ok {
		_spec.SetField(procinst.FieldStartUserID, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedStartUserID(); ok {
		_spec.AddField(procinst.FieldStartUserID, field.TypeInt, value)
	}
	if value, ok := piu.mutation.SupperInstanceID(); ok {
		_spec.SetField(procinst.FieldSupperInstanceID, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedSupperInstanceID(); ok {
		_spec.AddField(procinst.FieldSupperInstanceID, field.TypeInt, value)
	}
	if piu.mutation.SupperInstanceIDCleared() {
		_spec.ClearField(procinst.FieldSupperInstanceID, field.TypeInt)
	}
	if value, ok := piu.mutation.RootInstanceID(); ok {
		_spec.SetField(procinst.FieldRootInstanceID, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedRootInstanceID(); ok {
		_spec.AddField(procinst.FieldRootInstanceID, field.TypeInt, value)
	}
	if piu.mutation.RootInstanceIDCleared() {
		_spec.ClearField(procinst.FieldRootInstanceID, field.TypeInt)
	}
	if value, ok := piu.mutation.DeletedTime(); ok {
		_spec.SetField(procinst.FieldDeletedTime, field.TypeTime, value)
	}
	if piu.mutation.DeletedTimeCleared() {
		_spec.ClearField(procinst.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := piu.mutation.DeletedReason(); ok {
		_spec.SetField(procinst.FieldDeletedReason, field.TypeString, value)
	}
	if piu.mutation.DeletedReasonCleared() {
		_spec.ClearField(procinst.FieldDeletedReason, field.TypeString)
	}
	if value, ok := piu.mutation.Status(); ok {
		_spec.SetField(procinst.FieldStatus, field.TypeEnum, value)
	}
	if piu.mutation.ProcDefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procinst.ProcDefTable,
			Columns: []string{procinst.ProcDefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piu.schemaConfig.ProcInst
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ProcDefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procinst.ProcDefTable,
			Columns: []string{procinst.ProcDefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piu.schemaConfig.ProcInst
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piu.schemaConfig.Task
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !piu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piu.schemaConfig.Task
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piu.schemaConfig.Task
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = piu.schemaConfig.ProcInst
	ctx = internal.NewSchemaConfigContext(ctx, piu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{procinst.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// ProcInstUpdateOne is the builder for updating a single ProcInst entity.
type ProcInstUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProcInstMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (piuo *ProcInstUpdateOne) SetUpdatedBy(i int) *ProcInstUpdateOne {
	piuo.mutation.ResetUpdatedBy()
	piuo.mutation.SetUpdatedBy(i)
	return piuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableUpdatedBy(i *int) *ProcInstUpdateOne {
	if i != nil {
		piuo.SetUpdatedBy(*i)
	}
	return piuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (piuo *ProcInstUpdateOne) AddUpdatedBy(i int) *ProcInstUpdateOne {
	piuo.mutation.AddUpdatedBy(i)
	return piuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (piuo *ProcInstUpdateOne) ClearUpdatedBy() *ProcInstUpdateOne {
	piuo.mutation.ClearUpdatedBy()
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *ProcInstUpdateOne) SetUpdatedAt(t time.Time) *ProcInstUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProcInstUpdateOne {
	if t != nil {
		piuo.SetUpdatedAt(*t)
	}
	return piuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piuo *ProcInstUpdateOne) ClearUpdatedAt() *ProcInstUpdateOne {
	piuo.mutation.ClearUpdatedAt()
	return piuo
}

// SetProcDefID sets the "proc_def_id" field.
func (piuo *ProcInstUpdateOne) SetProcDefID(i int) *ProcInstUpdateOne {
	piuo.mutation.SetProcDefID(i)
	return piuo
}

// SetBusinessKey sets the "business_key" field.
func (piuo *ProcInstUpdateOne) SetBusinessKey(s string) *ProcInstUpdateOne {
	piuo.mutation.SetBusinessKey(s)
	return piuo
}

// SetStartTime sets the "start_time" field.
func (piuo *ProcInstUpdateOne) SetStartTime(t time.Time) *ProcInstUpdateOne {
	piuo.mutation.SetStartTime(t)
	return piuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableStartTime(t *time.Time) *ProcInstUpdateOne {
	if t != nil {
		piuo.SetStartTime(*t)
	}
	return piuo
}

// SetEndTime sets the "end_time" field.
func (piuo *ProcInstUpdateOne) SetEndTime(t time.Time) *ProcInstUpdateOne {
	piuo.mutation.SetEndTime(t)
	return piuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableEndTime(t *time.Time) *ProcInstUpdateOne {
	if t != nil {
		piuo.SetEndTime(*t)
	}
	return piuo
}

// ClearEndTime clears the value of the "end_time" field.
func (piuo *ProcInstUpdateOne) ClearEndTime() *ProcInstUpdateOne {
	piuo.mutation.ClearEndTime()
	return piuo
}

// SetDuration sets the "duration" field.
func (piuo *ProcInstUpdateOne) SetDuration(i int) *ProcInstUpdateOne {
	piuo.mutation.ResetDuration()
	piuo.mutation.SetDuration(i)
	return piuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableDuration(i *int) *ProcInstUpdateOne {
	if i != nil {
		piuo.SetDuration(*i)
	}
	return piuo
}

// AddDuration adds i to the "duration" field.
func (piuo *ProcInstUpdateOne) AddDuration(i int) *ProcInstUpdateOne {
	piuo.mutation.AddDuration(i)
	return piuo
}

// ClearDuration clears the value of the "duration" field.
func (piuo *ProcInstUpdateOne) ClearDuration() *ProcInstUpdateOne {
	piuo.mutation.ClearDuration()
	return piuo
}

// SetStartUserID sets the "start_user_id" field.
func (piuo *ProcInstUpdateOne) SetStartUserID(i int) *ProcInstUpdateOne {
	piuo.mutation.ResetStartUserID()
	piuo.mutation.SetStartUserID(i)
	return piuo
}

// AddStartUserID adds i to the "start_user_id" field.
func (piuo *ProcInstUpdateOne) AddStartUserID(i int) *ProcInstUpdateOne {
	piuo.mutation.AddStartUserID(i)
	return piuo
}

// SetSupperInstanceID sets the "supper_instance_id" field.
func (piuo *ProcInstUpdateOne) SetSupperInstanceID(i int) *ProcInstUpdateOne {
	piuo.mutation.ResetSupperInstanceID()
	piuo.mutation.SetSupperInstanceID(i)
	return piuo
}

// SetNillableSupperInstanceID sets the "supper_instance_id" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableSupperInstanceID(i *int) *ProcInstUpdateOne {
	if i != nil {
		piuo.SetSupperInstanceID(*i)
	}
	return piuo
}

// AddSupperInstanceID adds i to the "supper_instance_id" field.
func (piuo *ProcInstUpdateOne) AddSupperInstanceID(i int) *ProcInstUpdateOne {
	piuo.mutation.AddSupperInstanceID(i)
	return piuo
}

// ClearSupperInstanceID clears the value of the "supper_instance_id" field.
func (piuo *ProcInstUpdateOne) ClearSupperInstanceID() *ProcInstUpdateOne {
	piuo.mutation.ClearSupperInstanceID()
	return piuo
}

// SetRootInstanceID sets the "root_instance_id" field.
func (piuo *ProcInstUpdateOne) SetRootInstanceID(i int) *ProcInstUpdateOne {
	piuo.mutation.ResetRootInstanceID()
	piuo.mutation.SetRootInstanceID(i)
	return piuo
}

// SetNillableRootInstanceID sets the "root_instance_id" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableRootInstanceID(i *int) *ProcInstUpdateOne {
	if i != nil {
		piuo.SetRootInstanceID(*i)
	}
	return piuo
}

// AddRootInstanceID adds i to the "root_instance_id" field.
func (piuo *ProcInstUpdateOne) AddRootInstanceID(i int) *ProcInstUpdateOne {
	piuo.mutation.AddRootInstanceID(i)
	return piuo
}

// ClearRootInstanceID clears the value of the "root_instance_id" field.
func (piuo *ProcInstUpdateOne) ClearRootInstanceID() *ProcInstUpdateOne {
	piuo.mutation.ClearRootInstanceID()
	return piuo
}

// SetDeletedTime sets the "deleted_time" field.
func (piuo *ProcInstUpdateOne) SetDeletedTime(t time.Time) *ProcInstUpdateOne {
	piuo.mutation.SetDeletedTime(t)
	return piuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableDeletedTime(t *time.Time) *ProcInstUpdateOne {
	if t != nil {
		piuo.SetDeletedTime(*t)
	}
	return piuo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (piuo *ProcInstUpdateOne) ClearDeletedTime() *ProcInstUpdateOne {
	piuo.mutation.ClearDeletedTime()
	return piuo
}

// SetDeletedReason sets the "deleted_reason" field.
func (piuo *ProcInstUpdateOne) SetDeletedReason(s string) *ProcInstUpdateOne {
	piuo.mutation.SetDeletedReason(s)
	return piuo
}

// SetNillableDeletedReason sets the "deleted_reason" field if the given value is not nil.
func (piuo *ProcInstUpdateOne) SetNillableDeletedReason(s *string) *ProcInstUpdateOne {
	if s != nil {
		piuo.SetDeletedReason(*s)
	}
	return piuo
}

// ClearDeletedReason clears the value of the "deleted_reason" field.
func (piuo *ProcInstUpdateOne) ClearDeletedReason() *ProcInstUpdateOne {
	piuo.mutation.ClearDeletedReason()
	return piuo
}

// SetStatus sets the "status" field.
func (piuo *ProcInstUpdateOne) SetStatus(pr procinst.Status) *ProcInstUpdateOne {
	piuo.mutation.SetStatus(pr)
	return piuo
}

// SetProcDef sets the "proc_def" edge to the ProcDef entity.
func (piuo *ProcInstUpdateOne) SetProcDef(p *ProcDef) *ProcInstUpdateOne {
	return piuo.SetProcDefID(p.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (piuo *ProcInstUpdateOne) AddTaskIDs(ids ...int) *ProcInstUpdateOne {
	piuo.mutation.AddTaskIDs(ids...)
	return piuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (piuo *ProcInstUpdateOne) AddTasks(t ...*Task) *ProcInstUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return piuo.AddTaskIDs(ids...)
}

// Mutation returns the ProcInstMutation object of the builder.
func (piuo *ProcInstUpdateOne) Mutation() *ProcInstMutation {
	return piuo.mutation
}

// ClearProcDef clears the "proc_def" edge to the ProcDef entity.
func (piuo *ProcInstUpdateOne) ClearProcDef() *ProcInstUpdateOne {
	piuo.mutation.ClearProcDef()
	return piuo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (piuo *ProcInstUpdateOne) ClearTasks() *ProcInstUpdateOne {
	piuo.mutation.ClearTasks()
	return piuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (piuo *ProcInstUpdateOne) RemoveTaskIDs(ids ...int) *ProcInstUpdateOne {
	piuo.mutation.RemoveTaskIDs(ids...)
	return piuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (piuo *ProcInstUpdateOne) RemoveTasks(t ...*Task) *ProcInstUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return piuo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the ProcInstUpdate builder.
func (piuo *ProcInstUpdateOne) Where(ps ...predicate.ProcInst) *ProcInstUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *ProcInstUpdateOne) Select(field string, fields ...string) *ProcInstUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated ProcInst entity.
func (piuo *ProcInstUpdateOne) Save(ctx context.Context) (*ProcInst, error) {
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *ProcInstUpdateOne) SaveX(ctx context.Context) *ProcInst {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *ProcInstUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *ProcInstUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *ProcInstUpdateOne) check() error {
	if v, ok := piuo.mutation.Status(); ok {
		if err := procinst.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ProcInst.status": %w`, err)}
		}
	}
	if _, ok := piuo.mutation.ProcDefID(); piuo.mutation.ProcDefCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProcInst.proc_def"`)
	}
	return nil
}

func (piuo *ProcInstUpdateOne) sqlSave(ctx context.Context) (_node *ProcInst, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(procinst.Table, procinst.Columns, sqlgraph.NewFieldSpec(procinst.FieldID, field.TypeInt))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProcInst.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, procinst.FieldID)
		for _, f := range fields {
			if !procinst.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != procinst.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.UpdatedBy(); ok {
		_spec.SetField(procinst.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(procinst.FieldUpdatedBy, field.TypeInt, value)
	}
	if piuo.mutation.UpdatedByCleared() {
		_spec.ClearField(procinst.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(procinst.FieldUpdatedAt, field.TypeTime, value)
	}
	if piuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(procinst.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := piuo.mutation.BusinessKey(); ok {
		_spec.SetField(procinst.FieldBusinessKey, field.TypeString, value)
	}
	if value, ok := piuo.mutation.StartTime(); ok {
		_spec.SetField(procinst.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.EndTime(); ok {
		_spec.SetField(procinst.FieldEndTime, field.TypeTime, value)
	}
	if piuo.mutation.EndTimeCleared() {
		_spec.ClearField(procinst.FieldEndTime, field.TypeTime)
	}
	if value, ok := piuo.mutation.Duration(); ok {
		_spec.SetField(procinst.FieldDuration, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedDuration(); ok {
		_spec.AddField(procinst.FieldDuration, field.TypeInt, value)
	}
	if piuo.mutation.DurationCleared() {
		_spec.ClearField(procinst.FieldDuration, field.TypeInt)
	}
	if value, ok := piuo.mutation.StartUserID(); ok {
		_spec.SetField(procinst.FieldStartUserID, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedStartUserID(); ok {
		_spec.AddField(procinst.FieldStartUserID, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.SupperInstanceID(); ok {
		_spec.SetField(procinst.FieldSupperInstanceID, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedSupperInstanceID(); ok {
		_spec.AddField(procinst.FieldSupperInstanceID, field.TypeInt, value)
	}
	if piuo.mutation.SupperInstanceIDCleared() {
		_spec.ClearField(procinst.FieldSupperInstanceID, field.TypeInt)
	}
	if value, ok := piuo.mutation.RootInstanceID(); ok {
		_spec.SetField(procinst.FieldRootInstanceID, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedRootInstanceID(); ok {
		_spec.AddField(procinst.FieldRootInstanceID, field.TypeInt, value)
	}
	if piuo.mutation.RootInstanceIDCleared() {
		_spec.ClearField(procinst.FieldRootInstanceID, field.TypeInt)
	}
	if value, ok := piuo.mutation.DeletedTime(); ok {
		_spec.SetField(procinst.FieldDeletedTime, field.TypeTime, value)
	}
	if piuo.mutation.DeletedTimeCleared() {
		_spec.ClearField(procinst.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := piuo.mutation.DeletedReason(); ok {
		_spec.SetField(procinst.FieldDeletedReason, field.TypeString, value)
	}
	if piuo.mutation.DeletedReasonCleared() {
		_spec.ClearField(procinst.FieldDeletedReason, field.TypeString)
	}
	if value, ok := piuo.mutation.Status(); ok {
		_spec.SetField(procinst.FieldStatus, field.TypeEnum, value)
	}
	if piuo.mutation.ProcDefCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procinst.ProcDefTable,
			Columns: []string{procinst.ProcDefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piuo.schemaConfig.ProcInst
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ProcDefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procinst.ProcDefTable,
			Columns: []string{procinst.ProcDefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procdef.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piuo.schemaConfig.ProcInst
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piuo.schemaConfig.Task
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !piuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piuo.schemaConfig.Task
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   procinst.TasksTable,
			Columns: []string{procinst.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		edge.Schema = piuo.schemaConfig.Task
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = piuo.schemaConfig.ProcInst
	ctx = internal.NewSchemaConfigContext(ctx, piuo.schemaConfig)
	_node = &ProcInst{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{procinst.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
