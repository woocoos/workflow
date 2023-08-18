package engine

import (
	"context"
	"fmt"
	"github.com/woocoos/workflow/pkg/spec/vars"
	"sync"
)

type (
	HandlerFunc func(ctx context.Context, input vars.Mapping) (vars.Mapping, error)
)

var (
	handlerFuncMap = make(map[string]HandlerFunc)
	mu             sync.RWMutex
)

func init() {
	RegistryGlobalHandler("test", func(ctx context.Context, input vars.Mapping) (vars.Mapping, error) {
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

func (h *HandlerFuncMap) RunHandler(ctx context.Context, name string, input vars.Mapping) (vars.Mapping, error) {
	mu.RLock()
	defer mu.RUnlock()
	if fun, ok := handlerFuncMap[h.id+"."+name]; ok {
		return fun(ctx, input)
	}
	if fun, ok := handlerFuncMap[name]; ok {
		return fun(ctx, input)
	}
	return nil, fmt.Errorf("not support service task type: %s", name)
}

func RegistryGlobalHandler(name string, fun HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()
	handlerFuncMap[name] = fun
}
