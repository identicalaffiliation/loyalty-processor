package kafka

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/payments/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg *config.KafkaConfig) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(cfg.Brokers...),
			Topic:        cfg.ProduceTopic,
			MaxAttempts:  cfg.MaxAttempts,
			BatchSize:    cfg.BatchSize,
			BatchTimeout: cfg.BatchTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}

func (p *Producer) WriteMessages(ctx context.Context, events []*domain.Event) error {
	messages := make([]kafka.Message, 0, len(events))
	for _, event := range events {
		messages = append(messages, kafka.Message{
			Key:   []byte(event.OrderID.String()),
			Value: event.Payload,
		})
	}

	if err := p.writer.WriteMessages(ctx, messages...); err != nil {
		return fmt.Errorf("write batch messages to kafka: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
