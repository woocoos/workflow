package dmn

type Definitions struct {
	NamedElement
	Decisions []Decision `xml:"decision"`
}
