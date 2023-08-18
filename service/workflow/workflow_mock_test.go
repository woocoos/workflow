package workflow

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/api"
	"path/filepath"
)

var _ api.Exporter = (*FileExport)(nil)

type FileExport struct {
	DbExport
}

func (e *FileExport) GetProcDef(ctx context.Context, req *api.GetProcDefRequest) (*ent.ProcDef, error) {
	return &ent.ProcDef{
		Key:         req.ProcDefKey,
		ResourceKey: filepath.Join("case", req.ProcDefKey+".bpmn"),
	}, nil
}

func (e *FileExport) GetDecisionReqDef(ctx context.Context, req *api.GetDecisionReqDefRequest) (*ent.DecisionReqDef, error) {
	return &ent.DecisionReqDef{
		Key:         req.DecisionDefKey,
		ResourceKey: filepath.Join("case", "decisions.dmn"),
	}, nil
}

func (e *FileExport) GetUserIDs(ctx context.Context, req *api.GetUserIDsRequest) (*api.GetUserIDResponse, error) {
	ids := make([]int, 0, len(req.Names))
	for i, _ := range req.Names {
		ids[i] = i + 1
	}
	return &api.GetUserIDResponse{
		UserIDs: ids,
	}, nil
}

func (e *FileExport) GetGroupIDs(ctx context.Context, req *api.GetGroupIDsRequest) (*api.GetGroupIDResponse, error) {
	ids := make([]int, 0, len(req.Names))
	for i, _ := range req.Names {
		ids[i] = i + 1
	}
	return &api.GetGroupIDResponse{
		GroupIDs: ids,
	}, nil
}
