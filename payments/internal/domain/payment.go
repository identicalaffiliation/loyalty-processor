package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"orderId"`
	UserID    uuid.UUID `json:"userId"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPayment(orderId, userId uuid.UUID, amount int64) *Payment {
	return &Payment{
		ID:      uuid.New(),
		OrderID: orderId,
		UserID:  userId,
		Amount:  amount,
	}
}
