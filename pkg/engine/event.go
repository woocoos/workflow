package engine

import (
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"go.temporal.io/sdk/workflow"
)

type EventListen struct {
	loader     *BpmnLoader
	eles       []bpmn.Elementor
	isComplete bool
	next       bpmn.Elementor
	err        error
	input      map[string]any
	output     map[string]any
}

func (elsn *EventListen) MultiEventListen(ctx workflow.Context) {
	ao := DefaultActivityOption()
	ctx = workflow.WithActivityOptions(ctx, ao)

	childCtx, cancelHandler := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)
	for _, ele := range elsn.eles {
		switch el := ele.(type) {
		case *bpmn.IntermediateCatchEvent:
			switch {
			case el.MessageEventDefinition != nil:
				msg := elsn.loader.FindMessage(el.MessageEventDefinition.MessageRef)
				key, err := msg.EvaluationSubscription(elsn.input)
				if err != nil {
					elsn.err = err
					return
				}
				ch := workflow.GetSignalChannel(childCtx, key)
				selector.AddReceive(ch, func(c workflow.ReceiveChannel, more bool) {
					var signal map[string]any
					c.Receive(ctx, &signal)
					if elsn.DoneIntermediateCatchEvent(el, signal, cancelHandler) != nil {
						return
					}
				})
			case el.SignalEventDefinition != nil:
				sig := elsn.loader.FindSignal(el.SignalEventDefinition.SignalRef)
				ch := workflow.GetSignalChannel(childCtx, sig.Name)
				selector.AddReceive(ch, func(c workflow.ReceiveChannel, more bool) {
					if elsn.DoneIntermediateCatchEvent(el, nil, cancelHandler) != nil {
						return
					}
				})
			case el.TimerEventDefinition != nil:
				du, err := bpmn.EvaluateExpressionToDuration(el.TimerEventDefinition.TimeDuration.Content)
				if err != nil {
					elsn.err = err
					return
				}
				timeFuture := workflow.NewTimer(childCtx, du)
				selector.AddFuture(timeFuture, func(f workflow.Future) {
					if elsn.DoneIntermediateCatchEvent(el, nil, cancelHandler) != nil {
						return
					}
				})
			}
		}
	}
	selector.Select(ctx)
}

func (elsn *EventListen) DoneIntermediateCatchEvent(ele *bpmn.IntermediateCatchEvent, vars map[string]any, cancel workflow.CancelFunc) error {
	elsn.isComplete = true
	cancel()
	elsn.next = ele
	for _, output := range ele.Outputs {
		elsn.err = output.OutputTarget(vars, elsn.output)
		if elsn.err != nil {
			break
		}
	}
	return elsn.err
}
