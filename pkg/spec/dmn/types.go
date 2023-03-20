package dmn

type UnaryTests struct {
	DMNElement
	Text               string `xml:"text"`
	ExpressionLanguage string `xml:"expressionLanguage,attr,omitempty"`
}

type InputClause struct {
	DMNElement
	InputExpression *LiteralExpression `xml:"inputExpression"`
	InputValues     []UnaryTests       `xml:"inputValues"`
}

type OutputClause struct {
	DMNElement
	OutputValues []UnaryTests `xml:"outputValues"`
	Name         string       `xml:"name,attr"`
	TypeRef      TypeRef      `xml:"typeRef,attr"`
}

type LiteralExpression struct {
	Expression
	// choice of text or importedValues
	Text               string          `xml:"text"`
	ImportedValues     *ImportedValues `xml:"importedValues"`
	ExpressionLanguage string          `xml:"expressionLanguage,attr,omitempty"`
}

type ImportedValues struct {
	Import
	ImportedElement    string `xml:"importedElement"`
	ExpressionLanguage string `xml:"expressionLanguage,attr"`
}

type Import struct {
	Namespace   string `xml:"namespace,attr"`
	LocationURI string `xml:"locationURI,attr,omitempty"`
	ImportType  string `xml:"importType,attr"`
}
