package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrdersByUser(ctx context.Context, id uuid.UUID) ([]*domain.Order, error)
	MarkOrder(ctx context.Context, id uuid.UUID, status domain.OrderStatus) error
	GetProductByOrderID(ctx context.Context, id uuid.UUID) (*uuid.UUID, error)
}

type OutboxRepository interface {
	CreateEvent(ctx context.Context, event *domain.Event) error
	MarkPublished(ctx context.Context, id uuid.UUID, status domain.EventStatus) error
	GetUnpublishedEventsByLimit(ctx context.Context, limit int) ([]*domain.Event, error)
}
