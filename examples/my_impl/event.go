package main

const (
	betEvent = "bet"
	winEvent = "win"
)

type BetEvent struct {
}

func (b BetEvent) Event() string {
	return betEvent
}

type WinEvent struct {
}

func (w WinEvent) Event() string {
	return winEvent
}
