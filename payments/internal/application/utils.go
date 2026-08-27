package application

import (
	"encoding/json"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
)

func encodePaymentPayload(payment *domain.Payment) ([]byte, error) {
	payload, err := json.Marshal(payment)
	if err != nil {
		return nil, fmt.Errorf("encode payment payload: %w", err)
	}

	return payload, nil
}
