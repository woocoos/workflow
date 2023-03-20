// Code generated by ent, DO NOT EDIT.

package deployment

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
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
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldDeployTime holds the string denoting the deploy_time field in the database.
	FieldDeployTime = "deploy_time"
	// EdgeProcDefs holds the string denoting the proc_defs edge name in mutations.
	EdgeProcDefs = "proc_defs"
	// EdgeDecisionReqs holds the string denoting the decision_reqs edge name in mutations.
	EdgeDecisionReqs = "decision_reqs"
	// Table holds the table name of the deployment in the database.
	Table = "act_deployment"
	// ProcDefsTable is the table that holds the proc_defs relation/edge.
	ProcDefsTable = "act_proc_def"
	// ProcDefsInverseTable is the table name for the ProcDef entity.
	// It exists in this package in order to avoid circular dependency with the "procdef" package.
	ProcDefsInverseTable = "act_proc_def"
	// ProcDefsColumn is the table column denoting the proc_defs relation/edge.
	ProcDefsColumn = "deployment_id"
	// DecisionReqsTable is the table that holds the decision_reqs relation/edge.
	DecisionReqsTable = "act_decision_req_def"
	// DecisionReqsInverseTable is the table name for the DecisionReqDef entity.
	// It exists in this package in order to avoid circular dependency with the "decisionreqdef" package.
	DecisionReqsInverseTable = "act_decision_req_def"
	// DecisionReqsColumn is the table column denoting the decision_reqs relation/edge.
	DecisionReqsColumn = "deployment_id"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldOrgID,
	FieldAppID,
	FieldName,
	FieldSource,
	FieldDeployTime,
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
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultDeployTime holds the default value on creation for the "deploy_time" field.
	DefaultDeployTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)