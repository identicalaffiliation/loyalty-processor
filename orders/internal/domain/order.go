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
	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID uuid.UUID
	Amount    int64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
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
