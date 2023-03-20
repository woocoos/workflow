// Code generated by ent, DO NOT EDIT.

package decisionreqdef

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/workflow/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeploymentID applies equality check predicate on the "deployment_id" field. It's identical to DeploymentIDEQ.
func DeploymentID(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldDeploymentID, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldOrgID, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldAppID, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCategory, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldName, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldKey, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldVersion, v))
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldRevision, v))
}

// ResourceName applies equality check predicate on the "resource_name" field. It's identical to ResourceNameEQ.
func ResourceName(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldResourceName, v))
}

// DgrmResourceName applies equality check predicate on the "dgrm_resource_name" field. It's identical to DgrmResourceNameEQ.
func DgrmResourceName(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldDgrmResourceName, v))
}

// ResourceData applies equality check predicate on the "resource_data" field. It's identical to ResourceDataEQ.
func ResourceData(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldResourceData, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldUpdatedAt))
}

// DeploymentIDEQ applies the EQ predicate on the "deployment_id" field.
func DeploymentIDEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldDeploymentID, v))
}

// DeploymentIDNEQ applies the NEQ predicate on the "deployment_id" field.
func DeploymentIDNEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldDeploymentID, v))
}

// DeploymentIDIn applies the In predicate on the "deployment_id" field.
func DeploymentIDIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldDeploymentID, vs...))
}

// DeploymentIDNotIn applies the NotIn predicate on the "deployment_id" field.
func DeploymentIDNotIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldDeploymentID, vs...))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldOrgID, vs...))
}

// OrgIDGT applies the GT predicate on the "org_id" field.
func OrgIDGT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldOrgID, v))
}

// OrgIDGTE applies the GTE predicate on the "org_id" field.
func OrgIDGTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldOrgID, v))
}

// OrgIDLT applies the LT predicate on the "org_id" field.
func OrgIDLT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldOrgID, v))
}

// OrgIDLTE applies the LTE predicate on the "org_id" field.
func OrgIDLTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldOrgID, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v int) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldAppID, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryIsNil applies the IsNil predicate on the "category" field.
func CategoryIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldCategory))
}

// CategoryNotNil applies the NotNil predicate on the "category" field.
func CategoryNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldCategory))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContainsFold(FieldCategory, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContainsFold(FieldName, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContainsFold(FieldKey, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldVersion, v))
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldRevision, v))
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldRevision, v))
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldRevision, vs...))
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldRevision, vs...))
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldRevision, v))
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldRevision, v))
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldRevision, v))
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int32) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldRevision, v))
}

// RevisionIsNil applies the IsNil predicate on the "revision" field.
func RevisionIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldRevision))
}

// RevisionNotNil applies the NotNil predicate on the "revision" field.
func RevisionNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldRevision))
}

// ResourceNameEQ applies the EQ predicate on the "resource_name" field.
func ResourceNameEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldResourceName, v))
}

// ResourceNameNEQ applies the NEQ predicate on the "resource_name" field.
func ResourceNameNEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldResourceName, v))
}

// ResourceNameIn applies the In predicate on the "resource_name" field.
func ResourceNameIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldResourceName, vs...))
}

// ResourceNameNotIn applies the NotIn predicate on the "resource_name" field.
func ResourceNameNotIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldResourceName, vs...))
}

// ResourceNameGT applies the GT predicate on the "resource_name" field.
func ResourceNameGT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldResourceName, v))
}

// ResourceNameGTE applies the GTE predicate on the "resource_name" field.
func ResourceNameGTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldResourceName, v))
}

// ResourceNameLT applies the LT predicate on the "resource_name" field.
func ResourceNameLT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldResourceName, v))
}

// ResourceNameLTE applies the LTE predicate on the "resource_name" field.
func ResourceNameLTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldResourceName, v))
}

// ResourceNameContains applies the Contains predicate on the "resource_name" field.
func ResourceNameContains(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContains(FieldResourceName, v))
}

// ResourceNameHasPrefix applies the HasPrefix predicate on the "resource_name" field.
func ResourceNameHasPrefix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasPrefix(FieldResourceName, v))
}

// ResourceNameHasSuffix applies the HasSuffix predicate on the "resource_name" field.
func ResourceNameHasSuffix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasSuffix(FieldResourceName, v))
}

// ResourceNameIsNil applies the IsNil predicate on the "resource_name" field.
func ResourceNameIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldResourceName))
}

// ResourceNameNotNil applies the NotNil predicate on the "resource_name" field.
func ResourceNameNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldResourceName))
}

// ResourceNameEqualFold applies the EqualFold predicate on the "resource_name" field.
func ResourceNameEqualFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEqualFold(FieldResourceName, v))
}

// ResourceNameContainsFold applies the ContainsFold predicate on the "resource_name" field.
func ResourceNameContainsFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContainsFold(FieldResourceName, v))
}

// DgrmResourceNameEQ applies the EQ predicate on the "dgrm_resource_name" field.
func DgrmResourceNameEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldDgrmResourceName, v))
}

