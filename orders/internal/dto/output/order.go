package output

import (
	"time"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type OrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}

func NewOrdersResponse(orders []*domain.Order) *OrdersResponse {
	responses := make([]OrderResponse, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, NewOrderResponse(order))
	}
	
	return &OrdersResponse{Orders: responses}
}

type OrderResponse struct {
	Order Order `json:"order"`
}

func NewOrderResponse(o *domain.Order) OrderResponse {
	return OrderResponse{
		Order: *NewOrder(o.ID, o.UserID, o.ProductID, o.Amount, o.Status, o.CreatedAt, o.UpdatedAt),
	}
}

type Order struct {
	ID        uuid.UUID          `json:"id"`
	UserID    uuid.UUID          `json:"userId"`
	ProductID uuid.UUID          `json:"productId"`
	Amount    int64              `json:"amount"`
	Status    domain.OrderStatus `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func NewOrder(id, userId, productId uuid.UUID, amount int64, status domain.OrderStatus, created, updated time.Time) *Order {
	return &Order{
		ID:        id,
		UserID:    userId,
		ProductID: productId,
		Amount:    amount,
		Status:    status,
		CreatedAt: created,
		UpdatedAt: updated,
	}
}
