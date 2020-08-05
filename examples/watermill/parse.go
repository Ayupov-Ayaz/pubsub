package main

import (
	"strconv"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

func uint64ToStr(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

const (
	time_      = "time"
	betID      = "bet_id"
	stepID     = "step_id"
	bet        = "bet"
	win        = "win"
	balance    = "balance"
	stepsCount = "steps_count"
	stepType   = "step_type"
	seed       = "seed"
	roomID     = "room_id"
)

func NewMessage(h History) *message.Message {
	msg := message.NewMessage(watermill.NewUUID(), []byte(h.StepPayload))
	meta := msg.Metadata

	meta.Set(time_, h.Time)
	meta.Set(bet, uint64ToStr(h.Bet))
	meta.Set(betID, h.BetID.String())
	meta.Set(win, uint64ToStr(h.Win))
	meta.Set(stepID, h.StepID.String())
	meta.Set(balance, int64ToStr(h.Balance))
	meta.Set(stepsCount, strconv.Itoa(int(h.StepsCount)))
	meta.Set(stepType, strconv.Itoa(int(h.StepType)))
	meta.Set(seed, int64ToStr(h.Seed))
	meta.Set(roomID, h.RoomID)

	return msg
}
