package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	gen "github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/procdef"
)

// ProcDef holds the schema definition for the ProcDef entity.
type ProcDef struct {
	ent.Schema
}

func (ProcDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_proc_def"},
	}
}

func (ProcDef) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the ProcDef.
func (ProcDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deployment_id").Immutable().Comment("部署ID"),
		field.Int("org_id").Immutable().Comment("所属根组织ID").Annotations(entgql.Type("ID")),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("category").Optional().Comment("分类"),
		field.String("name").Optional().Comment("名称"),
		field.String("key").Comment("KEY"),
		field.Int32("version").Optional().Comment("版本"),
		field.Int32("revision").Optional().Comment("小版本"),
		field.String("version_tag").Optional().Comment("版本标签"),
		field.String("resource_name").Optional().Comment("资源名称"),
		field.String("dgrm_resource_name").Optional().Comment("流程图资源名称"),
		field.Enum("status").Values("active", "suspended").Default("active").Comment("状态"),
		field.Bytes("resource_data").Optional().Comment("资源数据").Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Edges of the ProcDef.
func (ProcDef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deployment", Deployment.Type).Ref("proc_defs").Unique().Immutable().Required().Field("deployment_id"),
		edge.To("proc_instances", ProcInst.Type).Annotations(entgql.RelayConnection()),
	}
}

func (ProcDef) Hooks() []ent.Hook {
	return []ent.Hook{
		versionHook,
	}
}

// 同组织下,同应用下,同key的流程定义,版本号+1
func versionHook(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if m.Op() == ent.OpCreate {
			if mx, ok := m.(interface {
				SetVersion(int32)
				SetRevision(int32)
				Client() *gen.Client
				GetVersion() int32
				GetOrgID() int
				GetAppID() int
				GetKey() string
			}); ok {
				max, err := mx.Client().ProcDef.Query().Where(procdef.KeyEQ(mx.GetKey()),
					procdef.OrgID(mx.GetOrgID()), procdef.AppID(mx.GetAppID()),
				).Aggregate(gen.Max(procdef.FieldVersion)).Int(ctx)
				if err != nil {
					return nil, err
				}
				mx.SetVersion(int32(max + 1))
			}
		}
		return next.Mutate(ctx, m)
	})
}
