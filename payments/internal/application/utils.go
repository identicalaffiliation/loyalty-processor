package application

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
)

func encodePaymentPayload(payment *domain.Message) ([]byte, error) {
	payload, err := json.Marshal(payment)
	if err != nil {
		return nil, fmt.Errorf("encode payment payload: %w", err)
	}

	return payload, nil
}

func encodePaymentFailedPayload(
	orderID uuid.UUID,
	userID uuid.UUID,
) ([]byte, error) {
	payload := domain.Message{
		Status:  domain.Fail,
		OrderID: orderID,
		UserID:  userID,
	}

	return json.Marshal(payload)
}
