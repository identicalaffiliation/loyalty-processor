package grpcclient

import (
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInventoryGRPClient(cfg *config.InventoryServiceConfig) (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
