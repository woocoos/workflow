// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/workflow/ent/orgrole"
)

// OrgRole is the model entity for the OrgRole schema.
type OrgRole struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 类型,group:组,role:角色
	Kind orgrole.Kind `json:"kind,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgRoleQuery when eager-loading is set.
	Edges        OrgRoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgRoleEdges holds the relations/edges for other nodes in the graph.
type OrgRoleEdges struct {
	// 组用户
	OrgUsers []*OrgUser `json:"org_users,omitempty"`
	// OrgRoleUser holds the value of the org_role_user edge.
	OrgRoleUser []*OrgRoleUser `json:"org_role_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool

	namedOrgUsers    map[string][]*OrgUser
	namedOrgRoleUser map[string][]*OrgRoleUser
}

// OrgUsersOrErr returns the OrgUsers value or an error if the edge
// was not loaded in eager-loading.
func (e OrgRoleEdges) OrgUsersOrErr() ([]*OrgUser, error) {
	if e.loadedTypes[0] {
		return e.OrgUsers, nil
	}
	return nil, &NotLoadedError{edge: "org_users"}
}

// OrgRoleUserOrErr returns the OrgRoleUser value or an error if the edge
// was not loaded in eager-loading.
func (e OrgRoleEdges) OrgRoleUserOrErr() ([]*OrgRoleUser, error) {
	if e.loadedTypes[1] {
		return e.OrgRoleUser, nil
	}
	return nil, &NotLoadedError{edge: "org_role_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgrole.FieldID, orgrole.FieldOrgID:
			values[i] = new(sql.NullInt64)
		case orgrole.FieldKind, orgrole.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgRole fields.
func (or *OrgRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case orgrole.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				or.OrgID = int(value.Int64)
			}
		case orgrole.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				or.Kind = orgrole.Kind(value.String)
			}
		case orgrole.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				or.Name = value.String
			}
		default:
			or.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgRole.
// This includes values selected through modifiers, order, etc.
func (or *OrgRole) Value(name string) (ent.Value, error) {
	return or.selectValues.Get(name)
}

// QueryOrgUsers queries the "org_users" edge of the OrgRole entity.
func (or *OrgRole) QueryOrgUsers() *OrgUserQuery {
	return NewOrgRoleClient(or.config).QueryOrgUsers(or)
}

// QueryOrgRoleUser queries the "org_role_user" edge of the OrgRole entity.
func (or *OrgRole) QueryOrgRoleUser() *OrgRoleUserQuery {
	return NewOrgRoleClient(or.config).QueryOrgRoleUser(or)
}

// Update returns a builder for updating this OrgRole.
// Note that you need to call OrgRole.Unwrap() before calling this method if this OrgRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OrgRole) Update() *OrgRoleUpdateOne {
	return NewOrgRoleClient(or.config).UpdateOne(or)
}

// Unwrap unwraps the OrgRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OrgRole) Unwrap() *OrgRole {
	_tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgRole is not a transactional entity")
	}
	or.config.driver = _tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OrgRole) String() string {
	var builder strings.Builder
	builder.WriteString("OrgRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", or.ID))
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", or.OrgID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", or.Kind))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(or.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOrgUsers returns the OrgUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (or *OrgRole) NamedOrgUsers(name string) ([]*OrgUser, error) {
	if or.Edges.namedOrgUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := or.Edges.namedOrgUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (or *OrgRole) appendNamedOrgUsers(name string, edges ...*OrgUser) {
	if or.Edges.namedOrgUsers == nil {
		or.Edges.namedOrgUsers = make(map[string][]*OrgUser)
	}
	if len(edges) == 0 {
		or.Edges.namedOrgUsers[name] = []*OrgUser{}
	} else {
		or.Edges.namedOrgUsers[name] = append(or.Edges.namedOrgUsers[name], edges...)
	}
}

// NamedOrgRoleUser returns the OrgRoleUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (or *OrgRole) NamedOrgRoleUser(name string) ([]*OrgRoleUser, error) {
	if or.Edges.namedOrgRoleUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := or.Edges.namedOrgRoleUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (or *OrgRole) appendNamedOrgRoleUser(name string, edges ...*OrgRoleUser) {
	if or.Edges.namedOrgRoleUser == nil {
		or.Edges.namedOrgRoleUser = make(map[string][]*OrgRoleUser)
	}
	if len(edges) == 0 {
		or.Edges.namedOrgRoleUser[name] = []*OrgRoleUser{}
	} else {
		or.Edges.namedOrgRoleUser[name] = append(or.Edges.namedOrgRoleUser[name], edges...)
	}
}

// OrgRoles is a parsable slice of OrgRole.
type OrgRoles []*OrgRole