package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OutboxRepository) MarkPublished(ctx context.Context, id uuid.UUID, status domain.EventStatus) error {
	const query = `UPDATE outbox SET status = $1, published_at = NOW() WHERE id = $2`
	tag, err := r.db.Exec(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrEventNotFound
	}

	return nil
}
