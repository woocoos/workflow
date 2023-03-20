package bpmn

type Elementor interface {
	GetId() string
	GetOutgoing() []string
}

type Definitions struct {
	BaseElement
	TargetNamespace    string    `xml:"targetNamespace,attr"`
	ExpressionLanguage string    `xml:"expressionLanguage,attr"`
	TypeLanguage       string    `xml:"typeLanguage,attr"`
	Exporter           string    `xml:"exporter,attr"`
	ExporterVersion    string    `xml:"exporterVersion,attr"`
	Process            []Process `json:"process,omitempty" xml:"process"`
	Messages           []Message `json:"message,omitempty" xml:"message"`
	Signals            []Signal  `json:"signal,omitempty" xml:"signal"`
}
