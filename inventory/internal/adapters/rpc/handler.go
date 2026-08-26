package rpc

import (
	inventoryv1 "github.com/identicalaffiliation/loyalty-processor/gen/inventory/v1"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/ports"
)

type GRPCHandler struct {
	cases ports.InventoryApplication
	inventoryv1.UnsafeInventoryServiceServer
}

func NewHandler(cases ports.InventoryApplication) *GRPCHandler {
	return &GRPCHandler{
		cases: cases,
	}
}
