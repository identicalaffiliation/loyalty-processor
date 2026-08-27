package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusCreated OrderStatus = "created"
	StatusPaid    OrderStatus = "paid"
	StatusFailed  OrderStatus = "failed"
)

type OrderStatus string

type Order struct {
	ID        uuid.UUID   `json:"id"`
	UserID    uuid.UUID   `json:"userId"`
	ProductID uuid.UUID   `json:"productId"`
	Amount    int64       `json:"amount"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}

func NewOrder(userId, productId uuid.UUID, amount int64, status OrderStatus) *Order {
	return &Order{
		ID:        uuid.New(),
		UserID:    userId,
		ProductID: productId,
		Amount:    amount,
		Status:    status,
	}
}
