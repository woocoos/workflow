package dmn

// Decision is a decision element in a DMN model.
//
// TODO not all fields are implemented yet.
type Decision struct {
	DRGElement
	Variable               *InformationItem         `xml:"variable"`
	InformationRequirement []InformationRequirement `xml:"informationRequirement"`
	DecisionTable          *DecisionTable           `xml:"decisionTable"`
	LiteralExpression      *LiteralExpression       `xml:"literalExpression"`
}

type DRGElement struct {
	DMNElement
	Name string `xml:"name,attr"`
}

type DMNElement struct {
	Id          string `xml:"id,attr,omitempty"`
	Description string `xml:"description"`
	Label       string `xml:"label,attr,omitempty"`
}

type InformationItem struct {
	NamedElement
	TypeRef string `xml:"typeRef,attr"`
}

type InformationRequirement struct {
	DMNElement
	RequiredInput    *DMNElementReference `xml:"requiredInput"`
	RequiredDecision *DMNElementReference `xml:"requiredDecision"`
}

type DMNElementReference struct {
	Href string `xml:"href,attr"`
}

type NamedElement struct {
	DMNElement
	Name string `xml:"name,attr"`
}
