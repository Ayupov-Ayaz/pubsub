package main

import (
	"fmt"
	"strconv"
)

var (
	loggerID int
)

type LoggerObserver struct {
	ID  string
	got int
}

func NewLoggerObserver() *LoggerObserver {
	loggerID++
	return &LoggerObserver{ID: strconv.Itoa(loggerID)}
}

func (l *LoggerObserver) Notify(e Event) {
	l.got++
	fmt.Println("ID_" + l.ID + " " + e.Event() + ": got_messages:" + strconv.Itoa(l.got))
}
