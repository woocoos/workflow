package engine

import (
	"context"
	"fmt"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"sync"
)

type (
	HandlerFunc func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error)
)

var (
	handlerFuncMap = make(map[string]HandlerFunc)
	mu             sync.RWMutex
)

func init() {
	RegistryGlobalHandler("test", func(ctx context.Context, vars bpmn.Mappings) (bpmn.Mappings, error) {
		return nil, nil
	})
}

type HandlerFuncMap struct {
	id string
}

func NewHandlerFuncMap(id string) *HandlerFuncMap {
	return &HandlerFuncMap{id: id}
}

func (h *HandlerFuncMap) RegistryHandler(name string, fun HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()
	handlerFuncMap[h.id+"."+name] = fun
}

func (h *HandlerFuncMap) RunHandler(ctx context.Context, name string, vars bpmn.Mappings) (bpmn.Mappings, error) {
	mu.RLock()
	defer mu.RUnlock()
	if fun, ok := handlerFuncMap[h.id+"."+name]; ok {
		return fun(ctx, vars)
	}
	if fun, ok := handlerFuncMap[name]; ok {
		return fun(ctx, vars)
	}
	return nil, fmt.Errorf("not support service task type: %s", name)
}

func RegistryGlobalHandler(name string, fun HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()
	handlerFuncMap[name] = fun
}
