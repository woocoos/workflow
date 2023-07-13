// Code generated by ent, DO NOT EDIT.

package procinst

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the procinst type in the database.
	Label = "proc_inst"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldProcDefID holds the string denoting the proc_def_id field in the database.
	FieldProcDefID = "proc_def_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldBusinessKey holds the string denoting the business_key field in the database.
	FieldBusinessKey = "business_key"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldStartUserID holds the string denoting the start_user_id field in the database.
	FieldStartUserID = "start_user_id"
	// FieldSupperInstanceID holds the string denoting the supper_instance_id field in the database.
	FieldSupperInstanceID = "supper_instance_id"
	// FieldRootInstanceID holds the string denoting the root_instance_id field in the database.
	FieldRootInstanceID = "root_instance_id"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldDeletedReason holds the string denoting the deleted_reason field in the database.
	FieldDeletedReason = "deleted_reason"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeProcDef holds the string denoting the proc_def edge name in mutations.
	EdgeProcDef = "proc_def"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// Table holds the table name of the procinst in the database.
	Table = "act_proc_inst"
	// ProcDefTable is the table that holds the proc_def relation/edge.
	ProcDefTable = "act_proc_inst"
	// ProcDefInverseTable is the table name for the ProcDef entity.
	// It exists in this package in order to avoid circular dependency with the "procdef" package.
	ProcDefInverseTable = "act_proc_def"
	// ProcDefColumn is the table column denoting the proc_def relation/edge.
	ProcDefColumn = "proc_def_id"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "act_task"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "act_task"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "proc_inst_id"
)

// Columns holds all SQL columns for procinst fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldTenantID,
	FieldProcDefID,
	FieldAppID,
	FieldBusinessKey,
	FieldStartTime,
	FieldEndTime,
	FieldDuration,
	FieldStartUserID,
	FieldSupperInstanceID,
	FieldRootInstanceID,
	FieldDeletedTime,
	FieldDeletedReason,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/woocoos/workflow/ent/runtime"
var (
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusReady      Status = "ready"
	StatusActive     Status = "active"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
	StatusTerminated Status = "terminated"
	StatusSuspended  Status = "suspended"
	StatusDeleted    Status = "deleted"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusReady, StatusActive, StatusCompleted, StatusFailed, StatusTerminated, StatusSuspended, StatusDeleted:
		return nil
	default:
		return fmt.Errorf("procinst: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the ProcInst queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByProcDefID orders the results by the proc_def_id field.
func ByProcDefID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProcDefID, opts...).ToFunc()
}

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByBusinessKey orders the results by the business_key field.
func ByBusinessKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessKey, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByStartUserID orders the results by the start_user_id field.
func ByStartUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartUserID, opts...).ToFunc()
}

// BySupperInstanceID orders the results by the supper_instance_id field.
func BySupperInstanceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupperInstanceID, opts...).ToFunc()
}

// ByRootInstanceID orders the results by the root_instance_id field.
func ByRootInstanceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRootInstanceID, opts...).ToFunc()
}

// ByDeletedTime orders the results by the deleted_time field.
func ByDeletedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedTime, opts...).ToFunc()
}

// ByDeletedReason orders the results by the deleted_reason field.
func ByDeletedReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedReason, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByProcDefField orders the results by proc_def field.
func ByProcDefField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProcDefStep(), sql.OrderByField(field, opts...))
	}
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProcDefStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProcDefInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProcDefTable, ProcDefColumn),
	)
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
