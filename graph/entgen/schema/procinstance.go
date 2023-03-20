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

// ProcInst holds process instance
type ProcInst struct {
	ent.Schema
}

func (ProcInst) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "act_proc_inst"},
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

func (ProcInst) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the ProcInst.
func (ProcInst) Fields() []ent.Field {
	return []ent.Field{
		field.Int("proc_def_id").Comment("流程定义ID"),
		field.Int("org_id").Immutable().Comment("所属根组织ID").Annotations(entgql.Type("ID")),
		field.Int("app_id").Immutable().Comment("所属应用ID").Annotations(entgql.Type("ID")),
		field.String("business_key").Comment("业务主键"),
		field.Time("start_time").Default(time.Now).Comment("开始时间"),
		field.Time("end_time").Optional().Comment("结束时间"),
		field.Int("duration").Optional().Comment("持续时间"),
		field.Int("start_user_id").Comment("发起人ID/名称").Annotations(entgql.Type("ID")),
		field.Int("supper_instance_id").Optional().Comment("父流程实例ID"),
		field.Int("root_instance_id").Optional().Comment("根流程实例ID"),
		field.Time("deleted_time").Optional().Comment("删除时间"),
		field.String("deleted_reason").Optional().Comment("删除原因"),
		field.Enum("status").Values("ready", "active", "completed", "failed", "terminated", "suspended", "deleted").Comment("状态"),
	}
}

// Edges of the ProcInst.
func (ProcInst) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("proc_def", ProcDef.Type).Ref("proc_instances").Unique().Required().Field("proc_def_id").Comment("流程定义"),
		edge.To("tasks", Task.Type).Comment("任务列表").Annotations(entgql.RelayConnection()),
	}
}
