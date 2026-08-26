package rpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	inventoryv1 "github.com/identicalaffiliation/loyalty-processor/gen/inventory/v1"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *GRPCHandler) ReleaseStock(
	ctx context.Context,
	req *inventoryv1.ReleaseStockRequest,
) (*inventoryv1.ReleaseStockResponse, error) {
	id, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid product id")
	}

	if err := h.cases.ReleaseStock(ctx, id); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &inventoryv1.ReleaseStockResponse{}, nil
}
