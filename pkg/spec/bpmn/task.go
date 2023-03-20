package bpmn

type Task struct {
	Activity
	// Camunda extensions
	Async bool `xml:"async,attr"`
}

type UserTask struct {
	Task
	Form          *FormDefinition                   `xml:"extensionElements>formDefinition"`
	Assignment    *AssignmentDefinition             `xml:"extensionElements>assignmentDefinition"`
	Inputs        []Input                           `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs       []Output                          `xml:"extensionElements>ioMapping>output,omitempty"`
	Headers       []Header                          `xml:"extensionElements>taskHeaders>header,omitempty"`
	MultiInstance *MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

type ServiceTask struct {
	Task
	TaskDefinition *TaskDefinition `xml:"extensionElements>taskDefinition,omitempty"`
	Inputs         []Input         `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs        []Output        `xml:"extensionElements>ioMapping>output,omitempty"`
	Headers        []Header        `xml:"extensionElements>taskHeaders>header,omitempty"`
}

type SendTask struct {
	Task
	TaskDefinition *TaskDefinition `xml:"extensionElements>taskDefinition,omitempty"`
	Inputs         []Input         `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs        []Output        `xml:"extensionElements>ioMapping>output,omitempty"`
	Headers        []Header        `xml:"extensionElements>taskHeaders>header,omitempty"`
}

type ManualTask struct {
	Task
}

type ReceiveTask struct {
	Task
	MessageRef   string        `xml:"messageRef,attr"`
	Outputs      []Output      `xml:"extensionElements>ioMapping>output,omitempty"`
	Subscription *Subscription `xml:"extensionElements>subscription"`
}

type ScriptTask struct {
	Task
	TaskDefinition *TaskDefinition `xml:"extensionElements>taskDefinition,omitempty"`
	Inputs         []Input         `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs        []Output        `xml:"extensionElements>ioMapping>output,omitempty"`
	Headers        []Header        `xml:"extensionElements>taskHeaders>header,omitempty"`
	Script         *Script         `xml:"extensionElements>script,omitempty"`
}

type BusinessRuleTask struct {
	Task
	TaskDefinition *TaskDefinition `xml:"extensionElements>taskDefinition,omitempty"`
	Inputs         []Input         `xml:"extensionElements>ioMapping>input,omitempty"`
	Outputs        []Output        `xml:"extensionElements>ioMapping>output,omitempty"`
	Headers        []Header        `xml:"extensionElements>taskHeaders>header,omitempty"`
	Script         *Script         `xml:"extensionElements>script"`
	CalledDecision *CalledDecision `xml:"extensionElements>calledDecision"`
}
