package engine

import (
	"fmt"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
)

func CompletionCondition(input string, total, undone, agree int32) (bool, error) {
	output, err := bpmn.Convert.Eval(input, map[string]any{
		"numberOfInstances":           total,
		"numberOfActiveInstances":     total - undone,
		"numberOfCompletedInstances":  agree,
		"numberOfTerminatedInstances": total - undone - agree,
	})
	if err != nil {
		return false, err
	}
	if v, ok := output.(bool); ok {
		return v, nil
	}
	return false, fmt.Errorf("completion condition is not a bool expression: %v", output)
}
