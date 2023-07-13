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
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_task"},
		entgql.QueryField(),
		entgql.RelayConnection(),
		schemax.TenantField("tenant_id"),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.NewTenantMixin[intercept.Query, *gen.Client](intercept.NewQuery),
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("proc_inst_id").Comment("流程实例ID"),
		field.Int("proc_def_id").Comment("流程定义ID"),
		field.String("execution_id").Comment("workflow id"),
		field.String("run_id").Optional().Comment("run id"),
		field.String("task_def_key").Comment("流程节点名称"),
		field.Int("parent_id").Optional().Default(0).Comment("父任务ID"),
		field.String("comments").Optional().Comment("任务描述").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("assignee").Optional().Comment("设定的受理人"),
		field.Int32("member_count").Comment("任务成员人数"),
		field.Int32("unfinished_count").Comment("未任务成员完成人数"),
		field.Int32("agree_count").Default(0).Comment("通过数量"),
		field.Enum("kind").Comment("会签类型").Values("AND", "OR").Default("OR"),
		field.Bool("sequential").Default(false).Comment("默认并行false,顺序执行true"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Immutable().Comment("创建时间"),
		field.Enum("status").Comment("任务状态").Values("created", "running", "finished", "canceled").
			Default("created"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("proc_inst", ProcInst.Type).Ref("tasks").Unique().Required().Field("proc_inst_id").Comment("流程实例"),
		edge.To("task_identities", IdentityLink.Type).Comment("任务主体"),
	}
}
