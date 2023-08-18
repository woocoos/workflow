package bpmn

import (
	"github.com/woocoos/workflow/pkg/spec/dmn"
	"github.com/woocoos/workflow/pkg/spec/vars"
)

var Convert Converter

// Converter convert dmn decision feel expression to other grammar
type Converter interface {
	// ConvertDecisionTable convert decision table to other grammar
	ConvertDecisionTable(decisionID string, table *dmn.DecisionTable) ([]string, error)
	// Match bool evaluate
	Match(rule string, input vars.Mapping) (bool, error)
	Eval(rule string, input vars.Mapping) (any, error)
}
