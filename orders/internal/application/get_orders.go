package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/dto/output"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
)

type GetOrdersUsecase struct {
	repo   ports.OrdersRepository
	logger ports.Logger
}

func NewGetOrdersUsecase(repo ports.OrdersRepository, logger ports.Logger) *GetOrdersUsecase {
	return &GetOrdersUsecase{
		repo:   repo,
		logger: logger,
	}
}

func (u *GetOrdersUsecase) GetOrdersByUser(ctx context.Context, id uuid.UUID) (*output.OrdersResponse, error) {
	if id == uuid.Nil {
		return nil, domain.ErrInvalidData
	}

	orders, err := u.repo.GetOrdersByUser(ctx, id)
	if err != nil {
		u.logger.Error(
			"failed to get orders by user id",
			"user id", id.String(),
			"error", err,
		)
		return nil, domain.ErrInternal
	}

	return output.NewOrdersResponse(orders), nil
}
