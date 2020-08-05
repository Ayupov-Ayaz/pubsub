package main

import "errors"

type Event interface {
	Event() string
}

type Observer interface {
	Notify(e Event)
}

type BasicBroker struct {
	subGroups map[string]chan<- Event
}

func NewBasicBroker() *BasicBroker {
	return &BasicBroker{
		subGroups: make(map[string]chan<- Event),
	}
}

func (b *BasicBroker) SubscribeToEvent(name string, observers []Observer) error {
	// нет возможности добавлять в режиме runtime дополнительных слушателей
	// сделано для того, чтобы исключить использование mutex во избежание тормозов
	if _, ok := b.subGroups[name]; ok {
		return errors.New("group exist")
	}

	b.subGroups[name] = NewSubGroup(observers)
	return nil
}

func (b BasicBroker) Send(e Event) {
	group, ok := b.subGroups[e.Event()]
	if ok {
		group <- e
	}
}
