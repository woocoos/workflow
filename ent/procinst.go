// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/workflow/ent/procdef"
	"github.com/woocoos/workflow/ent/procinst"
)

// ProcInst is the model entity for the ProcInst schema.
type ProcInst struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 流程定义ID
	ProcDefID int `json:"proc_def_id,omitempty"`
	// 所属根组织ID
	OrgID int `json:"org_id,omitempty"`
	// 所属应用ID
	AppID int `json:"app_id,omitempty"`
	// 业务主键
	BusinessKey string `json:"business_key,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 结束时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 持续时间
	Duration int `json:"duration,omitempty"`
	// 发起人ID/名称
	StartUserID int `json:"start_user_id,omitempty"`
	// 父流程实例ID
	SupperInstanceID int `json:"supper_instance_id,omitempty"`
	// 根流程实例ID
	RootInstanceID int `json:"root_instance_id,omitempty"`
	// 删除时间
	DeletedTime time.Time `json:"deleted_time,omitempty"`
	// 删除原因
	DeletedReason string `json:"deleted_reason,omitempty"`
	// 状态
	Status procinst.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProcInstQuery when eager-loading is set.
	Edges ProcInstEdges `json:"edges"`
}

// ProcInstEdges holds the relations/edges for other nodes in the graph.
type ProcInstEdges struct {
	// 流程定义
	ProcDef *ProcDef `json:"proc_def,omitempty"`
	// 任务列表
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedTasks map[string][]*Task
}

// ProcDefOrErr returns the ProcDef value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProcInstEdges) ProcDefOrErr() (*ProcDef, error) {
	if e.loadedTypes[0] {
		if e.ProcDef == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: procdef.Label}
		}
		return e.ProcDef, nil
	}
	return nil, &NotLoadedError{edge: "proc_def"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ProcInstEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProcInst) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case procinst.FieldID, procinst.FieldCreatedBy, procinst.FieldUpdatedBy, procinst.FieldProcDefID, procinst.FieldOrgID, procinst.FieldAppID, procinst.FieldDuration, procinst.FieldStartUserID, procinst.FieldSupperInstanceID, procinst.FieldRootInstanceID:
			values[i] = new(sql.NullInt64)
		case procinst.FieldBusinessKey, procinst.FieldDeletedReason, procinst.FieldStatus:
			values[i] = new(sql.NullString)
		case procinst.FieldCreatedAt, procinst.FieldUpdatedAt, procinst.FieldStartTime, procinst.FieldEndTime, procinst.FieldDeletedTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProcInst", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProcInst fields.
func (pi *ProcInst) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case procinst.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case procinst.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pi.CreatedBy = int(value.Int64)
			}
		case procinst.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case procinst.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pi.UpdatedBy = int(value.Int64)
			}
		case procinst.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		case procinst.FieldProcDefID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proc_def_id", values[i])
			} else if value.Valid {
				pi.ProcDefID = int(value.Int64)
			}
		case procinst.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				pi.OrgID = int(value.Int64)
			}
		case procinst.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				pi.AppID = int(value.Int64)
			}
		case procinst.FieldBusinessKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_key", values[i])
			} else if value.Valid {
				pi.BusinessKey = value.String
			}
		case procinst.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				pi.StartTime = value.Time
			}
		case procinst.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				pi.EndTime = value.Time
			}
		case procinst.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pi.Duration = int(value.Int64)
			}
		case procinst.FieldStartUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_user_id", values[i])
			} else if value.Valid {
				pi.StartUserID = int(value.Int64)
			}
		case procinst.FieldSupperInstanceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field supper_instance_id", values[i])
			} else if value.Valid {
				pi.SupperInstanceID = int(value.Int64)
			}
		case procinst.FieldRootInstanceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field root_instance_id", values[i])
			} else if value.Valid {
				pi.RootInstanceID = int(value.Int64)
			}
		case procinst.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				pi.DeletedTime = value.Time
			}
		case procinst.FieldDeletedReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_reason", values[i])
			} else if value.Valid {
				pi.DeletedReason = value.String
			}
		case procinst.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pi.Status = procinst.Status(value.String)
			}
		}
	}
	return nil
}

// QueryProcDef queries the "proc_def" edge of the ProcInst entity.
func (pi *ProcInst) QueryProcDef() *ProcDefQuery {
	return NewProcInstClient(pi.config).QueryProcDef(pi)
}

// QueryTasks queries the "tasks" edge of the ProcInst entity.
func (pi *ProcInst) QueryTasks() *TaskQuery {
	return NewProcInstClient(pi.config).QueryTasks(pi)
}

// Update returns a builder for updating this ProcInst.
// Note that you need to call ProcInst.Unwrap() before calling this method if this ProcInst
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *ProcInst) Update() *ProcInstUpdateOne {
	return NewProcInstClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the ProcInst entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *ProcInst) Unwrap() *ProcInst {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProcInst is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *ProcInst) String() string {
	var builder strings.Builder
	builder.WriteString("ProcInst(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pi.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("proc_def_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.ProcDefID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.OrgID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.AppID))
	builder.WriteString(", ")
	builder.WriteString("business_key=")
	builder.WriteString(pi.BusinessKey)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(pi.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(pi.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", pi.Duration))
	builder.WriteString(", ")
	builder.WriteString("start_user_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.StartUserID))
	builder.WriteString(", ")
	builder.WriteString("supper_instance_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.SupperInstanceID))
	builder.WriteString(", ")
	builder.WriteString("root_instance_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.RootInstanceID))
	builder.WriteString(", ")
	builder.WriteString("deleted_time=")
	builder.WriteString(pi.DeletedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_reason=")
	builder.WriteString(pi.DeletedReason)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pi.Status))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pi *ProcInst) NamedTasks(name string) ([]*Task, error) {
	if pi.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pi.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pi *ProcInst) appendNamedTasks(name string, edges ...*Task) {
	if pi.Edges.namedTasks == nil {
		pi.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		pi.Edges.namedTasks[name] = []*Task{}
	} else {
		pi.Edges.namedTasks[name] = append(pi.Edges.namedTasks[name], edges...)
	}
}

// ProcInsts is a parsable slice of ProcInst.
type ProcInsts []*ProcInst
