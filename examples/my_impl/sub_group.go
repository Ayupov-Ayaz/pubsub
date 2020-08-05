package main

import (
	"runtime"
)

type SubGroup struct {
	publish   chan Event
	observers []Observer
}

func (g *SubGroup) sendingMessages() {
	for e := range g.publish {
		for _, o := range g.observers {
			o.Notify(e)
		}
	}
}

func NewSubGroup(observers []Observer) chan<- Event {
	// нет возможности добавлять в режиме runtime дополнительных слушателей
	// сделано для того, чтобы исключить использование mutex во избежание тормозов
	ch := make(chan Event, runtime.NumCPU()*2)

	g := &SubGroup{
		publish:   ch,
		observers: observers,
	}

	go g.sendingMessages()

	return ch
}
