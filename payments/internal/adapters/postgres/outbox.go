package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/ports"
)

type OutboxRepository struct {
	db ports.DBTX
}

func NewOutboxRepository(db ports.DBTX) *OutboxRepository {
	return &OutboxRepository{db: db}
}

func (r *OutboxRepository) CreateEvent(ctx context.Context, event *domain.Event) error {
	const query = `INSERT INTO outbox (id, order_id, payload) VALUES ($1, $2, $3)`
	if _, err := r.db.Exec(ctx, query, event.ID, event.OrderID, event.Payload); err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	return nil
}

func (r *OutboxRepository) GetUnpublishedEventsByLimit(ctx context.Context, limit int) ([]*domain.Event, error) {
	const query = `SELECT id, order_id, payload, created_at, published_at
		FROM outbox WHERE published_at IS NULL LIMIT $1`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("get rows: %w", err)
	}

	defer rows.Close()

	events := make([]*domain.Event, 0, limit)
	for rows.Next() {
		var event domain.Event
		err = rows.Scan(
			&event.ID,
			&event.OrderID,
			&event.Payload,
			&event.CreatedAt,
			&event.PublishedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}

	return events, nil
}

func (r *OutboxRepository) MarkPublished(ctx context.Context, eventID uuid.UUID) error {
	const query = `UPDATE outbox SET published_at = NOW() WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, eventID)
	if err != nil {
		return fmt.Errorf("update event: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrEventNotFound
	}

	return nil
}
