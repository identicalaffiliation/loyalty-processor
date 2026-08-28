package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
)

type ProcessOrderUsecase struct {
	inventoryClient ports.InventoryClient
	repo            ports.OrdersRepository
	logger          ports.Logger
}

func NewProcessPaymentUsecase(
	client ports.InventoryClient,
	repository ports.OrdersRepository,
	logger ports.Logger,
) *ProcessOrderUsecase {
	return &ProcessOrderUsecase{
		inventoryClient: client,
		repo:            repository,
		logger:          logger,
	}
}

func (u *ProcessOrderUsecase) ProcessOrder(ctx context.Context, key []byte, payload []byte) error {
	id, err := uuid.ParseBytes(key)
	if err != nil {
		u.logger.Error(
			"failed to parse kafka key",
			"error", err,
		)

		return domain.ErrInternal
	}

	status, err := u.getOrderStatusByEventType(payload)
	if err != nil {
		u.logger.Error(
			"failed to get order status by event type",
			"error", err,
		)

		return domain.ErrInternal
	}

	if status != domain.StatusPaid {
		if err := u.repo.MarkOrder(ctx, id, domain.StatusFailed); err != nil {
			if errors.Is(err, domain.ErrInvalidData) {
				return err
			}

			u.logger.Error(
				"failed to mark order",
				"error", err,
			)
			return domain.ErrInternal
		}

		productID, err := u.repo.GetProductByOrderID(ctx, id)
		if err != nil {
			if errors.Is(err, domain.ErrProductNotFound) {
				return err
			}

			u.logger.Error(
				"failed to get product by order id",
				"error", err,
			)
			return domain.ErrInternal
		}

		if err = u.inventoryClient.Release(ctx, *productID); err != nil {
			if errors.Is(err, domain.ErrInvalidData) {
				return err
			}

			u.logger.Error(
				"failed to release stock",
				"product id", productID.String(),
				"error", err,
			)
			return domain.ErrInternal
		}

		return nil
	}

	return u.repo.MarkOrder(ctx, id, domain.StatusPaid)
}

func (u *ProcessOrderUsecase) getOrderStatusByEventType(payload []byte) (domain.OrderStatus, error) {
	var msg domain.Message
	if err := json.Unmarshal(payload, &msg); err != nil {
		return "", fmt.Errorf("unmarshal payload: %w", err)
	}

	if msg.Status == domain.Success {
		return domain.StatusPaid, nil
	}

	return domain.StatusFailed, nil
}
