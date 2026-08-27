package postgres

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OutboxRepository) GetUnpublishedEventsByLimit(ctx context.Context, limit int) ([]*domain.Event, error) {
	const query = `SELECT id, order_id, status, payload, created_at, published_at
		FROM outbox WHERE status = 'created' AND published_at IS NULL LIMIT $1`

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
			&event.Status,
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
