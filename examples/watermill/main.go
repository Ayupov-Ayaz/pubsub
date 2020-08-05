package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

func send(broker Publisher, topic string) {
	count := 0

	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte(topic+" "+strconv.Itoa(count)))
		count++
		if err := broker.Publish(topic, msg); err != nil {
			log.Fatal(err)
		}
	}
}

func Run() error {
	broker := NewBasicPublisher()
	observer := NewLoggerObserver()
	for _, topic := range observer.Subscriptions() {
		if err := broker.Subscribe(context.Background(), topic, observer); err != nil {
			return err
		}
		go send(broker, topic)
	}

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
