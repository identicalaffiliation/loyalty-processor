package postgres

import "github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"

type OrdersRepository struct {
	db ports.DBTX
}

func NewOrdersRepository(db ports.DBTX) *OrdersRepository {
	return &OrdersRepository{db: db}
}
