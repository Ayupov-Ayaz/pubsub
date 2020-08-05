package main

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type Observer interface {
	Notify(<-chan *message.Message)
}

type Publisher interface {
	Subscribe(ctx context.Context, topic string, observer Observer) error
	Publish(topic string, messages ...*message.Message) error
}

type BasicPublisher struct {
	pubSub *gochannel.GoChannel
}

func NewBasicPublisher() Publisher {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	b := BasicPublisher{pubSub: pubSub}

	return b
}

func (b BasicPublisher) Subscribe(ctx context.Context, topic string, observer Observer) error {
	ch, err := b.pubSub.Subscribe(ctx, topic)
	if err != nil {
		return fmt.Errorf("subscribe failed: %w", err)
	}

	go observer.Notify(ch)

	return nil
}

func (b BasicPublisher) Publish(topic string, messages ...*message.Message) error {
	if err := b.pubSub.Publish(topic, messages...); err != nil {
		return fmt.Errorf("publish failed: %w", err)
	}

	return nil
}
