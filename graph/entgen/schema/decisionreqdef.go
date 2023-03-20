package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
)

// DecisionReqDef holds the schema definition for the DecisionReqDef entity.
type DecisionReqDef struct {
	ent.Schema
}

func (DecisionReqDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_decision_req_def"},
	}
}

func (DecisionReqDef) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the DecisionReqDef.
func (DecisionReqDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deployment_id").Immutable().Comment("部署ID"),
		field.Int("org_id").Immutable().Comment("所属根组织ID").Annotations(entgql.Type("ID")),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("category").Optional().Comment("分类"),
		field.String("name").Optional().Comment("名称"),
		field.String("key").Comment("KEY"),
		field.Int32("version").Comment("版本"),
		field.Int32("revision").Optional().Comment("小版本"),
		field.String("resource_name").Optional().Comment("资源名称"),
		field.String("dgrm_resource_name").Optional().Comment("流程图资源名称"),
		field.Bytes("resource_data").Optional().Comment("资源数据").Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Edges of the DecisionReqDef.
func (DecisionReqDef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).Ref("decision_reqs").Unique().Immutable().Required().Field("deployment_id"),
		edge.To("decision_defs", DecisionDef.Type),
	}
}
