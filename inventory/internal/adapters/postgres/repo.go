package postgres

import "github.com/identicalaffiliation/loyalty-processor/invetory/internal/ports"

type InventoryRepository struct {
	db ports.DBTX
}

func NewInventoryRepository(db ports.DBTX) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}
