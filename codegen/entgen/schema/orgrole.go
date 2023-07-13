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

// OrgRole holds the schema definition for the OrgRole entity.
type OrgRole struct {
	ent.Schema
}

func (OrgRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_role"},
	}
}

func (OrgRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
	}
}

func (OrgRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("组织ID"),
		field.Enum("kind").Values("group", "role").Comment("类型,group:组,role:角色"),
		field.String("name").Comment("名称"),
	}
}

// Edges of the OrgRole.
func (OrgRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("org_users", OrgUser.Type).Comment("组用户").
			Through("org_role_user", OrgRoleUser.Type).Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
