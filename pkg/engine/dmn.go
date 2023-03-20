package engine

import (
	"fmt"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/pkg/spec/dmn"
	"reflect"
	"strconv"
)

type DMN struct {
	Id string
}

func NewDMN(id string) *DMN {
	return &DMN{
		Id: id,
	}
}

// StartInstance start a dmn instance,return a single feel value
func (d *DMN) StartInstance(loader *DMNLoader, task *bpmn.BusinessRuleTask, pi *InstanceRequest) (bpmn.Mappings, error) {
	dcn := loader.FindDecision(task.CalledDecision.DecisionId)
	if dcn == nil {
		return nil, fmt.Errorf("[StartInstance]decision %s not found", task.CalledDecision.DecisionId)
	}
	output, err := d.EvaluateDecision(loader, dcn.Id, pi.Variables)
	if err != nil {
		return nil, err
	}
	resultDt := dcn.DecisionTable
	v, err := d.Result(resultDt, output)
	if err != nil {
		return nil, err
	}
	return bpmn.Mappings{
		task.CalledDecision.ResultVariable: v,
	}, nil
}

// Result handle the result of decision DRG
// dt is final decision table
func (d *DMN) Result(final *dmn.DecisionTable, input []bpmn.Mappings) (output any, err error) {
	switch final.HitPolicy {
	case dmn.HitPolicyCOLLECT, dmn.HitPolicyRULEORDER:
		return input, nil
	}
	tmp := make(bpmn.Mappings)
	if final.Aggregation != nil {
		tmp, err = aggregation(final, input)
		if err != nil {
			return nil, err
		}
	}
	if len(final.Outputs) == 1 {
		for _, v := range tmp { // only one output
			return v, nil
		}
	}
	return tmp, nil
}

func (d *DMN) EvaluateDecision(loader *DMNLoader, did string, ctxMap bpmn.Mappings) ([]bpmn.Mappings, error) {
	dcn := loader.FindDecision(did)
	if dcn == nil {
		return nil, fmt.Errorf("[EvaluateDecision]decision %s not found", did)
	}

	// choice
	if dcn.LiteralExpression != nil {
		if dcn.LiteralExpression.Text != "" {
			v, err := bpmn.Convert.Eval(dcn.LiteralExpression.Text, ctxMap)
			if err != nil {
				return nil, err
			}
			return []bpmn.Mappings{
				{dcn.Variable.Name: v},
			}, nil
		}
		return nil, nil
	}

	vars := []bpmn.Mappings{ctxMap}
	for _, rd := range dcn.InformationRequirement {
		did := rd.RequiredDecision.Href[1:]
		result, err := d.EvaluateDecision(loader, did, ctxMap)
		if err != nil {
			return nil, err
		}
		vars, err = cross(vars, result)
		if err != nil {
			return nil, err
		}
	}
	if dcn.DecisionTable == nil {
		return nil, fmt.Errorf("[EvaluateDecision]decision %s not found", did)
	}
	var output []bpmn.Mappings
	for _, mappings := range vars {
		data, err := d.EvaluateTable(did, dcn.DecisionTable, mappings)
		if err != nil {
			return nil, err
		}
		output = append(output, data...)
	}
	return output, nil
}

// EvaluateTable evaluate decision table
//
// IF THE COLLECT HIT POLICY IS USED WITH AN AGGREGATOR, THE DECISION TABLE CAN ONLY HAVE ONE OUTPUT.
func (d *DMN) EvaluateTable(did string, dt *dmn.DecisionTable, input bpmn.Mappings) (output []bpmn.Mappings, err error) {
	rules, err := bpmn.Convert.ConvertDecisionTable(did, dt)
	if err != nil {
		return nil, err
	}
	for i, rule := range rules {
		match, err := bpmn.Convert.Match(rule, input)
		if err != nil {
			return nil, err
		}
		if !match {
			continue
		}
		rp := bpmn.Mappings{}
		for j, column := range dt.Rules[i].OutputEntries {
			// literal expression
			p, err := bpmn.Convert.Eval(column.Text, nil)
			if err != nil {
				return nil, err
			}
			key := dt.Outputs[j].Name
			rp[key] = p
		}
		output = append(output, rp)
	}
	if len(output) < 2 {
		return output, nil
	}
	switch dt.HitPolicy {
	case dmn.HitPolicyUNIQUE:
		if len(output) > 1 {
			return nil, fmt.Errorf("hit policy not unique error:%s,%v", dt.Id, input)
		}
		return output[:1], nil
	case dmn.HitPolicyANY, dmn.HitPolicyFIRST:
		if len(output) > 0 {
			return output[:1], nil
		}
	case dmn.HitPolicyCOLLECT, dmn.HitPolicyRULEORDER:
		return output, nil
	}
	gr, err := aggregation(dt, output)
	if err != nil {
		return nil, err
	}
	return []bpmn.Mappings{gr}, nil
}

