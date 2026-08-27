package output

import "github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

func NewCreateOrderResponse(order *domain.Order) *CreateOrderResponse {
	return &CreateOrderResponse{
		Order: *NewOrder(order.ID,
			order.UserID,
			order.ProductID,
			order.Amount,
			order.Status,
			order.CreatedAt,
			order.UpdatedAt,
		),
	}
}
