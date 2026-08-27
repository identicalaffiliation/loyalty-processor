package application

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
)

func encodePaymentPayload(payment *domain.Payment) ([]byte, error) {
	payload, err := json.Marshal(payment)
	if err != nil {
		return nil, fmt.Errorf("encode payment payload: %w", err)
	}

	return payload, nil
}

func encodePaymentFailedPayload(
	orderID uuid.UUID,
	userID uuid.UUID,
	amount int64,
	reason string,
) ([]byte, error) {
	payload := domain.PaymentFailedPayload{
		OrderID: orderID,
		UserID:  userID,
		Amount:  amount,
		Reason:  reason,
	}

	return json.Marshal(payload)
}
