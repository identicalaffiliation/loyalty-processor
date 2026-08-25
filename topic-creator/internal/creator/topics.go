package creator

import (
	"errors"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/topic-creator/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/topic-creator/internal/ports"
	"github.com/identicalaffiliation/loyalty-processor/topic-creator/pkg/kafkaconn"
	"github.com/segmentio/kafka-go"
)

func CreateKafkaTopics(logger ports.Logger, cfg *config.KafkaConfig) error {
	conn, err := kafkaconn.GetKafkaConn(cfg)
	if err != nil {
		return fmt.Errorf("get kafka conn: %w", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			logger.Error(
				"failed to close kafka conn",
				"error", err,
			)
		}
	}()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("get kafka controller: %w", err)
	}

	controllerConn, err := kafkaconn.GetKafkaControllerConn(&controller, cfg)
	if err != nil {
		return fmt.Errorf("get kafka controller conn: %w", err)
	}

	defer func() {
		if err := controllerConn.Close(); err != nil {
			logger.Error(
				"failed to close kafka controller conn",
				"error", err,
			)
		}
	}()

	return createTopics(controllerConn, cfg)
}

func createTopics(controllerConn *kafka.Conn, cfg *config.KafkaConfig) error {
	for _, topic := range cfg.Topics {
		err := controllerConn.CreateTopics(
			kafka.TopicConfig{
				Topic:             topic,
				NumPartitions:     cfg.Partitions,
				ReplicationFactor: cfg.Replicas,
			},
		)
		if err != nil && !errors.Is(err, kafka.TopicAlreadyExists) {
			return fmt.Errorf("create kafka topic: %w", err)
		}
	}

	return nil
}
