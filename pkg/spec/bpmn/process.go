package bpmn

type Process struct {
	BaseElement
	ProcessType                  ProcessType               `xml:"processType,attr"`
	IsClosed                     bool                      `xml:"isClosed,attr"`
	IsExecutable                 bool                      `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string                    `xml:"definitionalCollaborationRef,attr"`
	StartEvents                  []*StartEvent             `xml:"startEvent"`
	EndEvents                    []*EndEvent               `xml:"endEvent"`
	UserTasks                    []*UserTask               `xml:"userTask"`
	ServiceTasks                 []*ServiceTask            `xml:"serviceTask"`
	BusinessRules                []*BusinessRuleTask       `xml:"businessRuleTask"`
	SequenceFlows                []*SequenceFlow           `xml:"sequenceFlow"`
	CallActivities               []*CallActivity           `xml:"callActivity"`
	ExclusiveGateways            []*ExclusiveGateway       `xml:"exclusiveGateway"`
	InclusiveGateways            []*InclusiveGateway       `xml:"inclusiveGateway"`
	ParallelGateways             []*ParallelGateway        `xml:"parallelGateway"`
	EventBasedGateways           []*EventBasedGateway      `xml:"eventBasedGateway"`
	IntermediateCatchEvents      []*IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	// camunda:properties
	VersionTag string `xml:"versionTag,attr"`
}

type ProcessType string

const (
	ProcessTypeNode    ProcessType = "None"
	ProcessTypePublic  ProcessType = "Public"
	ProcessTypePrivate ProcessType = "Private"
)

func (pr Process) FindBaseElementsById(id string) (elements []Elementor) {
	appender := func(element Elementor) {
		if element.GetId() == id {
			elements = append(elements, element)
		}
	}
	for _, startEvent := range pr.StartEvents {
		appender(startEvent)
	}
	for _, endEvent := range pr.EndEvents {
		appender(endEvent)
	}
	for _, task := range pr.ServiceTasks {
		appender(task)
	}
	for _, task := range pr.UserTasks {
		appender(task)
	}
	for _, task := range pr.BusinessRules {
		appender(task)
	}
	for _, parallelGateway := range pr.ParallelGateways {
		appender(parallelGateway)
	}
	for _, exclusiveGateway := range pr.ExclusiveGateways {
		appender(exclusiveGateway)
	}
	for _, intermediateCatchEvent := range pr.IntermediateCatchEvents {
		appender(intermediateCatchEvent)
	}
	for _, eventBasedGateway := range pr.EventBasedGateways {
		appender(eventBasedGateway)
	}
	for _, activity := range pr.CallActivities {
		appender(activity)
	}

	return elements
}
