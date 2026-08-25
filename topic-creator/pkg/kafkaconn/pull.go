package kafkaconn

import (
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/topic-creator/internal/config"
	"github.com/segmentio/kafka-go"
)

func GetKafkaConn(cfg *config.KafkaConfig) (*kafka.Conn, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	return kafka.Dial(cfg.ConnectionType, addr)
}

func GetKafkaControllerConn(controller *kafka.Broker, cfg *config.KafkaConfig) (*kafka.Conn, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, controller.Port)
	return kafka.Dial(cfg.ConnectionType, addr)
}