// DgrmResourceNameNEQ applies the NEQ predicate on the "dgrm_resource_name" field.
func DgrmResourceNameNEQ(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldDgrmResourceName, v))
}

// DgrmResourceNameIn applies the In predicate on the "dgrm_resource_name" field.
func DgrmResourceNameIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldDgrmResourceName, vs...))
}

// DgrmResourceNameNotIn applies the NotIn predicate on the "dgrm_resource_name" field.
func DgrmResourceNameNotIn(vs ...string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldDgrmResourceName, vs...))
}

// DgrmResourceNameGT applies the GT predicate on the "dgrm_resource_name" field.
func DgrmResourceNameGT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldDgrmResourceName, v))
}

// DgrmResourceNameGTE applies the GTE predicate on the "dgrm_resource_name" field.
func DgrmResourceNameGTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldDgrmResourceName, v))
}

// DgrmResourceNameLT applies the LT predicate on the "dgrm_resource_name" field.
func DgrmResourceNameLT(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldDgrmResourceName, v))
}

// DgrmResourceNameLTE applies the LTE predicate on the "dgrm_resource_name" field.
func DgrmResourceNameLTE(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldDgrmResourceName, v))
}

// DgrmResourceNameContains applies the Contains predicate on the "dgrm_resource_name" field.
func DgrmResourceNameContains(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContains(FieldDgrmResourceName, v))
}

// DgrmResourceNameHasPrefix applies the HasPrefix predicate on the "dgrm_resource_name" field.
func DgrmResourceNameHasPrefix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasPrefix(FieldDgrmResourceName, v))
}

// DgrmResourceNameHasSuffix applies the HasSuffix predicate on the "dgrm_resource_name" field.
func DgrmResourceNameHasSuffix(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldHasSuffix(FieldDgrmResourceName, v))
}

// DgrmResourceNameIsNil applies the IsNil predicate on the "dgrm_resource_name" field.
func DgrmResourceNameIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldDgrmResourceName))
}

// DgrmResourceNameNotNil applies the NotNil predicate on the "dgrm_resource_name" field.
func DgrmResourceNameNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldDgrmResourceName))
}

// DgrmResourceNameEqualFold applies the EqualFold predicate on the "dgrm_resource_name" field.
func DgrmResourceNameEqualFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEqualFold(FieldDgrmResourceName, v))
}

// DgrmResourceNameContainsFold applies the ContainsFold predicate on the "dgrm_resource_name" field.
func DgrmResourceNameContainsFold(v string) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldContainsFold(FieldDgrmResourceName, v))
}

// ResourceDataEQ applies the EQ predicate on the "resource_data" field.
func ResourceDataEQ(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldEQ(FieldResourceData, v))
}

// ResourceDataNEQ applies the NEQ predicate on the "resource_data" field.
func ResourceDataNEQ(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNEQ(FieldResourceData, v))
}

// ResourceDataIn applies the In predicate on the "resource_data" field.
func ResourceDataIn(vs ...[]byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIn(FieldResourceData, vs...))
}

// ResourceDataNotIn applies the NotIn predicate on the "resource_data" field.
func ResourceDataNotIn(vs ...[]byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotIn(FieldResourceData, vs...))
}

// ResourceDataGT applies the GT predicate on the "resource_data" field.
func ResourceDataGT(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGT(FieldResourceData, v))
}

// ResourceDataGTE applies the GTE predicate on the "resource_data" field.
func ResourceDataGTE(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldGTE(FieldResourceData, v))
}

// ResourceDataLT applies the LT predicate on the "resource_data" field.
func ResourceDataLT(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLT(FieldResourceData, v))
}

// ResourceDataLTE applies the LTE predicate on the "resource_data" field.
func ResourceDataLTE(v []byte) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldLTE(FieldResourceData, v))
}

// ResourceDataIsNil applies the IsNil predicate on the "resource_data" field.
func ResourceDataIsNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldIsNull(FieldResourceData))
}

// ResourceDataNotNil applies the NotNil predicate on the "resource_data" field.
func ResourceDataNotNil() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(sql.FieldNotNull(FieldResourceData))
}

// HasDeployment applies the HasEdge predicate on the "deployment" edge.
func HasDeployment() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeploymentWith applies the HasEdge predicate on the "deployment" edge with a given conditions (other predicates).
func HasDeploymentWith(preds ...predicate.Deployment) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeploymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeploymentTable, DeploymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDecisionDefs applies the HasEdge predicate on the "decision_defs" edge.
func HasDecisionDefs() predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DecisionDefsTable, DecisionDefsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDecisionDefsWith applies the HasEdge predicate on the "decision_defs" edge with a given conditions (other predicates).
func HasDecisionDefsWith(preds ...predicate.DecisionDef) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DecisionDefsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DecisionDefsTable, DecisionDefsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DecisionReqDef) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DecisionReqDef) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DecisionReqDef) predicate.DecisionReqDef {
	return predicate.DecisionReqDef(func(s *sql.Selector) {
		p(s.Not())
	})
}
