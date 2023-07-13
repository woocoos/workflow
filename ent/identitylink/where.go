// Code generated by ent, DO NOT EDIT.

package identitylink

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/workflow/ent/predicate"

	"github.com/woocoos/workflow/ent/internal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldID, id))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldTenantID, v))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldTaskID, v))
}

// ProcDefID applies equality check predicate on the "proc_def_id" field. It's identical to ProcDefIDEQ.
func ProcDefID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldProcDefID, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldGroupID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldUserID, v))
}

// AssignerID applies equality check predicate on the "assigner_id" field. It's identical to AssignerIDEQ.
func AssignerID(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldAssignerID, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldComments, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldTenantID, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldTaskID, vs...))
}

// ProcDefIDEQ applies the EQ predicate on the "proc_def_id" field.
func ProcDefIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldProcDefID, v))
}

// ProcDefIDNEQ applies the NEQ predicate on the "proc_def_id" field.
func ProcDefIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldProcDefID, v))
}

// ProcDefIDIn applies the In predicate on the "proc_def_id" field.
func ProcDefIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldProcDefID, vs...))
}

// ProcDefIDNotIn applies the NotIn predicate on the "proc_def_id" field.
func ProcDefIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldProcDefID, vs...))
}

// ProcDefIDGT applies the GT predicate on the "proc_def_id" field.
func ProcDefIDGT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldProcDefID, v))
}

// ProcDefIDGTE applies the GTE predicate on the "proc_def_id" field.
func ProcDefIDGTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldProcDefID, v))
}

// ProcDefIDLT applies the LT predicate on the "proc_def_id" field.
func ProcDefIDLT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldProcDefID, v))
}

// ProcDefIDLTE applies the LTE predicate on the "proc_def_id" field.
func ProcDefIDLTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldProcDefID, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldGroupID, v))
}

// GroupIDIsNil applies the IsNil predicate on the "group_id" field.
func GroupIDIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIsNull(FieldGroupID))
}

// GroupIDNotNil applies the NotNil predicate on the "group_id" field.
func GroupIDNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotNull(FieldGroupID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotNull(FieldUserID))
}

// AssignerIDEQ applies the EQ predicate on the "assigner_id" field.
func AssignerIDEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldAssignerID, v))
}

// AssignerIDNEQ applies the NEQ predicate on the "assigner_id" field.
func AssignerIDNEQ(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldAssignerID, v))
}

// AssignerIDIn applies the In predicate on the "assigner_id" field.
func AssignerIDIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldAssignerID, vs...))
}

// AssignerIDNotIn applies the NotIn predicate on the "assigner_id" field.
func AssignerIDNotIn(vs ...int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldAssignerID, vs...))
}

// AssignerIDGT applies the GT predicate on the "assigner_id" field.
func AssignerIDGT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldAssignerID, v))
}

// AssignerIDGTE applies the GTE predicate on the "assigner_id" field.
func AssignerIDGTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldAssignerID, v))
}

// AssignerIDLT applies the LT predicate on the "assigner_id" field.
func AssignerIDLT(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldAssignerID, v))
}

// AssignerIDLTE applies the LTE predicate on the "assigner_id" field.
func AssignerIDLTE(v int) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldAssignerID, v))
}

// AssignerIDIsNil applies the IsNil predicate on the "assigner_id" field.
func AssignerIDIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIsNull(FieldAssignerID))
}

// AssignerIDNotNil applies the NotNil predicate on the "assigner_id" field.
func AssignerIDNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotNull(FieldAssignerID))
}

// LinkTypeEQ applies the EQ predicate on the "link_type" field.
func LinkTypeEQ(v LinkType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldLinkType, v))
}

// LinkTypeNEQ applies the NEQ predicate on the "link_type" field.
func LinkTypeNEQ(v LinkType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldLinkType, v))
}

// LinkTypeIn applies the In predicate on the "link_type" field.
func LinkTypeIn(vs ...LinkType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldLinkType, vs...))
}

// LinkTypeNotIn applies the NotIn predicate on the "link_type" field.
func LinkTypeNotIn(vs ...LinkType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldLinkType, vs...))
}

// OperationTypeEQ applies the EQ predicate on the "operation_type" field.
func OperationTypeEQ(v OperationType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldOperationType, v))
}

// OperationTypeNEQ applies the NEQ predicate on the "operation_type" field.
func OperationTypeNEQ(v OperationType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldOperationType, v))
}

// OperationTypeIn applies the In predicate on the "operation_type" field.
func OperationTypeIn(vs ...OperationType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldOperationType, vs...))
}

// OperationTypeNotIn applies the NotIn predicate on the "operation_type" field.
func OperationTypeNotIn(vs ...OperationType) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldOperationType, vs...))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.IdentityLink {
	return predicate.IdentityLink(sql.FieldContainsFold(FieldComments, v))
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TaskTable, TaskColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Task
		step.Edge.Schema = schemaConfig.IdentityLink
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.IdentityLink {
	return predicate.IdentityLink(func(s *sql.Selector) {
		step := newTaskStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Task
		step.Edge.Schema = schemaConfig.IdentityLink
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.IdentityLink) predicate.IdentityLink {
	return predicate.IdentityLink(sql.NotPredicates(p))
}
