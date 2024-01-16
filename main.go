package main

import (
	"context"
	"os"
)

func main() {
	helloEvent := NewHelloEvent()
	goodbyeEvent := NewGoodbyeEvent()

	echoHandler := NewEchoHandler(os.Stdout)
	printHandler := NewPrintHandler(os.Stdout)

	dispatcher := NewEventDispatcher()
	dispatcher.Subscribe(helloEvent, echoHandler)
	dispatcher.Subscribe(helloEvent, printHandler)
	dispatcher.Subscribe(goodbyeEvent, echoHandler)

	dispatcher.Publish(context.Background(), helloEvent)
	dispatcher.Publish(context.Background(), goodbyeEvent)
}
