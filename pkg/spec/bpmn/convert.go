package bpmn

import "github.com/woocoos/workflow/pkg/spec/dmn"

var Convert Converter

// Converter convert dmn decision feel expression to other grammar
type Converter interface {
	// ConvertDecisionTable convert decision table to other grammar
	ConvertDecisionTable(decisionID string, table *dmn.DecisionTable) ([]string, error)
	// Match bool evaluate
	Match(rule string, input Mappings) (bool, error)
	Eval(rule string, input Mappings) (any, error)
}

type Mappings map[string]any

func (Mappings) Equals(a, b interface{}) bool {
	return a == b
}
