package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
)

type PayOrderUsecase interface {
	ProcessOrder(ctx context.Context, orderID uuid.UUID, payload *domain.Order) error
}
