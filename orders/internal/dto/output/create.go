package output

import "github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"

type CreateOrderResponse struct {
	Order OrderResponse `json:"order"`
}

func NewCreateOrderResponse(order *domain.Order) *CreateOrderResponse {
	return &CreateOrderResponse{
		Order: *NewOrderResponse(order.ID,
			order.UserID,
			order.ProductID,
			order.Amount,
			order.Status,
			order.CreatedAt,
			order.UpdatedAt,
		),
	}
}
