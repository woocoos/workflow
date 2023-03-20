package engine

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
)

type (
	// GetProcDefRequest is for exporter GetProcDef
	GetProcDefRequest struct {
		OrgID      int // tenant id plus
		ProcDefKey string
	}
	GetDecisionReqDefRequest struct {
		OrgID          int
		DecisionDefKey string
	}
	GetUserIDsRequest struct {
		OrgID int
		Names []string
	}
	GetUserIDResponse struct {
		UserIDs []int
	}
	GetGroupIDsRequest struct {
		OrgID int
		Names []string
	}
	GetGroupIDResponse struct {
		GroupIDs []int
	}
	// Exporter is an interface for export data or handle data
	Exporter interface {
		// GetProcDef get process definition by key
		GetProcDef(context.Context, *GetProcDefRequest) (*ent.ProcDef, error)
		GetDecisionReqDef(context.Context, *GetDecisionReqDefRequest) (*ent.DecisionReqDef, error)
		CreateUserTask(context.Context, *InstanceRequest, *bpmn.UserTask) (*ent.Task, error)
		ClaimUserTask(context.Context, *ent.CreateIdentityLinkInput) error
		Review(ctx context.Context, input *ent.UpdateIdentityLinkInput, lastUndone int) error
		// GetUserIDs get user id by name. BPMN designer usually use username to identify user. the function convert username to id
		GetUserIDs(context.Context, *GetUserIDsRequest) (*GetUserIDResponse, error)
		// GetGroupIDs the function convert group name to id
		GetGroupIDs(context.Context, *GetGroupIDsRequest) (*GetGroupIDResponse, error)
	}
)

type InstanceRequest struct {
	*ent.ProcInst
	ProcDefKey   string        `json:"proc_def_key"`
	ResourceName string        `json:"resource_name"`
	Variables    bpmn.Mappings `json:"variables"`
	ResourceData []byte        `json:"resource_data"`
	ResourcePath string        `json:"resource_path"`
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
