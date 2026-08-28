package postgres

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/ports"
)

type PaymentsRepository struct {
	db ports.DBTX
}

func NewPaymentsRepository(db ports.DBTX) *PaymentsRepository {
	return &PaymentsRepository{db: db}
}

func (r *PaymentsRepository) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	const query = `INSERT INTO payments (id, order_id, user_id, bonuses_amount) 
		VALUES ($1, $2, $3, $4) RETURNING *`

	var created domain.Payment
	err := r.db.QueryRow(ctx, query, payment.ID, payment.OrderID, payment.UserID, payment.Amount).Scan(
		&created.ID,
		&created.OrderID,
		&created.UserID,
		&created.Amount,
		&created.CreatedAt,
	)
	if err != nil {
		if checkUniqueViolation(err) {
			return nil, domain.ErrPaymentAlreadyExists
		}
		
		if checkConstraits(err) {
			return nil, domain.ErrInvalidData
		}

		return nil, fmt.Errorf("create payment: %w", err)
	}

	return &created, nil
}
