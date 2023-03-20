package workflow

import (
	"context"
	"github.com/woocoos/workflow/ent"
	"github.com/woocoos/workflow/pkg/engine"
	"github.com/woocoos/workflow/test"
	"os"
	"path/filepath"
)

var _ engine.Exporter = (*FileExport)(nil)

type FileExport struct {
	DbExport
}

func (e *FileExport) GetProcDef(ctx context.Context, req *engine.GetProcDefRequest) (*ent.ProcDef, error) {
	file := filepath.Join(test.BaseDir(), "case", req.ProcDefKey+".bpmn")
	bs, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &ent.ProcDef{
		Key:          req.ProcDefKey,
		ResourceName: filepath.Base(file),
		ResourceData: bs,
	}, nil
}

func (e *FileExport) GetDecisionReqDef(ctx context.Context, req *engine.GetDecisionReqDefRequest) (*ent.DecisionReqDef, error) {
	file := filepath.Join(test.BaseDir(), "case", "decisions.dmn")
	bs, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return &ent.DecisionReqDef{
		Key:          req.DecisionDefKey,
		ResourceName: "decisions.dmn",
		ResourceData: bs,
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
