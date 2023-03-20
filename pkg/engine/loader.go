package engine

import (
	"encoding/xml"
	"github.com/woocoos/workflow/pkg/spec"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"github.com/woocoos/workflow/pkg/spec/dmn"
	"io"
	"os"
)

type BpmnLoader struct {
	Id   string
	Name string
	Spec *bpmn.Definitions
}

func (ld *BpmnLoader) GetName() string {
	return ld.Name
}

func (ld *BpmnLoader) GetId() string {
	return ld.Id
}

func (ld *BpmnLoader) Load(data []byte) error {
	var definitions bpmn.Definitions
	if err := xml.Unmarshal(data, &definitions); err != nil {
		return err
	}
	ld.Spec = &definitions
	ld.Id = ld.Spec.Id
	return ld.Validator()
}

func (ld *BpmnLoader) LoadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	return ld.Load(b)
}

// Validator 一致化输入数据,由于bpmn的扩展性,有些输入源是一致的.可以提前处理
func (ld *BpmnLoader) Validator() error {
	return nil
}

func (ld *BpmnLoader) FindProcess(id string) *bpmn.Process {
	for _, process := range ld.Spec.Process {
		if process.Id == id {
			return &process
		}
	}
	return nil
}

func (ld *BpmnLoader) FindMessage(id string) *bpmn.Message {
	for _, msg := range ld.Spec.Messages {
		if msg.Id == id {
			return &msg
		}
	}
	return nil
}

func (ld *BpmnLoader) FindSignal(ref string) *bpmn.Signal {
	for _, signal := range ld.Spec.Signals {
		if signal.Id == ref {
			return &signal
		}
	}
	return nil
}

func findSequenceFlows(sequenceFlows []*bpmn.SequenceFlow, ids []string) (vals []*bpmn.SequenceFlow) {
	for _, flow := range sequenceFlows {
		for _, id := range ids {
			if id == flow.Id {
				vals = append(vals, flow)
			}
		}
	}
	return vals
}

func flowFilter(flows []*bpmn.SequenceFlow, inst *InstanceRequest) (ret []*bpmn.SequenceFlow, err error) {
	for _, flow := range flows {
		if flow.ConditionExpression != nil {
			out, err := flow.ConditionExpression.Evaluate(inst.Variables)
			if err != nil {
				return nil, err
			}
			if out == true {
				ret = append(ret, flow)
			}
		}
	}
	if len(ret) == 0 {
		ret = append(ret, flowDefault(flows)...)
	}
	return ret, nil
}

func flowDefault(flows []*bpmn.SequenceFlow) (ret []*bpmn.SequenceFlow) {
	for _, flow := range flows {
		if flow.ConditionExpression == nil || flow.ConditionExpression.Content == "" {
			ret = append(ret, flow)
			break
		}
	}
	return ret
}

//--------------------------------------------------------------------------------------------

var _ spec.Loader = (*DMNLoader)(nil)

type DMNLoader struct {
	Id   string
	Spec *dmn.Definitions
}

func NewDMNLoader(id string) *DMNLoader {
	return &DMNLoader{
		Id: id,
	}
}

func (dl *DMNLoader) GetName() string {
	return dl.Spec.Name
}

func (dl *DMNLoader) GetId() string {
	return dl.Id
}

func (dl *DMNLoader) Load(data []byte) error {
	var definitions dmn.Definitions
	if err := xml.Unmarshal(data, &definitions); err != nil {
		return err
	}
	dl.Spec = &definitions
	dl.Id = dl.Spec.Id
	return dl.Validator()
}

func (dl *DMNLoader) LoadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	return dl.Load(b)
}

// Validator 一致化输入数据,由于bpmn的扩展性,有些输入源是一致的.可以提前处理
func (dl *DMNLoader) Validator() error {
	return nil
}

func (dl *DMNLoader) FindDecision(id string) *dmn.Decision {
	for _, decision := range dl.Spec.Decisions {
		if decision.Id == id {
			return &decision
		}
	}
	return nil
}
