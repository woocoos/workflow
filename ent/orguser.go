// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/workflow/ent/orguser"
)

// OrgUser is the model entity for the OrgUser schema.
type OrgUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 用户ID
	UserID int `json:"user_id,omitempty"`
	// 加入时间
	JoinedAt time.Time `json:"joined_at,omitempty"`
	// 在组织内的显示名称
	DisplayName string `json:"display_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgUserQuery when eager-loading is set.
	Edges        OrgUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgUserEdges holds the relations/edges for other nodes in the graph.
type OrgUserEdges struct {
	// OrgRoles holds the value of the org_roles edge.
	OrgRoles []*OrgRole `json:"org_roles,omitempty"`
	// OrgRoleUser holds the value of the org_role_user edge.
	OrgRoleUser []*OrgRoleUser `json:"org_role_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool

	namedOrgRoles    map[string][]*OrgRole
	namedOrgRoleUser map[string][]*OrgRoleUser
}

// OrgRolesOrErr returns the OrgRoles value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUserEdges) OrgRolesOrErr() ([]*OrgRole, error) {
	if e.loadedTypes[0] {
		return e.OrgRoles, nil
	}
	return nil, &NotLoadedError{edge: "org_roles"}
}

// OrgRoleUserOrErr returns the OrgRoleUser value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUserEdges) OrgRoleUserOrErr() ([]*OrgRoleUser, error) {
	if e.loadedTypes[1] {
		return e.OrgRoleUser, nil
	}
	return nil, &NotLoadedError{edge: "org_role_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orguser.FieldID, orguser.FieldOrgID, orguser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case orguser.FieldDisplayName:
			values[i] = new(sql.NullString)
		case orguser.FieldJoinedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgUser fields.
func (ou *OrgUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orguser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ou.ID = int(value.Int64)
		case orguser.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				ou.OrgID = int(value.Int64)
			}
		case orguser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ou.UserID = int(value.Int64)
			}
		case orguser.FieldJoinedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field joined_at", values[i])
			} else if value.Valid {
				ou.JoinedAt = value.Time
			}
		case orguser.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				ou.DisplayName = value.String
			}
		default:
			ou.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgUser.
// This includes values selected through modifiers, order, etc.
func (ou *OrgUser) Value(name string) (ent.Value, error) {
	return ou.selectValues.Get(name)
}

// QueryOrgRoles queries the "org_roles" edge of the OrgUser entity.
func (ou *OrgUser) QueryOrgRoles() *OrgRoleQuery {
	return NewOrgUserClient(ou.config).QueryOrgRoles(ou)
}

// QueryOrgRoleUser queries the "org_role_user" edge of the OrgUser entity.
func (ou *OrgUser) QueryOrgRoleUser() *OrgRoleUserQuery {
	return NewOrgUserClient(ou.config).QueryOrgRoleUser(ou)
}

// Update returns a builder for updating this OrgUser.
// Note that you need to call OrgUser.Unwrap() before calling this method if this OrgUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ou *OrgUser) Update() *OrgUserUpdateOne {
	return NewOrgUserClient(ou.config).UpdateOne(ou)
}

// Unwrap unwraps the OrgUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ou *OrgUser) Unwrap() *OrgUser {
	_tx, ok := ou.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgUser is not a transactional entity")
	}
	ou.config.driver = _tx.drv
	return ou
}

// String implements the fmt.Stringer.
func (ou *OrgUser) String() string {
	var builder strings.Builder
	builder.WriteString("OrgUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ou.ID))
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", ou.OrgID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ou.UserID))
	builder.WriteString(", ")
	builder.WriteString("joined_at=")
	builder.WriteString(ou.JoinedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(ou.DisplayName)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOrgRoles returns the OrgRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ou *OrgUser) NamedOrgRoles(name string) ([]*OrgRole, error) {
	if ou.Edges.namedOrgRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ou.Edges.namedOrgRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ou *OrgUser) appendNamedOrgRoles(name string, edges ...*OrgRole) {
	if ou.Edges.namedOrgRoles == nil {
		ou.Edges.namedOrgRoles = make(map[string][]*OrgRole)
	}
	if len(edges) == 0 {
		ou.Edges.namedOrgRoles[name] = []*OrgRole{}
	} else {
		ou.Edges.namedOrgRoles[name] = append(ou.Edges.namedOrgRoles[name], edges...)
	}
}

// NamedOrgRoleUser returns the OrgRoleUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ou *OrgUser) NamedOrgRoleUser(name string) ([]*OrgRoleUser, error) {
	if ou.Edges.namedOrgRoleUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ou.Edges.namedOrgRoleUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ou *OrgUser) appendNamedOrgRoleUser(name string, edges ...*OrgRoleUser) {
	if ou.Edges.namedOrgRoleUser == nil {
		ou.Edges.namedOrgRoleUser = make(map[string][]*OrgRoleUser)
	}
	if len(edges) == 0 {
		ou.Edges.namedOrgRoleUser[name] = []*OrgRoleUser{}
	} else {
		ou.Edges.namedOrgRoleUser[name] = append(ou.Edges.namedOrgRoleUser[name], edges...)
	}
}

// OrgUsers is a parsable slice of OrgUser.
type OrgUsers []*OrgUser
