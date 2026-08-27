package input

import (
	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type CreateOrderRequest struct {
	UserID    uuid.UUID
	ProductID uuid.UUID
	Amount    int64 `json:"amount"`
}

func NewCreateOrderRequest(userID, productID uuid.UUID, amount int64) *CreateOrderRequest {
	return &CreateOrderRequest{
		UserID:    userID,
		ProductID: productID,
		Amount:    amount,
	}
}

func (in *CreateOrderRequest) Validate() error {
	if in.UserID == uuid.Nil {
		return domain.ErrInvalidData
	}

	if in.ProductID == uuid.Nil {
		return domain.ErrInvalidData
	}

	if in.Amount < 1 {
		return domain.ErrInvalidData
	}

	return nil
}
