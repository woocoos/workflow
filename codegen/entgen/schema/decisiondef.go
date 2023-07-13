package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	gen "github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/intercept"
)

// DecisionDef holds the schema definition for the DecisionDef entity.
type DecisionDef struct {
	ent.Schema
}

func (DecisionDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_decision_def"},
		schemax.TenantField("tenant_id"),
	}
}

func (DecisionDef) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the DecisionDef.
func (DecisionDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deployment_id").Immutable().Comment("部署ID"),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.Int("req_def_id").Comment("决策定义ID"),
		field.String("category").Optional().Comment("分类"),
		field.String("name").Optional().Comment("名称"),
		field.String("key").Comment("KEY"),
		field.String("req_def_key").Comment("决策定义key"),
		field.Int32("version").Comment("版本"),
		field.Int32("revision").Optional().Comment("小版本"),
		field.String("version_tag").Optional().Comment("版本标签"),
	}
}

// Edges of the DecisionDef.
func (DecisionDef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("req_def", DecisionReqDef.Type).Ref("decision_defs").Unique().Required().Field("req_def_id"),
	}
}
