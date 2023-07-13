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

type OrgUser struct {
	ent.Schema
}

func (OrgUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_user"},
	}
}

func (OrgUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
	}
}

// Fields of the OrgUser.
func (OrgUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Time("joined_at").Default(time.Now).Comment("加入时间"),
		field.String("display_name").Comment("在组织内的显示名称"),
	}
}

// Edges of the OrgUser.
func (OrgUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org_roles", OrgRole.Type).Ref("org_users").
			Through("org_role_user", OrgRoleUser.Type).
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
