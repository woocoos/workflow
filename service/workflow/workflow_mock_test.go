package workflow

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/engine"
	"path/filepath"
)

var _ engine.Exporter = (*FileExport)(nil)

type FileExport struct {
	DbExport
}

func (e *FileExport) GetProcDef(ctx context.Context, req *engine.GetProcDefRequest) (*ent.ProcDef, error) {
	return &ent.ProcDef{
		Key:         req.ProcDefKey,
		ResourceKey: filepath.Join("case", req.ProcDefKey+".bpmn"),
	}, nil
}

func (e *FileExport) GetDecisionReqDef(ctx context.Context, req *engine.GetDecisionReqDefRequest) (*ent.DecisionReqDef, error) {
	return &ent.DecisionReqDef{
		Key:         req.DecisionDefKey,
		ResourceKey: filepath.Join("case", "decisions.dmn"),
	}, nil
}

func (e *FileExport) GetUserIDs(ctx context.Context, req *engine.GetUserIDsRequest) (*engine.GetUserIDResponse, error) {
	ids := make([]int, 0, len(req.Names))
	for i, _ := range req.Names {
		ids[i] = i + 1
	}
	return &engine.GetUserIDResponse{
		UserIDs: ids,
	}, nil
}

func (e *FileExport) GetGroupIDs(ctx context.Context, req *engine.GetGroupIDsRequest) (*engine.GetGroupIDResponse, error) {
	ids := make([]int, 0, len(req.Names))
	for i, _ := range req.Names {
		ids[i] = i + 1
	}
	return &engine.GetGroupIDResponse{
		GroupIDs: ids,
	}, nil
}
