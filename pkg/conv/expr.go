package conv

import (
	"github.com/antonmedv/expr"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/pkg/spec/dmn"
	"regexp"
	"strings"
	"unicode"
)

var ruleCache = make(map[string][]string)

// DefaultConverter use expr as evaluate expression engine
type DefaultConverter struct {
}

func (d *DefaultConverter) Match(rule string, input bpmn.Mappings) (bool, error) {
	v, err := Eval(rule, input, expr.AsBool())
	if err != nil {
		return false, err
	}
	return v.(bool), nil
}

func (d *DefaultConverter) Eval(rule string, input bpmn.Mappings) (any, error) {
	if strings.HasPrefix(rule, "=") {
		rule = strings.TrimPrefix(rule, "=")
	}
	return Eval(rule, input)
}

func NewDefaultConverter() *DefaultConverter {
	return &DefaultConverter{}
}

func (d *DefaultConverter) ConvertDecisionTable(decisionID string, dt *dmn.DecisionTable) ([]string, error) {
	key := decisionID + dt.Id
	pres, ok := ruleCache[key]
	if !ok {
		pres = make([]string, len(dt.Rules))
		for i, rule := range dt.Rules {
			pre := ""
			for i, input := range rule.InputEntries {
				converted, err := d.convertDecisionRule(input, dt.Inputs[i])
				if err != nil {
					return nil, err
				}
				if pre != "" {
					pre += " && "
				}
				pre += converted
			}
			pres[i] = pre
		}
		ruleCache[key] = pres
	}
	return pres, nil
}

var (
	stringListPattern = regexp.MustCompile(`"[^"]+"`)
)

func (d *DefaultConverter) convertDecisionRule(tests dmn.UnaryTests, clause dmn.InputClause) (string, error) {
	exp := tests.Text
	if exp == "" {
		return "", nil
	}
	name := clause.InputExpression.Text
	switch clause.InputExpression.TypeRef {
	case dmn.TypeRefString:
		if strings.HasPrefix(exp, "\"") {
			matches := stringListPattern.FindAllString(exp, -1)
			if len(matches) > 1 {
				return name + " in [" + strings.Join(matches, ",") + "]", nil
			}
			return name + " == " + exp, nil
		}
	case dmn.TypeRefNumber:
		if strings.Contains(exp, "..") {
			return name + " in " + exp, nil
		}
		// if start with number
		if unicode.IsDigit(rune(exp[0])) {
			return name + " == " + exp, nil
		}
	}
	return name + " " + exp, nil
}
