package main

import (
	"runtime"
)

func sendingMessages(publish <-chan Event, observers []Observer) {
	for e := range publish {
		for _, o := range observers {
			o.Notify(e)
		}
	}
}

func NewSubGroup(observers []Observer) chan<- Event {
	ch := make(chan Event, runtime.NumCPU()*2)

	go sendingMessages(ch, observers)

	return ch
}
