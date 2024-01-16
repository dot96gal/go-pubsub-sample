package main

import (
	"context"
	"fmt"
	"io"

	pubsub "github.com/dot96gal/go-pubsub-sample/internal"
)

var _ pubsub.EventHandler = (*EchoHandler)(nil)

type EchoHandler struct {
	writer io.Writer
}

func NewEchoHandler(writer io.Writer) EchoHandler {
	return EchoHandler{
		writer: writer,
	}
}

func (h EchoHandler) Handle(ctx context.Context, event pubsub.Event) error {
	_, err := fmt.Fprintf(h.writer, "Echo: %s, %s\n", event.ID(), event.Payload())
	if err != nil {
		return err
	}

	return nil
}

var _ pubsub.EventHandler = (*PrintHandler)(nil)

type PrintHandler struct {
	writer io.Writer
}

func NewPrintHandler(writer io.Writer) PrintHandler {
	return PrintHandler{
		writer: writer,
	}
}

func (h PrintHandler) Handle(ctx context.Context, event pubsub.Event) error {
	_, err := fmt.Fprintf(h.writer, "Print: %s, %s\n", event.ID(), event.Payload())
	if err != nil {
		return err
	}

	return nil
}
