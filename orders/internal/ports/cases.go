package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/input"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/output"
)

type CreateOrderUsecase interface {
	CreateOrder(ctx context.Context, req *input.CreateOrderRequest) (*output.CreateOrderResponse, error)
}

type GetOrdersUsecase interface {
	GetOrdersByUser(ctx context.Context, id uuid.UUID) (*output.OrdersResponse, error)
}
