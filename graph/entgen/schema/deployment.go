package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"time"
)

// Deployment 流程发布.
//
// 流程是应用的扩展功能.发布的流程图,组织可以有不同的流程定义.
type Deployment struct {
	ent.Schema
}

func (Deployment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_deployment"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate(), entgql.MutationCreate()),
	}
}

func (Deployment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Deployment.
func (Deployment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Immutable().Comment("所属根组织ID").Annotations(entgql.Type("ID")),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("name").Optional().Comment("名称"),
		field.String("source").Optional().Comment("来源"),
		field.Time("deploy_time").Default(time.Now).Immutable().Comment("部署时间"),
	}
}

// Edges of the Deployment.
func (Deployment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proc_defs", ProcDef.Type),
		edge.To("decision_reqs", DecisionReqDef.Type),
	}
}
