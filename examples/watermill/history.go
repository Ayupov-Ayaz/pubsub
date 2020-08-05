package main

import (
	"encoding/json"

	"github.com/google/uuid"
)

type History struct {
	Time        string          `json:"time" validate:"required"`
	Bet         uint64          `json:"bet"`
	Win         uint64          `json:"win"`
	Balance     int64           `json:"balance"`
	RoomID      string          `json:"roomId" validate:"required"`
	BetID       uuid.UUID       `json:"betId" validate:"required"`
	StepsCount  uint32          `json:"stepsCount"`
	StepID      uuid.UUID       `json:"stepId" validate:"required"`
	StepType    uint32          `json:"stepType" validate:"required"`
	StepPayload json.RawMessage `json:"stepPayload" validate:"required"`
	Seed        int64           `json:"seed" validate:"required"`
	//todo: gift
}
