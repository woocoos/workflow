package api

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/pkg/spec/vars"
)

type (
	// Exporter 定义了外部系统接入工作流所需要实现的操作.
	Exporter interface {
		// GetProcDef get process definition by key
		GetProcDef(context.Context, *GetProcDefRequest) (*ent.ProcDef, error)
		// GetDecisionReqDef get decision requirement definition by key
		GetDecisionReqDef(context.Context, *GetDecisionReqDefRequest) (*ent.DecisionReqDef, error)
		// CreateUserTask create user task
		CreateUserTask(context.Context, *InstanceRequest, *bpmn.UserTask) (*ent.Task, error)
		// ClaimUserTask claim user task, if user task is claimed or error, return false
		ClaimUserTask(context.Context, *ent.CreateIdentityLinkInput) (added bool, err error)
		// Review approve or reject user task
		Review(ctx context.Context, input *ent.UpdateIdentityLinkInput, lastUndone int) error
		// GetUserIDs get user id by name. BPMN designer usually use username to identify user. the function convert username to id
		// if user not exist return empty slice.
		GetUserIDs(context.Context, *GetUserIDsRequest) (*GetUserIDResponse, error)
		// GetGroupIDs the function convert group name to id
		// if group not exist return empty slice.
		GetGroupIDs(context.Context, *GetGroupIDsRequest) (*GetGroupIDResponse, error)
	}
	// GetProcDefRequest is for exporter GetProcDef
	GetProcDefRequest struct {
		TenantID   int // tenant id plus
		ProcDefKey string
	}
	// GetDecisionReqDefRequest is for exporter GetDecisionReqDef
	GetDecisionReqDefRequest struct {
		OrgID          int
		DecisionDefKey string
	}
	// GetUserIDsRequest is for exporter GetUserIDs
	GetUserIDsRequest struct {
		OrgID int
		Names []string
	}
	// GetUserIDResponse is for exporter GetUserIDs
	GetUserIDResponse struct {
		UserIDs []int
	}
	// GetGroupIDsRequest is for exporter GetGroupIDs
	GetGroupIDsRequest struct {
		OrgID int
		Names []string
	}
	// GetGroupIDResponse is for exporter GetGroupIDs
	GetGroupIDResponse struct {
		GroupIDs []int
	}
	// GetGroupUserRequest is for exporter GetGroupUser
	GetGroupUserRequest struct {
		TenantID int
		Name     string
	}
	// GetGroupUserResponse is for exporter GetGroupUser
	GetGroupUserResponse struct {
		UserIDs []int
	}
)

// InstanceRequest 流程启动请求参数,流程请求序列化为json格式,传递给流程引擎.
//
// 嵌入了ent.ProcInst.
type InstanceRequest struct {
	*ent.ProcInst
	ProcDefKey string `json:"proc_def_key"`
	// ResourceKey 流程图资源路径,相对于流程图资源根目录
	ResourceKey string `json:"resource_name"`
	// 参数,用于在流程中传递业务数据.业务需要转化为map形式.
	Variables vars.Mapping `json:"variables"`
	// ResourcePath 流程图资源路径.测试时使用该参数简化数据初始化
	ResourcePath string `json:"-"`
}

type ServiceTaskActivityRequest struct {
	InstanceRequest *InstanceRequest
	Element         *bpmn.ServiceTask
}

type UserTaskActivityRequest struct {
	InstanceRequest *InstanceRequest
	Element         *bpmn.UserTask
}

type BusinessRuleTaskActivityRequest struct {
	InstanceRequest *InstanceRequest
	Element         *bpmn.BusinessRuleTask
}
