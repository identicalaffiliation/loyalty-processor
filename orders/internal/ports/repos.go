package ports

import (
	"context"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
