package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	betTop = "bet"
	winTop = "win"
)

type LoggerObserver struct{}

func NewLoggerObserver() LoggerObserver {
	return LoggerObserver{}
}

func (o LoggerObserver) Subscriptions() []string {
	return []string{
		betTop,
		winTop,
	}
}

func (o LoggerObserver) Notify(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
