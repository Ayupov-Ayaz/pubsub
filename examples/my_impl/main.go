package main

import (
	"log"
	"time"
)

const (
	betObserversCount = 2
	winObserversCount = 3
)

func betEvents(broker *BasicBroker) {
	for {
		broker.Send(&BetEvent{})
	}
}

func winEvents(broker *BasicBroker) {
	for {
		broker.Send(&WinEvent{})
	}
}

func Run() error {
	broker := NewBasicBroker()

	var (
		betObservers []Observer
		winObservers []Observer
	)

	for i := 0; i < betObserversCount; i++ {
		betObservers = append(betObservers, NewLoggerObserver())
	}

	for i := 0; i < winObserversCount; i++ {
		winObservers = append(winObservers, NewLoggerObserver())
	}

	if err := broker.SubscribeToEvent(betEvent, betObservers); err != nil {
		return err
	}

	if err := broker.SubscribeToEvent(winEvent, winObservers); err != nil {
		return err
	}

	go winEvents(broker)
	go betEvents(broker)

	time.Sleep(2 * time.Millisecond) //

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
