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
	"github.com/woocoos/entco/schemax/typex"
	gen "github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/ent/intercept"
	"github.com/woocoos/workflow/ent/procdef"
)

// ProcDef 流程定义,对应流程文件中的一个process.
//
// 建议一个流程文件中为一个Process.流程定义中的图表信息未支持,所以目前dgrm_resource_name都是空.
type ProcDef struct {
	ent.Schema
}

func (ProcDef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_proc_def"},
		schemax.TenantField("tenant_id"),
	}
}

func (ProcDef) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the ProcDef.
func (ProcDef) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deployment_id").Immutable().Comment("部署ID"),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("category").Optional().Comment("分类"),
		field.String("name").Optional().Comment("名称"),
		field.String("key").Comment("process id"),
		field.Int32("version").Optional().Comment("版本"),
		field.Int32("revision").Optional().Comment("小版本"),
		field.String("version_tag").Optional().Comment("版本标签"),
		field.String("resource_key").Optional().Comment("流程文件key").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int("resource_id").Optional().Comment("流程文件id").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Comment("状态"),
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
					procdef.TenantID(mx.GetOrgID()), procdef.AppID(mx.GetAppID()),
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
