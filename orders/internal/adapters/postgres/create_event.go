package postgres

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OutboxRepository) CreateEvent(ctx context.Context, event *domain.Event) error {
	const query = `INSERT INTO outbox (id, order_id, status, payload) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, event.ID, event.OrderID, event.Status, event.Payload)
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}
	
	return nil
}
