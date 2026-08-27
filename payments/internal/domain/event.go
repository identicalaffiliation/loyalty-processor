package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	OrderID     uuid.UUID
	Payload     []byte
	CreatedAt   time.Time
	PublishedAt *time.Time
}

func NewEvent(orderId uuid.UUID, data []byte) *Event {
	return &Event{
		ID:      uuid.New(),
		OrderID: orderId,
		Payload: data,
	}
}
