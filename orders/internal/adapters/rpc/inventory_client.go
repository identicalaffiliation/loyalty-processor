package rpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	inventoryv1 "github.com/identicalaffiliation/loyalty-processor/gen/inventory/v1"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/pkg/grpcclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryClient struct {
	client inventoryv1.InventoryServiceClient
}

func NewInventoryClient(cfg *config.InventoryServiceConfig) (*InventoryClient, func(), error) {
	client, err := grpcclient.NewInventoryGRPClient(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create inventory grpc conn: %w", err)
	}

	cleanup := func() {
		if err := client.Close(); err != nil {
			fmt.Println("grpc client conn close error", err)
		}
	}

	return &InventoryClient{client: inventoryv1.NewInventoryServiceClient(client)}, cleanup, nil
}

func (c *InventoryClient) Reserve(ctx context.Context, productID uuid.UUID) error {
	in := &inventoryv1.ReserveStockRequest{
		ProductId: productID.String(),
	}
	_, err := c.client.ReserveStock(ctx, in)
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return domain.ErrProductNotFound
		case codes.FailedPrecondition:
			return domain.ErrOutOfStock
		default:
			return domain.ErrInternal
		}
	}

	return nil
}

func (c *InventoryClient) Release(ctx context.Context, productID uuid.UUID) error {
	in := &inventoryv1.ReleaseStockRequest{
		ProductId: productID.String(),
	}

	if _, err := c.client.ReleaseStock(ctx, in); err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return domain.ErrProductNotFound
		default:
			return domain.ErrInternal
		}
	}

	return nil
}
