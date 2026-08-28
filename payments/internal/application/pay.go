package application

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/ports"
	"github.com/jackc/pgx/v5"
)

type PayOrderUsecase struct {
	logger    ports.Logger
	txManager ports.TxManager
}

func NewPayOrderUsecase(logger ports.Logger, manager ports.TxManager) *PayOrderUsecase {
	return &PayOrderUsecase{
		logger:    logger,
		txManager: manager,
	}
}

func (u *PayOrderUsecase) ProcessOrder(
	ctx context.Context,
	orderID uuid.UUID,
	payload *domain.Order,
) error {
	return u.txManager.WithTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		balances := postgres.NewBalancesRepository(tx)
		payments := postgres.NewPaymentsRepository(tx)
		outbox := postgres.NewOutboxRepository(tx)

		raw := domain.NewPayment(
			orderID,
			payload.UserID,
			payload.Amount,
		)
		payment, err := payments.CreatePayment(ctx, raw)
		if err != nil {
			if errors.Is(err, domain.ErrPaymentAlreadyExists) {
				return nil
			}

			if errors.Is(err, domain.ErrInvalidData) {
				data, err := encodePaymentFailedPayload(
					orderID,
					payload.UserID,
					payload.Amount,
					err.Error(),
				)
				if err != nil {
					u.logger.Error(
						"failed to encode payment failed payload",
						"order id", orderID.String(),
						"error", err,
					)
					return domain.ErrInternal
				}

				event := domain.NewEvent(orderID, data)
				if err := outbox.CreateEvent(ctx, event); err != nil {
					u.logger.Error(
						"failed to create payment failed event",
						"order id", orderID.String(),
						"error", err,
					)

					return domain.ErrInternal
				}

				return nil
			}

			u.logger.Error(
				"failed to create payment",
				"order id", orderID.String(),
				"error", err,
			)
			return domain.ErrInternal
		}

		if err := balances.PayOrder(
			ctx,
			payload.UserID,
			payload.Amount,
		); err != nil {
			if errors.Is(err, domain.ErrInvalidBalance) {
				data, err := encodePaymentFailedPayload(
					orderID,
					payload.UserID,
					payload.Amount,
					err.Error(),
				)
				if err != nil {
					u.logger.Error(
						"failed to encode payment failed payload",
						"order id", orderID.String(),
						"error", err,
					)
					return domain.ErrInternal
				}

				event := domain.NewEvent(orderID, data)
				if err := outbox.CreateEvent(ctx, event); err != nil {
					u.logger.Error(
						"failed to create payment failed event",
						"order id", orderID.String(),
						"error", err,
					)
					return domain.ErrInternal
				}

				return nil
			}

			u.logger.Error(
				"failed to pay order",
				"order id", orderID.String(),
				"error", err,
			)
			return domain.ErrInternal
		}

		data, err := encodePaymentPayload(payment)
		if err != nil {
			u.logger.Error(
				"failed to encode payment payload",
				"order id", orderID.String(),
				"error", err,
			)
			return domain.ErrInternal
		}

		event := domain.NewEvent(orderID, data)
		if err := outbox.CreateEvent(ctx, event); err != nil {
			u.logger.Error(
				"failed to create payment event",
				"order id", orderID.String(),
				"error", err,
			)
			return domain.ErrInternal
		}

		return nil
	})
}
