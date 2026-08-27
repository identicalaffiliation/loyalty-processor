package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/ports"
)

type BalancesRepository struct {
	db ports.DBTX
}

func NewBalancesRepository(db ports.DBTX) *BalancesRepository {
	return &BalancesRepository{db: db}
}

func (r *BalancesRepository) CreateBalance(ctx context.Context, balance *domain.Balance) error {
	const query = `INSERT INTO balances (user_id, bonuses) VALUES ($1, $2) 
    ON CONFLICT (user_id) DO NOTHING`

	if _, err := r.db.Exec(ctx, query, balance.UserID, balance.Bonuses); err != nil {
		return fmt.Errorf("create balance: %w", err)
	}

	return nil
}

func (r *BalancesRepository) PayOrder(ctx context.Context, userId uuid.UUID, amount int64) error {
	const query = `UPDATE balances SET bonuses = bonuses - $1, updated_at = NOW() WHERE user_id = $2 AND bonuses >= $1`
	tag, err := r.db.Exec(ctx, query, amount, userId)
	if err != nil {
		if checkConstraits(err) {
			return domain.ErrInvalidBalance
		}

		return fmt.Errorf("update balance: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrInvalidBalance
	}

	return nil
}
