// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/workflow/ent/decisionreqdef"
	"github.com/woocoos/workflow/ent/deployment"
)

// DecisionReqDef is the model entity for the DecisionReqDef schema.
type DecisionReqDef struct {
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
	// 部署ID
	DeploymentID int `json:"deployment_id,omitempty"`
	// 所属根组织ID
	OrgID int `json:"org_id,omitempty"`
	// 所属应用ID
	AppID int `json:"app_id,omitempty"`
	// 分类
	Category string `json:"category,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// KEY
	Key string `json:"key,omitempty"`
	// 版本
	Version int32 `json:"version,omitempty"`
	// 小版本
	Revision int32 `json:"revision,omitempty"`
	// 资源名称
	ResourceName string `json:"resource_name,omitempty"`
	// 流程图资源名称
	DgrmResourceName string `json:"dgrm_resource_name,omitempty"`
	// 资源数据
	ResourceData []byte `json:"resource_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DecisionReqDefQuery when eager-loading is set.
	Edges DecisionReqDefEdges `json:"edges"`
}

// DecisionReqDefEdges holds the relations/edges for other nodes in the graph.
type DecisionReqDefEdges struct {
	// Deployment holds the value of the deployment edge.
	Deployment *Deployment `json:"deployment,omitempty"`
	// DecisionDefs holds the value of the decision_defs edge.
	DecisionDefs []*DecisionDef `json:"decision_defs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedDecisionDefs map[string][]*DecisionDef
}

// DeploymentOrErr returns the Deployment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DecisionReqDefEdges) DeploymentOrErr() (*Deployment, error) {
	if e.loadedTypes[0] {
		if e.Deployment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deployment.Label}
		}
		return e.Deployment, nil
	}
	return nil, &NotLoadedError{edge: "deployment"}
}

// DecisionDefsOrErr returns the DecisionDefs value or an error if the edge
// was not loaded in eager-loading.
func (e DecisionReqDefEdges) DecisionDefsOrErr() ([]*DecisionDef, error) {
	if e.loadedTypes[1] {
		return e.DecisionDefs, nil
	}
	return nil, &NotLoadedError{edge: "decision_defs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DecisionReqDef) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case decisionreqdef.FieldResourceData:
			values[i] = new([]byte)
		case decisionreqdef.FieldID, decisionreqdef.FieldCreatedBy, decisionreqdef.FieldUpdatedBy, decisionreqdef.FieldDeploymentID, decisionreqdef.FieldOrgID, decisionreqdef.FieldAppID, decisionreqdef.FieldVersion, decisionreqdef.FieldRevision:
			values[i] = new(sql.NullInt64)
		case decisionreqdef.FieldCategory, decisionreqdef.FieldName, decisionreqdef.FieldKey, decisionreqdef.FieldResourceName, decisionreqdef.FieldDgrmResourceName:
			values[i] = new(sql.NullString)
		case decisionreqdef.FieldCreatedAt, decisionreqdef.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DecisionReqDef", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DecisionReqDef fields.
func (drd *DecisionReqDef) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case decisionreqdef.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			drd.ID = int(value.Int64)
		case decisionreqdef.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				drd.CreatedBy = int(value.Int64)
			}
		case decisionreqdef.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				drd.CreatedAt = value.Time
			}
		case decisionreqdef.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				drd.UpdatedBy = int(value.Int64)
			}
		case decisionreqdef.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				drd.UpdatedAt = value.Time
			}
		case decisionreqdef.FieldDeploymentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_id", values[i])
			} else if value.Valid {
				drd.DeploymentID = int(value.Int64)
			}
		case decisionreqdef.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				drd.OrgID = int(value.Int64)
			}
		case decisionreqdef.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				drd.AppID = int(value.Int64)
			}
		case decisionreqdef.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				drd.Category = value.String
			}
		case decisionreqdef.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				drd.Name = value.String
			}
		case decisionreqdef.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				drd.Key = value.String
			}
		case decisionreqdef.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				drd.Version = int32(value.Int64)
			}
		case decisionreqdef.FieldRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				drd.Revision = int32(value.Int64)
			}
		case decisionreqdef.FieldResourceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource_name", values[i])
			} else if value.Valid {
				drd.ResourceName = value.String
			}
		case decisionreqdef.FieldDgrmResourceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dgrm_resource_name", values[i])
			} else if value.Valid {
				drd.DgrmResourceName = value.String
			}
		case decisionreqdef.FieldResourceData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field resource_data", values[i])
			} else if value != nil {
				drd.ResourceData = *value
			}
		}
	}
	return nil
}

// QueryDeployment queries the "deployment" edge of the DecisionReqDef entity.
func (drd *DecisionReqDef) QueryDeployment() *DeploymentQuery {
	return NewDecisionReqDefClient(drd.config).QueryDeployment(drd)
}

// QueryDecisionDefs queries the "decision_defs" edge of the DecisionReqDef entity.
func (drd *DecisionReqDef) QueryDecisionDefs() *DecisionDefQuery {
	return NewDecisionReqDefClient(drd.config).QueryDecisionDefs(drd)
}

// Update returns a builder for updating this DecisionReqDef.
// Note that you need to call DecisionReqDef.Unwrap() before calling this method if this DecisionReqDef
// was returned from a transaction, and the transaction was committed or rolled back.
func (drd *DecisionReqDef) Update() *DecisionReqDefUpdateOne {
	return NewDecisionReqDefClient(drd.config).UpdateOne(drd)
}

// Unwrap unwraps the DecisionReqDef entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (drd *DecisionReqDef) Unwrap() *DecisionReqDef {
	_tx, ok := drd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DecisionReqDef is not a transactional entity")
	}
	drd.config.driver = _tx.drv
	return drd
}

// String implements the fmt.Stringer.
func (drd *DecisionReqDef) String() string {
	var builder strings.Builder
	builder.WriteString("DecisionReqDef(")
	builder.WriteString(fmt.Sprintf("id=%v, ", drd.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", drd.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(drd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", drd.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(drd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deployment_id=")
	builder.WriteString(fmt.Sprintf("%v", drd.DeploymentID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", drd.OrgID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", drd.AppID))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(drd.Category)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(drd.Name)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(drd.Key)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", drd.Version))
	builder.WriteString(", ")
	builder.WriteString("revision=")
	builder.WriteString(fmt.Sprintf("%v", drd.Revision))
	builder.WriteString(", ")
	builder.WriteString("resource_name=")
	builder.WriteString(drd.ResourceName)
	builder.WriteString(", ")
	builder.WriteString("dgrm_resource_name=")
	builder.WriteString(drd.DgrmResourceName)
	builder.WriteString(", ")
	builder.WriteString("resource_data=")
	builder.WriteString(fmt.Sprintf("%v", drd.ResourceData))
	builder.WriteByte(')')
	return builder.String()
}

// NamedDecisionDefs returns the DecisionDefs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (drd *DecisionReqDef) NamedDecisionDefs(name string) ([]*DecisionDef, error) {
	if drd.Edges.namedDecisionDefs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := drd.Edges.namedDecisionDefs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (drd *DecisionReqDef) appendNamedDecisionDefs(name string, edges ...*DecisionDef) {
	if drd.Edges.namedDecisionDefs == nil {
		drd.Edges.namedDecisionDefs = make(map[string][]*DecisionDef)
	}
	if len(edges) == 0 {
		drd.Edges.namedDecisionDefs[name] = []*DecisionDef{}
	} else {
		drd.Edges.namedDecisionDefs[name] = append(drd.Edges.namedDecisionDefs[name], edges...)
	}
}

// DecisionReqDefs is a parsable slice of DecisionReqDef.
type DecisionReqDefs []*DecisionReqDef
