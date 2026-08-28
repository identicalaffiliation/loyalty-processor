package outboxworker

import (
	"context"
	"time"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
)

func Run(
	ctx context.Context,
	limit int,
	tick time.Duration,
	outbox ports.OutboxRepository,
	sender ports.Sender,
	logger ports.Logger,
) error {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			events, err := outbox.GetUnpublishedEventsByLimit(ctx, limit)
			if err != nil {
				logger.Error(
					"failed to batch events get",
					"error", err,
				)
				continue
			}

			if len(events) == 0 {
				continue
			}

			if err := sender.WriteMessages(ctx, events); err != nil {
				logger.Error(
					"failed to batch send events",
					"error", err,
				)
				continue
			}

			for _, event := range events {
				if err := outbox.MarkPublished(ctx, event.ID, domain.Published); err != nil {
					logger.Error(
						"failed to mark published event",
						"event id", event.ID.String(),
						"error", err,
					)
				}
			}
		}
	}
}
