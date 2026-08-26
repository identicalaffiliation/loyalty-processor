package application

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/ports"
)

type InventoryUsecase struct {
	repo ports.InventoryRepository
	log  ports.Logger
}

func NewUsecase(repo ports.InventoryRepository, log ports.Logger) *InventoryUsecase {
	return &InventoryUsecase{
		repo: repo,
		log:  log,
	}
}

func (u *InventoryUsecase) ReleaseStock(ctx context.Context, productID uuid.UUID) error {
	if err := u.repo.ReleaseStock(ctx, productID); err != nil {
		if !errors.Is(err, domain.ErrProductNotFound) {
			u.log.Error(
				"failed to release stock",
				"error", err,
			)

			return domain.ErrInternal
		}

		return err
	}

	return nil
}
