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

// OrgRoleUser holds the schema definition for the OrgRoleUser entity.
type OrgRoleUser struct {
	ent.Schema
}

func (OrgRoleUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_role_user"},
		entgql.Skip(entgql.SkipAll),
	}
}

func (OrgRoleUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
	}
}

// Fields of the OrgRoleUser.
func (OrgRoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_role_id").Comment("组织角色ID"),
		field.Int("org_user_id").Comment("组织用户ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Int("org_id").Comment("组织ID"),
	}
}

// Edges of the OrgRoleUser.
func (OrgRoleUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("org_role", OrgRole.Type).Unique().Required().
			Field("org_role_id").Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("org_user", OrgUser.Type).Unique().Required().
			Field("org_user_id").Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
