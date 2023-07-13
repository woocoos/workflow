package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
)

type OrgApp struct {
	ent.Schema
}

func (OrgApp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_app"},
		entgql.Skip(entgql.SkipAll),
	}
}

func (OrgApp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
	}
}

// Fields of the OrgApp.
func (OrgApp) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("app_id").Comment("应用ID"),
	}
}
