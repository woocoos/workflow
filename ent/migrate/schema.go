// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActDecisionDefColumns holds the columns for the "act_decision_def" table.
	ActDecisionDefColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deployment_id", Type: field.TypeInt},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "category", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "key", Type: field.TypeString},
		{Name: "req_def_key", Type: field.TypeString},
		{Name: "version", Type: field.TypeInt32},
		{Name: "revision", Type: field.TypeInt32, Nullable: true},
		{Name: "version_tag", Type: field.TypeString, Nullable: true},
		{Name: "resource_name", Type: field.TypeString, Nullable: true},
		{Name: "dgrm_resource_name", Type: field.TypeString, Nullable: true},
		{Name: "req_def_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActDecisionDefTable holds the schema information for the "act_decision_def" table.
	ActDecisionDefTable = &schema.Table{
		Name:       "act_decision_def",
		Columns:    ActDecisionDefColumns,
		PrimaryKey: []*schema.Column{ActDecisionDefColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_decision_def_act_decision_req_def_decision_defs",
				Columns:    []*schema.Column{ActDecisionDefColumns[17]},
				RefColumns: []*schema.Column{ActDecisionReqDefColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ActDecisionReqDefColumns holds the columns for the "act_decision_req_def" table.
	ActDecisionReqDefColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "category", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "key", Type: field.TypeString},
		{Name: "version", Type: field.TypeInt32},
		{Name: "revision", Type: field.TypeInt32, Nullable: true},
		{Name: "resource_name", Type: field.TypeString, Nullable: true},
		{Name: "dgrm_resource_name", Type: field.TypeString, Nullable: true},
		{Name: "resource_data", Type: field.TypeBytes, Nullable: true},
		{Name: "deployment_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActDecisionReqDefTable holds the schema information for the "act_decision_req_def" table.
	ActDecisionReqDefTable = &schema.Table{
		Name:       "act_decision_req_def",
		Columns:    ActDecisionReqDefColumns,
		PrimaryKey: []*schema.Column{ActDecisionReqDefColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_decision_req_def_act_deployment_decision_reqs",
				Columns:    []*schema.Column{ActDecisionReqDefColumns[15]},
				RefColumns: []*schema.Column{ActDeploymentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ActDeploymentColumns holds the columns for the "act_deployment" table.
	ActDeploymentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "source", Type: field.TypeString, Nullable: true},
		{Name: "deploy_time", Type: field.TypeTime},
	}
	// ActDeploymentTable holds the schema information for the "act_deployment" table.
	ActDeploymentTable = &schema.Table{
		Name:       "act_deployment",
		Columns:    ActDeploymentColumns,
		PrimaryKey: []*schema.Column{ActDeploymentColumns[0]},
	}
	// ActIdentityLinkColumns holds the columns for the "act_identity_link" table.
	ActIdentityLinkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "proc_def_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "assigner_id", Type: field.TypeInt, Nullable: true},
		{Name: "link_type", Type: field.TypeEnum, Enums: []string{"assignee", "candidate", "participant", "manager", "notifier"}},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "operation_type", Type: field.TypeEnum, Enums: []string{"add", "delete", "pass", "reject"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "task_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActIdentityLinkTable holds the schema information for the "act_identity_link" table.
	ActIdentityLinkTable = &schema.Table{
		Name:       "act_identity_link",
		Columns:    ActIdentityLinkColumns,
		PrimaryKey: []*schema.Column{ActIdentityLinkColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_identity_link_act_task_task_identities",
				Columns:    []*schema.Column{ActIdentityLinkColumns[9]},
				RefColumns: []*schema.Column{ActTaskColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ActProcDefColumns holds the columns for the "act_proc_def" table.
	ActProcDefColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "category", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "key", Type: field.TypeString},
		{Name: "version", Type: field.TypeInt32, Nullable: true},
		{Name: "revision", Type: field.TypeInt32, Nullable: true},
		{Name: "version_tag", Type: field.TypeString, Nullable: true},
		{Name: "resource_name", Type: field.TypeString, Nullable: true},
		{Name: "dgrm_resource_name", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "suspended"}, Default: "active"},
		{Name: "resource_data", Type: field.TypeBytes, Nullable: true},
		{Name: "deployment_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActProcDefTable holds the schema information for the "act_proc_def" table.
	ActProcDefTable = &schema.Table{
		Name:       "act_proc_def",
		Columns:    ActProcDefColumns,
		PrimaryKey: []*schema.Column{ActProcDefColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_proc_def_act_deployment_proc_defs",
				Columns:    []*schema.Column{ActProcDefColumns[17]},
				RefColumns: []*schema.Column{ActDeploymentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ActProcInstColumns holds the columns for the "act_proc_inst" table.
	ActProcInstColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "app_id", Type: field.TypeInt},
		{Name: "business_key", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "duration", Type: field.TypeInt, Nullable: true},
		{Name: "start_user_id", Type: field.TypeInt},
		{Name: "supper_instance_id", Type: field.TypeInt, Nullable: true},
		{Name: "root_instance_id", Type: field.TypeInt, Nullable: true},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_reason", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ready", "active", "completed", "failed", "terminated", "suspended", "deleted"}},
		{Name: "proc_def_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActProcInstTable holds the schema information for the "act_proc_inst" table.
	ActProcInstTable = &schema.Table{
		Name:       "act_proc_inst",
		Columns:    ActProcInstColumns,
		PrimaryKey: []*schema.Column{ActProcInstColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_proc_inst_act_proc_def_proc_instances",
				Columns:    []*schema.Column{ActProcInstColumns[17]},
				RefColumns: []*schema.Column{ActProcDefColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ActTaskColumns holds the columns for the "act_task" table.
	ActTaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "proc_def_id", Type: field.TypeInt},
		{Name: "execution_id", Type: field.TypeString},
		{Name: "run_id", Type: field.TypeString, Nullable: true},
		{Name: "task_def_key", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "assignee", Type: field.TypeString, Nullable: true},
		{Name: "member_count", Type: field.TypeInt32},
		{Name: "unfinished_count", Type: field.TypeInt32},
		{Name: "agree_count", Type: field.TypeInt32, Default: 0},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"AND", "OR"}, Default: "OR"},
		{Name: "sequential", Type: field.TypeBool, Default: false},
		{Name: "org_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"created", "running", "finished", "canceled"}, Default: "created"},
		{Name: "proc_inst_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// ActTaskTable holds the schema information for the "act_task" table.
	ActTaskTable = &schema.Table{
		Name:       "act_task",
		Columns:    ActTaskColumns,
		PrimaryKey: []*schema.Column{ActTaskColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_task_act_proc_inst_tasks",
				Columns:    []*schema.Column{ActTaskColumns[17]},
				RefColumns: []*schema.Column{ActProcInstColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActDecisionDefTable,
		ActDecisionReqDefTable,
		ActDeploymentTable,
		ActIdentityLinkTable,
		ActProcDefTable,
		ActProcInstTable,
		ActTaskTable,
	}
)

func init() {
	ActDecisionDefTable.ForeignKeys[0].RefTable = ActDecisionReqDefTable
	ActDecisionDefTable.Annotation = &entsql.Annotation{
		Table: "act_decision_def",
	}
	ActDecisionReqDefTable.ForeignKeys[0].RefTable = ActDeploymentTable
	ActDecisionReqDefTable.Annotation = &entsql.Annotation{
		Table: "act_decision_req_def",
	}
	ActDeploymentTable.Annotation = &entsql.Annotation{
		Table: "act_deployment",
	}
	ActIdentityLinkTable.ForeignKeys[0].RefTable = ActTaskTable
	ActIdentityLinkTable.Annotation = &entsql.Annotation{
		Table: "act_identity_link",
	}
	ActProcDefTable.ForeignKeys[0].RefTable = ActDeploymentTable
	ActProcDefTable.Annotation = &entsql.Annotation{
		Table: "act_proc_def",
	}
	ActProcInstTable.ForeignKeys[0].RefTable = ActProcDefTable
	ActProcInstTable.Annotation = &entsql.Annotation{
		Table: "act_proc_inst",
	}
	ActTaskTable.ForeignKeys[0].RefTable = ActProcInstTable
	ActTaskTable.Annotation = &entsql.Annotation{
		Table: "act_task",
	}
}