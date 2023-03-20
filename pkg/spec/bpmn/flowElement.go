package bpmn

type FlowElement struct {
	BaseElement
	Name string `xml:"name,attr"`
}

type FlowNode struct {
	FlowElement
	Tag      string   `xml:"tag,attr"`
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`
}

func (f FlowNode) GetId() string {
	return f.Id
}

func (f FlowNode) GetOutgoing() []string {
	return f.Outgoing
}
