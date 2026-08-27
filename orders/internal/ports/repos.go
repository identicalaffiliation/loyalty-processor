package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}

type OutboxRepository interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
	MarkPublished(ctx context.Context, id uuid.UUID) error
	GetUnpublishedEventsByLimit(ctx context.Context, limit int) ([]*domain.Event, error)
}
