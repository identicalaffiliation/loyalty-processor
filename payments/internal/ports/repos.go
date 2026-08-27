package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
)

type OutboxRepository interface {
	GetUnpublishedEventsByLimit(ctx context.Context, limit int) ([]*domain.Event, error)
	MarkPublished(ctx context.Context, id uuid.UUID) error
}
