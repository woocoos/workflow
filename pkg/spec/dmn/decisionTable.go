package dmn

type HitPolicy string

const (
	HitPolicyUNIQUE      HitPolicy = "UNIQUE"
	HitPolicyFIRST       HitPolicy = "FIRST"
	HitPolicyPRIORITY    HitPolicy = "PRIORITY"
	HitPolicyANY         HitPolicy = "ANY"
	HitPolicyCOLLECT     HitPolicy = "COLLECT"
	HitPolicyRULEORDER   HitPolicy = "RULE ORDER"
	HitPolicyOUTPUTORDER HitPolicy = "OUTPUT ORDER"
)

type BuiltinAggregator string

const (
	BuiltinAggregatorSUM   BuiltinAggregator = "SUM"
	BuiltinAggregatorMIN   BuiltinAggregator = "MIN"
	BuiltinAggregatorMAX   BuiltinAggregator = "MAX"
	BuiltinAggregatorCOUNT BuiltinAggregator = "COUNT"
)

type PreferredOrientation string

const (
	PreferredOrientationRuleAsRow    = "Rule-as-Row"
	PreferredOrientationRuleAsColumn = "Rule-as-Column"
	PreferredOrientationCrossTable   = "CrossTable"
)

type TypeRef string

var (
	TypeRefString            TypeRef = "string"
	TypeRefNumber            TypeRef = "number"
	TypeRefBoolean           TypeRef = "boolean"
	TypeRefDate              TypeRef = "date"
	TypeRefTime              TypeRef = "time"
	TypeRefDateTime          TypeRef = "dateTime"
	TypeRefDayTimeDuration   TypeRef = "dayTimeDuration"
	TypeRefYearMonthDuration TypeRef = "yearMonthDuration"
	TypeRefAny               TypeRef = "any"
)

type DecisionTable struct {
	Expression
	Inputs      []InputClause      `xml:"input"`
	Outputs     []OutputClause     `xml:"output"`
	Rules       []DecisionRule     `xml:"rule"`
	HitPolicy   HitPolicy          `xml:"hitPolicy,attr,omitempty"`
	Aggregation *BuiltinAggregator `xml:"aggregation,attr,omitempty"`
	//PreferredOrientation *PreferredOrientation `xml:"preferredOrientation,attr,omitempty"`
	OutputLabel string `xml:"outputLabel,attr,omitempty"`
}

type Expression struct {
	DMNElement
	TypeRef TypeRef `xml:"typeRef,attr"`
}
