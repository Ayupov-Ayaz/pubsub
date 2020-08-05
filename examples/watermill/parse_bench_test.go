package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func BenchmarkNewMessage(b *testing.B) {

	history := History{
		Time:        time.Now().Format(time.RFC3339),
		Bet:         15,
		Win:         100,
		Balance:     232323,
		RoomID:      "room_id",
		BetID:       uuid.New(),
		StepsCount:  2,
		StepID:      uuid.New(),
		StepType:    1,
		StepPayload: []byte(`{"linesCount":10}`),
		Seed:        131313131,
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		msg := NewMessage(history)
		if len(msg.Payload) == 0 {
			b.Fatal("len == 0")
		}
	}
}
