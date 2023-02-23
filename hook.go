package hook

import "sync"

type Handler[T any] func(data T)

type Hook[T any] struct {
	mu       sync.RWMutex
	handlers []*Handler[T]
}

func New[T any]() *Hook[T] {

	hook := &Hook[T]{
		mu:       sync.RWMutex{},
		handlers: []*Handler[T]{},
	}

	return hook
}

func (hook *Hook[T]) Trigger(d T) {
	hook.mu.RLock()
	defer hook.mu.RUnlock()
	for _, handler := range hook.handlers {
		go func(h *Handler[T]) {
			(*h)(d)
		}(handler)
	}

}

func (hook *Hook[T]) AddHandler(handler Handler[T]) *Handler[T] {
	hook.mu.Lock()
	defer hook.mu.Unlock()
	hook.handlers = append(hook.handlers, &handler)
	return &handler
}

func (hook *Hook[T]) RemoveHandler(handler *Handler[T]) {
	hook.mu.Lock()
	defer hook.mu.Unlock()

	for i, h := range hook.handlers {
		if h == handler {
			hook.handlers = append(hook.handlers[:i], hook.handlers[i+1:]...)
			break
		}
	}
}
