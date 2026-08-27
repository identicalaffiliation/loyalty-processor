package domain

import (
	"time"

	"github.com/google/uuid"
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
