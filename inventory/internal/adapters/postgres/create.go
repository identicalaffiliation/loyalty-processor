package postgres

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
)

func (r *InventoryRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	const query = `INSERT INTO products (id, title, description, stock) VALUES ($1, $2, $3, $4)`

	if _, err := r.db.Exec(ctx, query, product.ID, product.Title, product.Description, product.Stock); err != nil {
		return fmt.Errorf("create product: %w", err)
	}

	return nil
}
