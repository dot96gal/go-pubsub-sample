package main

import (
	"context"
	"sync"

	pubsub "github.com/dot96gal/go-pubsub-sample/internal"
)

var _ interface {
	pubsub.EventPublisher
	pubsub.EventSubscriber
} = (*EventDispatcher)(nil)

type EventDispatcher struct {
	mu       sync.Mutex
	handlers map[string][]pubsub.EventHandler
}

func NewEventDispatcher() EventDispatcher {
	return EventDispatcher{
		handlers: make(map[string][]pubsub.EventHandler),
	}
}

func (h *EventDispatcher) Subscribe(event pubsub.Event, handler pubsub.EventHandler) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[event.ID()] = append(h.handlers[event.ID()], handler)

	return nil
}

func (h *EventDispatcher) Publish(ctx context.Context, event pubsub.Event) error {
	for _, handler := range h.handlers[event.ID()] {
		err := handler.Handle(ctx, event)
		if err != nil {
			return err
		}
	}

	return nil
}
