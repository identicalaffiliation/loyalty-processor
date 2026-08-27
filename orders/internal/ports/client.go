package ports

import (
	"context"

	"github.com/google/uuid"
)

type InventoryClient interface {
	Reserve(ctx context.Context, id uuid.UUID) error
	Release(ctx context.Context, id uuid.UUID) error
}
