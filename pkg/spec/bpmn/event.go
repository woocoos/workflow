package bpmn

type Event struct {
	FlowNode
}

type EventDefinition struct {
	RootElement
}

type StartEvent struct {
	CatchEvent
	MessageEventDefinition *MessageEventDefinition `xml:"messageEventDefinition,omitempty"`
	IsInterrupting         bool                    `xml:"isInterrupting,attr"`
	Outputs                []Output                `xml:"extensionElements>ioMapping>output,omitempty"`
}

type EndEvent struct {
	ThrowEvent
}

type ThrowEvent struct {
	Event
}

type CatchEvent struct {
	Event
}

type IntermediateCatchEvent struct {
	CatchEvent
	TimerEventDefinition   *TimerEventDefinition   `xml:"timerEventDefinition,omitempty"`
	MessageEventDefinition *MessageEventDefinition `xml:"messageEventDefinition,omitempty"`
	SignalEventDefinition  *SignalEventDefinition  `xml:"signalEventDefinition,omitempty"`
	Outputs                []Output                `xml:"extensionElements>ioMapping>output,omitempty"`
}

type MessageEventDefinition struct {
	EventDefinition
	MessageRef string `xml:"messageRef,attr"`
}

type SignalEventDefinition struct {
	EventDefinition
	SignalRef string `xml:"signalRef,attr"`
}
