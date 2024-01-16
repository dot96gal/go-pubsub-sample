package main

import (
	pubsub "github.com/dot96gal/go-pubsub-sample/internal"
	"github.com/google/uuid"
)

var _ pubsub.Event = (*HelloEvent)(nil)

type HelloEvent struct {
	id      uuid.UUID
	message string
}

func NewHelloEvent() HelloEvent {
	return HelloEvent{
		id:      uuid.New(),
		message: "hello",
	}
}

func (e HelloEvent) ID() string {
	return e.id.String()
}

func (e HelloEvent) Payload() []byte {
	return []byte(e.message)
}

var _ pubsub.Event = (*GoodbyeEvent)(nil)

type GoodbyeEvent struct {
	id      uuid.UUID
	message string
}

func NewGoodbyeEvent() GoodbyeEvent {
	return GoodbyeEvent{
		id:      uuid.New(),
		message: "goodbye",
	}
}

func (e GoodbyeEvent) ID() string {
	return e.id.String()
}

func (e GoodbyeEvent) Payload() []byte {
	return []byte(e.message)
}