func aggregation(dt *dmn.DecisionTable, output []bpmn.Mappings) (bpmn.Mappings, error) {
	switch *dt.Aggregation {
	case dmn.BuiltinAggregatorSUM:
		return sum(output)
	case dmn.BuiltinAggregatorCOUNT:
		return count(output), nil
	case dmn.BuiltinAggregatorMIN:
		return min(output)
	case dmn.BuiltinAggregatorMAX:
		return max(output)
	}
	return nil, fmt.Errorf("hit policy not support error:%s,%s", dt.Id, *dt.Aggregation)
}

func cross(left, right []bpmn.Mappings) ([]bpmn.Mappings, error) {
	output := make([]bpmn.Mappings, 0, len(left)*len(right))
	for _, l := range left {
		for _, r := range right {
			for k, v := range r {
				l[k] = v
			}
			output = append(output, l)
		}
	}
	return output, nil
}

// cross product
func crossMap(left, right bpmn.Mappings) ([]bpmn.Mappings, error) {
	ls := make([]bpmn.Mappings, 0)
	for _, v := range left {
		if _, ok := v.(bpmn.Mappings); ok {
			ls = append(ls, v.(bpmn.Mappings))
		} else {
			ls = append(ls, left)
			break
		}
	}
	rs := make([]bpmn.Mappings, 0)
	for _, v := range right {
		if _, ok := v.(bpmn.Mappings); ok {
			rs = append(rs, v.(bpmn.Mappings))
		} else {
			rs = append(rs, right)
			break
		}
	}
	return cross(ls, rs)
}

// sum aggregation for Mappings
func sum(input []bpmn.Mappings) (bpmn.Mappings, error) {
	output := make(bpmn.Mappings)
	for _, v := range input {
		for k, p := range v {
			if _, ok := output[k]; !ok {
				output[k] = p
			} else {
				f, err := AsFloat(p.(float64))
				if err != nil {
					return nil, err
				}
				output[k] = output[k].(float64) + f
			}
		}
	}
	return output, nil
}

// max aggregation for Mappings
func max(input []bpmn.Mappings) (bpmn.Mappings, error) {
	output := make(bpmn.Mappings)
	for _, v := range input {
		for k, p := range v {
			if _, ok := output[k]; !ok {
				output[k] = p
			} else {
				f, err := AsFloat(p.(float64))
				if err != nil {
					return nil, err
				}
				if output[k].(float64) < f {
					output[k] = p
				}
			}
		}
	}
	return output, nil
}

// min aggregation for Mappings
func min(input []bpmn.Mappings) (bpmn.Mappings, error) {
	output := make(bpmn.Mappings)
	for _, v := range input {
		for k, p := range v {
			if _, ok := output[k]; !ok {
				output[k] = p
			} else {
				f, err := AsFloat(p.(float64))
				if err != nil {
					return nil, err
				}
				if output[k].(float64) > f {
					output[k] = p
				}
			}
		}
	}
	return output, nil
}

// count aggregation for Mappings
func count(input []bpmn.Mappings) bpmn.Mappings {
	output := make(bpmn.Mappings)
	for _, v := range input {
		for k, _ := range v {
			if _, ok := output[k]; !ok {
				output[k] = 1
			} else {
				output[k] = output[k].(int) + 1
			}
		}
	}
	return output
}

func AsStr(val interface{}) interface{} {
	return fmt.Sprintf("%v", val)
}

func AsInt(in interface{}) (int64, error) {
	switch i := in.(type) {
	case float64:
		return int64(i), nil
	case float32:
		return int64(i), nil
	case int64:
		return i, nil
	case int32:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case int:
		return int64(i), nil
	case uint64:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case uint16:
		return int64(i), nil
	case uint8:
		return int64(i), nil
	case uint:
		return int64(i), nil
	case string:
		inAsInt, err := strconv.ParseInt(i, 10, 64)
		if err == nil {
			return inAsInt, nil
		}
	}
	return 0, fmt.Errorf("asInt() not supported on %v %v", reflect.TypeOf(in), in)
}

func AsFloat(in interface{}) (float64, error) {
	switch i := in.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int8:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint16:
		return float64(i), nil
	case uint8:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		inAsFloat, err := strconv.ParseFloat(i, 64)
		if err == nil {
			return inAsFloat, nil
		}
	}
	return 0, fmt.Errorf("asFloat() not supported on %v %v", reflect.TypeOf(in), in)
}
