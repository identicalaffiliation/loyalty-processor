package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/ports"
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

func (c *Consumer) ReadMessages(ctx context.Context, usecase ports.PayOrderUsecase) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return ctx.Err()
			}

			c.logger.Error(
				"failed to fetch kafka message",
				"error", err,
			)
			continue
		}

		orderID, err := uuid.ParseBytes(msg.Key)
		if err != nil {
			c.logger.Error(
				"failed to parse key",
				"error", err,
			)
			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				return fmt.Errorf("commit invalid message: %w", err)
			}

			continue
		}

		var payload domain.Order
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			c.logger.Error(
				"failed to parse payload",
				"error", err,
			)

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				return fmt.Errorf("commit invalid message: %w", err)
			}
			continue
		}

		if err := usecase.ProcessOrder(ctx, orderID, &payload); err != nil {
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("commit messages: %w", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
