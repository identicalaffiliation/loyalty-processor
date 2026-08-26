package ports

import (
	"context"

	"github.com/google/uuid"
)

type InventoryRepository interface {
	ReserveStock(ctx context.Context, productID uuid.UUID) error
	ReleaseStock(ctx context.Context, productID uuid.UUID) error
}
