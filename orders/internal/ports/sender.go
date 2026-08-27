package ports

import (
	"context"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
)

type Sender interface {
	WriteMessages(ctx context.Context, messages []*domain.Event) error
}
