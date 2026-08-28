package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	Created   EventStatus = "created"
	Published EventStatus = "published"
)

type EventStatus string

type Event struct {
	ID          uuid.UUID
	OrderID     uuid.UUID
	Status      EventStatus
	Payload     []byte
	CreatedAt   time.Time
	PublishedAt *time.Time
}

func NewEvent(orderID uuid.UUID, status EventStatus, payload []byte) *Event {
	return &Event{
		ID:      uuid.New(),
		OrderID: orderID,
		Status:  status,
		Payload: payload,
	}
}

const (
	Success MessageStatus = "success"
	Fail    MessageStatus = "fail"
)

type MessageStatus string

type Message struct {
	Status  MessageStatus `json:"status"`
	OrderID uuid.UUID     `json:"orderId"`
	UserID  uuid.UUID     `json:"userId"`
}
