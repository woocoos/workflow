package bpmn

type Activity struct {
	FlowNode
	Default           string `xml:"default,attr"`
	IsForCompensation bool   `xml:"isForCompensation,attr"`
}

type CallActivity struct {
	Activity
	CalledElement CalledElement `xml:"extensionElements>calledElement"`
	Inputs        []Input       `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs       []Output      `xml:"extensionElements>ioMapping>output,omitempty"`
}
