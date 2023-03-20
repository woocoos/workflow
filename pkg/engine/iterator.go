package engine

import (
	"fmt"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"sync"
)

type Iterator struct {
	loader  *BpmnLoader
	process *bpmn.Process
	current bpmn.Elementor
	pos     int
	err     error

	queue []bpmn.Elementor
	mu    sync.RWMutex

	Request *InstanceRequest
	// 针对当前节点的处理函数
	CurrentHandler func(it *Iterator, ele bpmn.Elementor) (err error)
	// 针对当前队列中的全部元素的处理函数,一般用于消息型网关
	QueueHandler func(it *Iterator, subs []bpmn.Elementor) (next bpmn.Elementor, err error)
}

func (it *Iterator) Loader() *BpmnLoader {
	return it.loader
}

func (it *Iterator) Err() error {
	it.mu.RLock()
	defer it.mu.RUnlock()
	return it.err
}

// Next 获取下一个节点.
//
// 消息型网关需要等待外部重置当前队列.
func (it *Iterator) Next() (exist bool) {
	it.mu.Lock()
	defer it.mu.Unlock()

	if it.err != nil {
		return false
	}

	waitChoice := false
	for len(it.queue) > 0 {
		it.current = it.queue[0]
		it.queue = it.queue[1:]

		if !waitChoice && it.CurrentHandler != nil {
			it.err = it.CurrentHandler(it, it.current)
			if it.err != nil {
				return false
			}
		}
		waitChoice = false
		nextFlows := findSequenceFlows(it.process.SequenceFlows, it.current.GetOutgoing())
		if len(nextFlows) == 0 {
			return false
		}
		switch it.current.(type) {
		case *bpmn.ExclusiveGateway:
			nextFlows, it.err = flowFilter(nextFlows, it.Request)
			if it.err != nil {
				return false
			}
		case *bpmn.EventBasedGateway:
			waitChoice = true
		}
		for _, flow := range nextFlows {
			els := it.process.FindBaseElementsById(flow.TargetRef)
			if len(els) == 0 {
				it.err = fmt.Errorf("target element not found: %s", flow.TargetRef)
				return false
			}
			it.queue = append(it.queue, els[0])
		}
		if waitChoice {
			next, err := it.QueueHandler(it, it.queue)
			if err != nil {
				it.err = err
				return false
			}
			it.queue = it.queue[:0]
			it.queue = append(it.queue, next)
		}
	}
	return false
}
