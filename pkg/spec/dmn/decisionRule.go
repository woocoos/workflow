package dmn

type DecisionRule struct {
	DMNElement
	InputEntries  []UnaryTests        `xml:"inputEntry"`
	OutputEntries []LiteralExpression `xml:"outputEntry"`
}
