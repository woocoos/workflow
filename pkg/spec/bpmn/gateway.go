package bpmn

type GatewayDirection string

const (
	GatewayDirectionUnspecified GatewayDirection = "Unspecified"
	GatewayDirectionConverging  GatewayDirection = "Converging"
	GatewayDirectionDiverging   GatewayDirection = "Diverging"
	GatewayDirectionMixed       GatewayDirection = "Mixed"
)

type EventBasedGatewayType string

const (
	EventBasedGatewayTypeExclusive EventBasedGatewayType = "Exclusive"
	EventBasedGatewayTypeParallel  EventBasedGatewayType = "Parallel"
)

type Gateway struct {
	FlowNode
	GatewayDirection GatewayDirection `xml:"gatewayDirection,attr"`
}

type EventBasedGateway struct {
	Gateway
	Instantiate           bool                  `xml:"instantiate,attr"`
	EventBasedGatewayType EventBasedGatewayType `xml:"eventGatewayType,attr"`
}

type ParallelGateway struct {
	Gateway
	// Camunda extensions
	Async     bool `xml:"async,attr"`
	Exclusive bool `xml:"exclusive,attr"`
}

type InclusiveGateway struct {
	Gateway
	Default string `xml:"default,attr,omitempty"`
}

type ComplexGateway struct {
	Gateway
	ActivationCondition *Expression `xml:"activationCondition,attr,omitempty"`
	Default             string      `xml:"default,attr"`
}

type ExclusiveGateway struct {
	Gateway
}

func (e ExclusiveGateway) GetId() string {
	return e.Id
}

func (e ExclusiveGateway) GetOutgoing() []string {
	return e.Outgoing
}
