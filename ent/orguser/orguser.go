// Code generated by ent, DO NOT EDIT.

package orguser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orguser type in the database.
	Label = "org_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldJoinedAt holds the string denoting the joined_at field in the database.
	FieldJoinedAt = "joined_at"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// EdgeOrgRoles holds the string denoting the org_roles edge name in mutations.
	EdgeOrgRoles = "org_roles"
	// EdgeOrgRoleUser holds the string denoting the org_role_user edge name in mutations.
	EdgeOrgRoleUser = "org_role_user"
	// Table holds the table name of the orguser in the database.
	Table = "org_user"
	// OrgRolesTable is the table that holds the org_roles relation/edge. The primary key declared below.
	OrgRolesTable = "org_role_user"
	// OrgRolesInverseTable is the table name for the OrgRole entity.
	// It exists in this package in order to avoid circular dependency with the "orgrole" package.
	OrgRolesInverseTable = "org_role"
	// OrgRoleUserTable is the table that holds the org_role_user relation/edge.
	OrgRoleUserTable = "org_role_user"
	// OrgRoleUserInverseTable is the table name for the OrgRoleUser entity.
	// It exists in this package in order to avoid circular dependency with the "orgroleuser" package.
	OrgRoleUserInverseTable = "org_role_user"
	// OrgRoleUserColumn is the table column denoting the org_role_user relation/edge.
	OrgRoleUserColumn = "org_user_id"
)

// Columns holds all SQL columns for orguser fields.
var Columns = []string{
	FieldID,
	FieldOrgID,
	FieldUserID,
	FieldJoinedAt,
	FieldDisplayName,
}

var (
	// OrgRolesPrimaryKey and OrgRolesColumn2 are the table columns denoting the
	// primary key for the org_roles relation (M2M).
	OrgRolesPrimaryKey = []string{"org_role_id", "org_user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultJoinedAt holds the default value on creation for the "joined_at" field.
	DefaultJoinedAt func() time.Time
)

// OrderOption defines the ordering options for the OrgUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByJoinedAt orders the results by the joined_at field.
func ByJoinedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJoinedAt, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByOrgRolesCount orders the results by org_roles count.
func ByOrgRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrgRolesStep(), opts...)
	}
}

// ByOrgRoles orders the results by org_roles terms.
func ByOrgRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrgRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrgRoleUserCount orders the results by org_role_user count.
func ByOrgRoleUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrgRoleUserStep(), opts...)
	}
}

// ByOrgRoleUser orders the results by org_role_user terms.
func ByOrgRoleUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrgRoleUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrgRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrgRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OrgRolesTable, OrgRolesPrimaryKey...),
	)
}
func newOrgRoleUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrgRoleUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OrgRoleUserTable, OrgRoleUserColumn),
	)
}
