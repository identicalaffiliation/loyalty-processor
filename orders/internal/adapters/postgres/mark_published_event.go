package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OutboxRepository) MarkPublished(ctx context.Context, eventID uuid.UUID) error {
	const query = `UPDATE outbox SET status = 'published', published_at = NOW() WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, eventID)
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrEventNotFound
	}
	
	return nil
}
