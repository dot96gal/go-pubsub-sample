package pubsub

import "context"

type Event interface {
	ID() string
	Payload() []byte
}

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type EventPublisher interface {
	Publish(ctx context.Context, event Event) error
}

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler) error
}
