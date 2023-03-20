package bpmn

import (
	"fmt"
	"strconv"
	"strings"
)

type TaskDefinition struct {
	TypeName string `xml:"type,attr"`
	Retries  string `xml:"retries,attr"`
}

func (t TaskDefinition) GetRetries() (retries int, err error) {
	if t.Retries == "" {
		return
	}

	val, err := Convert.Eval(t.Retries, nil)
	if err != nil {
		return 0, fmt.Errorf("[TaskDefinition.GetRetries]failed to evaluate expression %s: %w", t.Retries, err)
	}
	switch v := val.(type) {
	case int:
		return v, nil
	case string:
		retries, err = strconv.Atoi(v)
		return
	}
	return
}

type FormDefinition struct {
	FormKey string `xml:"formKey,attr"`
}

type Header struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type Input struct {
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
}

type Output struct {
	Source string `xml:"source,attr" xml:"name,attr"`
	Target string `xml:"target,attr" xml:",innerxml"`
}

func (o Output) OutputTarget(src, tar Mappings) error {
	v, err := Convert.Eval(o.Source, src)
	if err != nil {
		return err
	}
	tar[o.Target] = v
	return nil
}

type AssignmentDefinition struct {
	Assignee        string `xml:"assignee,attr"`
	CandidateGroups string `xml:"candidateGroups,attr"`
	CandidateUsers  string `xml:"candidateUsers,attr"`
}

// CountAssignees count assignees
func (a AssignmentDefinition) CountAssignees() (count int) {
	if a.Assignee != "" {
		return 1
	}
	if a.CandidateUsers != "" {
		count += len(strings.Split(a.CandidateUsers, ","))
	}
	if a.CandidateGroups != "" {
		count += len(strings.Split(a.CandidateGroups, ","))
	}
	return
}

func (a AssignmentDefinition) GetAssignee(vars Mappings) (string, error) {
	v, err := Convert.Eval(a.Assignee, vars)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate assignee: %w", err)
	}
	ret, err := a.ConvertValue(v)
	if err != nil {
		return "", err
	}

	return ret[0], nil
}

func (a AssignmentDefinition) GetCandidateGroups(vars Mappings) ([]string, error) {
	v, err := Convert.Eval(a.CandidateGroups, vars)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate candidate groups: %w", err)
	}
	return a.ConvertValue(v)
}

func (a AssignmentDefinition) GetCandidateUsers(vars Mappings) ([]string, error) {
	v, err := Convert.Eval(a.CandidateUsers, vars)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate candidate users: %w", err)
	}
	return a.ConvertValue(v)
}

func (a AssignmentDefinition) ConvertValue(val any) ([]string, error) {
	switch val.(type) {
	case []string:
		return val.([]string), nil
	case string:
		return strings.Split(val.(string), ","), nil
	case []any:
		var res []string
		for _, v := range val.([]any) {
			for _, vv := range v.(map[string]any) {
				res = append(res, vv.(string))
			}
		}
		return res, nil
	}
	return nil, fmt.Errorf("unsupported type %T", val)
}

type ExtensionProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Subscription struct {
	CorrelationKey string `xml:"correlationKey,attr"`
}

type Script struct {
	Expression     string `xml:"expression,attr"`
	ResultVariable string `xml:"resultVariable,attr"`
}

type CalledDecision struct {
	DecisionId     string `xml:"decisionId,attr"`
	ResultVariable string `xml:"resultVariable,attr"`
}

type CalledElement struct {
	ProcessId                  string `xml:"processId,attr"`
	PropagateAllChildVariables bool   `xml:"propagateAllChildVariables,attr"`
}

// MultiInstanceLoopCharacteristics is a loop characteristics for multi instance
type MultiInstanceLoopCharacteristics struct {
	IsSequential        bool                `xml:"isSequential,attr"`
	LoopCharacteristics LoopCharacteristics `xml:"extensionElements>loopCharacteristics"`
	CompletionCondition *FormalExpression   `xml:"completionCondition"`
}

type LoopCharacteristics struct {
	InputCollection  string `xml:"inputCollection,attr"`
	InputElement     string `xml:"inputElement,attr"`
	OutputCollection string `xml:"outputCollection,attr"`
	OutputElement    string `xml:"outputElement,attr"`
}
