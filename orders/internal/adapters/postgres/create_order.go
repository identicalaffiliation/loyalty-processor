package postgres

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func (r *OrdersRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	const query = `INSERT INTO orders (id, user_id, product_id, bonus_amount, status)
		VALUES ($1, $2, $3, $4, $5) RETURNING *`

	var created domain.Order
	err := r.db.QueryRow(
		ctx,
		query,
		order.ID,
		order.UserID,
		order.ProductID,
		order.Amount,
		order.Status,
	).Scan(
		&created.ID,
		&created.UserID,
		&created.ProductID,
		&created.Amount,
		&created.Status,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		if checkConstraits(err) {
			return nil, domain.ErrInvalidData
		}

		return nil, fmt.Errorf("create order: %w", err)
	}

	return &created, nil
}
