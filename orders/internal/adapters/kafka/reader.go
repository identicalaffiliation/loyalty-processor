package kafka

import (
	"context"
	"errors"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/ports"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	logger ports.Logger
}

func NewConsumer(cfg *config.KafkaConfig, logger ports.Logger) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:          cfg.Brokers,
			GroupID:          cfg.ID,
			Topic:            cfg.ConsumeTopic,
			Partition:        cfg.Partition,
			QueueCapacity:    cfg.QueueCapacity,
			ReadBatchTimeout: cfg.BatchTimeout,
		}),
		logger: logger,
	}
}

func (c *Consumer) ReadMessages(ctx context.Context, usecase ports.ProcessOrderUsecase) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return ctx.Err()
			}

			c.logger.Error(
				"failed to fetch message from kafka",
				"error", err,
			)
			continue
		}

		err = usecase.ProcessOrder(ctx, msg.Key, msg.Value)
		if err != nil {
			if errors.Is(err, domain.ErrInvalidData) {
				if err := c.commitMessage(ctx, msg); err != nil {
					return err
				}

				continue
			}

			c.logger.Error(
				"failed to process kafka message",
				"error", err,
			)
			continue
		}

		if err := c.commitMessage(ctx, msg); err != nil {
			return err
		}
	}
}

func (c *Consumer) commitMessage(ctx context.Context, msg kafka.Message) error {
	if err := c.reader.CommitMessages(ctx, msg); err != nil {
		c.logger.Error(
			"failed to commit kafka message",
			"error", err,
		)

		return fmt.Errorf("commit message: %w", err)
	}

	return nil
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
