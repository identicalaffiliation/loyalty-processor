package ports

import (
	"context"

	"github.com/google/uuid"
)

type InventoryApplication interface {
	ReserveStock(ctx context.Context, id uuid.UUID) error
	ReleaseStock(ctx context.Context, id uuid.UUID) error
}
