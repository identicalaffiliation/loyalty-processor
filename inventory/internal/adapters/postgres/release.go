package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
)

func (r *InventoryRepository) ReleaseStock(ctx context.Context, productID uuid.UUID) error {
	const query = `UPDATE products SET stock = stock + 1 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, productID)
	if err != nil {
		return fmt.Errorf("update stock: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}
