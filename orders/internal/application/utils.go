package application

import (
	"encoding/json"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

func encodeOrderPayload(order *domain.Order) ([]byte, error) {
	payload, err := json.Marshal(order)
	if err != nil {
		return nil, fmt.Errorf("encode order payload: %w", err)
	}

	return payload, nil
}
