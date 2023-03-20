package bpmn

type Documentation struct {
	Content    string `xml:",innerxml"`
	Id         string `xml:"id,attr"`
	TextFormat string `xml:"textFormat,attr"`
}
