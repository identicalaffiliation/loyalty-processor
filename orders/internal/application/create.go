package application

import (
	"context"
	"errors"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/input"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/output"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
)

type CreateOrderUsecase struct {
	ordersRepo      ports.OrdersRepository
	logger          ports.Logger
	inventoryClient ports.InventoryClient
}

func NewCreateOrderUsecase(
	ordersRepo ports.OrdersRepository,
	logger ports.Logger,
	client ports.InventoryClient,
) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		ordersRepo:      ordersRepo,
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

	order, err := u.ordersRepo.CreateOrder(ctx, domain.NewOrder(
		req.UserID,
		req.ProductID,
		req.Amount,
		domain.StatusCreated,
	))
	if err != nil {
		if err := u.inventoryClient.Release(ctx, req.ProductID); err != nil {
			u.logger.Error(
				"failed to release stock",
				"product id", req.ProductID,
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

	return output.NewCreateOrderResponse(order), nil
}
