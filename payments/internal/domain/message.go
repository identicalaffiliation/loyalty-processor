package domain

import "github.com/google/uuid"

type Message struct {
	Status  EventStatus `json:"status"`
	OrderID uuid.UUID   `json:"orderId"`
	UserID  uuid.UUID   `json:"userId"`
}
