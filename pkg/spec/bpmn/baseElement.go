package bpmn

import (
	"encoding/xml"
	"github.com/tsingsun/woocoo/pkg/conf"
	"html"
)

type ExtensionProperties []ExtensionProperty

func (ep ExtensionProperties) Decode(valuePtr any) error {
	mp := make(map[string]any)
	for _, p := range ep {
		mp[p.Name] = p.Value
	}
	return conf.NewFromStringMap(mp).Unmarshal(valuePtr)
}

type BaseElement struct {
	Id                  string              `xml:"id,attr"`
	Documentation       []Documentation     `xml:"documentation,omitempty"`
	ExtensionProperties ExtensionProperties `xml:"extensionElements>properties>property,omitempty"`
}

func (be *BaseElement) GetDocumentation() string {
	if len(be.Documentation) == 0 {
		return ""
	}
	return be.Documentation[0].Content
}

type BaseElementWithMixedContent struct {
	Id            string         `xml:"id,attr,omitempty"`
	Content       string         `xml:",innerxml"`
	Documentation *Documentation `xml:"documentation,omitempty"`
}

func (b *BaseElementWithMixedContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type baseElementWithMixedContent BaseElementWithMixedContent
	var bb baseElementWithMixedContent
	if err := d.DecodeElement(&bb, &start); err != nil {
		return err
	}
	*b = BaseElementWithMixedContent(bb)
	// parse html char to string
	b.Content = html.UnescapeString(bb.Content)
	return nil
}
