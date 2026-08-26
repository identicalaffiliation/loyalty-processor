package application

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
)

func (u *InventoryUsecase) ReserveStock(ctx context.Context, productID uuid.UUID) error {
	if err := u.repo.ReserveStock(ctx, productID); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) || errors.Is(err, domain.ErrOutOfStock) {
			return err
		}

		u.log.Error(
			"failed to reserve stock",
			"error", err,
		)
		return err
	}

	return nil
}
