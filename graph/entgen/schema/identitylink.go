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

// IdentityLink holds the schema definition for the IdentityLink entity.
type IdentityLink struct {
	ent.Schema
}

func (IdentityLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_identity_link"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (IdentityLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
	}
}

// Fields of the IdentityLink.
func (IdentityLink) Fields() []ent.Field {
	return []ent.Field{
		field.Int("task_id").Comment("流程实例ID"),
		field.Int("proc_def_id").Immutable().Comment("流程定义ID").Annotations(entgql.Type("ID")),
		field.Int("group_id").Optional().Comment("组ID").Annotations(entgql.Type("ID")),
		field.Int("user_id").Optional().Comment("用户ID").Annotations(entgql.Type("ID")),
		field.Int("assigner_id").Optional().Comment("分配人").Annotations(entgql.Type("ID")),
		field.Enum("link_type").Values("assignee", "candidate", "participant", "manager", "notifier").
			Comment("分配,候选,参与,上级,抄送"),
		field.Int("org_id").Immutable().Comment("组织ID").Annotations(entgql.Type("ID")),
		field.Enum("operation_type").Values("add", "delete", "pass", "reject").Comment("操作类型"),
		field.String("comments").Optional().Comment("评论"),
	}
}

// Edges of the IdentityLink.
func (IdentityLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("task_identities").Unique().Required().Field("task_id").
			Comment("流程任务"),
	}
}
