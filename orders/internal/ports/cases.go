package ports

import (
	"context"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/input"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/output"
)

type CreateOrderUsecase interface {
	CreateOrder(ctx context.Context, req *input.CreateOrderRequest) (*output.CreateOrderResponse, error)
}
