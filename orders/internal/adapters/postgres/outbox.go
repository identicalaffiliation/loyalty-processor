package postgres

import "github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"

type OutboxRepository struct {
	db ports.DBTX
}

func NewOutboxRepository(db ports.DBTX) *OutboxRepository {
	return &OutboxRepository{db: db}
}
