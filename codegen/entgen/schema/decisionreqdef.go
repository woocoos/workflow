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

// DecisionReqDef 决策文件需求定义.
type DecisionReqDef struct {
	ent.Schema
}

func (DecisionReqDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_decision_req_def"},
		schemax.TenantField("tenant_id"),
	}
}

func (DecisionReqDef) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the DecisionReqDef.
func (DecisionReqDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deployment_id").Immutable().Comment("部署ID"),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("category").Optional().Comment("分类"),
		field.String("name").Optional().Comment("名称"),
		field.String("key").Comment("KEY"),
		field.Int32("version").Comment("版本"),
		field.Int32("revision").Optional().Comment("小版本"),
		field.String("resource_key").Optional().Comment("流程文件key").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int("resource_id").Optional().Comment("流程文件id").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the DecisionReqDef.
func (DecisionReqDef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).Ref("decision_reqs").Unique().Immutable().Required().Field("deployment_id"),
		edge.To("decision_defs", DecisionDef.Type),
	}
}
