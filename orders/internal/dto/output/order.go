package output

import (
	"time"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type OrderResponse struct {
	ID        uuid.UUID          `json:"id"`
	UserID    uuid.UUID          `json:"userId"`
	ProductID uuid.UUID          `json:"productId"`
	Amount    int64              `json:"amount"`
	Status    domain.OrderStatus `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func NewOrderResponse(id, userId, productId uuid.UUID, amount int64, status domain.OrderStatus, created, updated time.Time) *OrderResponse {
	return &OrderResponse{
		ID:        id,
		UserID:    userId,
		ProductID: productId,
		Amount:    amount,
		Status:    status,
		CreatedAt: created,
		UpdatedAt: updated,
	}
}
