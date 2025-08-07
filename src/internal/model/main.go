package model

import (
	"time"

	"github.com/google/uuid"
)

type QueMessage struct {
	Id  uuid.UUID `json:"id"`
	Dt  time.Time `json:"dt"`
	CAt uint8     `json:"ca_t"`
	Ack bool      `json:"ack"`
	Msg []byte
}
