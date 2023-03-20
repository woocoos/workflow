package bpmn

type Message struct {
	RootElement
	Name         string        `xml:"name,attr"`
	Subscription *Subscription `xml:"extensionElements>subscription"`
}

func (m Message) EvaluationSubscription(vars map[string]any) (string, error) {
	if m.Subscription == nil {
		return "", nil
	}
	return EvaluateExpressionToString(m.Subscription.CorrelationKey, vars)
}

type Signal struct {
	RootElement
	Name string `xml:"name,attr"`
}
