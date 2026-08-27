package application

import (
	"context"
	"errors"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/input"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/output"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
	"github.com/jackc/pgx/v5"
)

type CreateOrderUsecase struct {
	manager         ports.TxManager
	logger          ports.Logger
	inventoryClient ports.InventoryClient
}

func NewCreateOrderUsecase(
	manager ports.TxManager,
	logger ports.Logger,
	client ports.InventoryClient,
) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		manager:         manager,
		logger:          logger,
		inventoryClient: client,
	}
}

func (u *CreateOrderUsecase) CreateOrder(
	ctx context.Context,
	req *input.CreateOrderRequest,
) (*output.CreateOrderResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if err := u.inventoryClient.Reserve(ctx, req.ProductID); err != nil {
		if !errors.Is(err, domain.ErrInternal) {
			return nil, err
		}

		u.logger.Error(
			"failed to reserve stock",
			"product id", req.ProductID,
			"error", err,
		)
		return nil, domain.ErrInternal
	}

	var created *domain.Order

	err := u.manager.WithTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		orders := postgres.NewOrdersRepository(tx)
		outbox := postgres.NewOutboxRepository(tx)

		order := domain.NewOrder(
			req.UserID,
			req.ProductID,
			req.Amount,
			domain.StatusCreated,
		)
		raw, err := orders.CreateOrder(ctx, order)
		if err != nil {
			return err
		}

		payload, err := encodeOrderPayload(raw)
		if err != nil {
			return err
		}

		event := domain.NewEvent(raw.ID, domain.Created, payload)
		if err := outbox.CreateEvent(ctx, event); err != nil {
			return err
		}

		created = raw
		return nil
	})
	if err != nil {
		if err := u.inventoryClient.Release(ctx, req.ProductID); err != nil {
			u.logger.Error(
				"failed to release stock",
				"product id", req.ProductID.String(),
				"error", err,
			)
		}

		if errors.Is(err, domain.ErrInvalidData) {
			return nil, err
		}

		u.logger.Error(
			"failed to create order",
			"error", err,
		)
		return nil, domain.ErrInternal
	}

	return output.NewCreateOrderResponse(created), nil
}
