package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OrdersRepository) MarkOrder(ctx context.Context, id uuid.UUID, status domain.OrderStatus) error {
	const query = `UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`
	tag, err := r.db.Exec(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("update order: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrInvalidData
	}
	
	return nil
}
