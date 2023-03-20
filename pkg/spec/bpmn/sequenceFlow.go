package bpmn

type SequenceFlow struct {
	FlowElement
	ConditionExpression *ConditionExpression `xml:"conditionExpression"`
	SourceRef           string               `xml:"sourceRef,attr"`
	TargetRef           string               `xml:"targetRef,attr"`
	IsImmediate         bool                 `xml:"isImmediate,attr"`
}

func (s SequenceFlow) GetID() string {
	return s.Id
}
