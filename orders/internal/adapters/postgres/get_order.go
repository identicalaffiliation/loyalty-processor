package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *OrdersRepository) GetProductByOrderID(ctx context.Context, id uuid.UUID) (*uuid.UUID, error) {
	const query = `SELECT product_id FROM orders WHERE id = $1`
	var productID uuid.UUID
	if err := r.db.QueryRow(ctx, query, id).Scan(&productID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrProductNotFound
		}

		return nil, fmt.Errorf("get product id: %w", err)
	}
	
	return &productID, nil
}
