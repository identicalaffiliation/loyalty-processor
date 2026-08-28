package config

import "time"

type Config struct {
	LoggerConfig           LoggerConfig           `yaml:"logger"`
	PostgresConfig         PostgresConfig         `yaml:"postgres"`
	ServerConfig           ServerConfig           `yaml:"server"`
	InventoryServiceConfig InventoryServiceConfig `yaml:"inventory_service"`
	KafkaConfig            KafkaConfig            `yaml:"kafka"`
	OutboxConfig           OutboxConfig           `yaml:"outbox"`
}

type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	IdleTimeout     time.Duration `yaml:"idle_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type InventoryServiceConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PostgresConfig struct {
	DbURL          string        `env:"DB_URL"`
	MaxLifetime    time.Duration `yaml:"max_lifetime"`
	MaxConnections int32         `yaml:"max_conns"`
}

type KafkaConfig struct {
	Brokers       []string      `yaml:"brokers"`
	ConsumeTopic  string        `yaml:"consume_topic"`
	ProduceTopic  string        `yaml:"produce_topic"`
	ID            string        `yaml:"id"`
	QueueCapacity int           `yaml:"capacity"`
	BatchSize     int           `yaml:"batch_size"`
	MaxAttempts   int           `yaml:"attempts"`
	Partition     int           `yaml:"partition"`
	BatchTimeout  time.Duration `yaml:"batch_timeout"`
	ReadTimeout   time.Duration `yaml:"read_timeout"`
	WriteTimeout  time.Duration `yaml:"write_timeout"`
}

type OutboxConfig struct {
	Limit int           `yaml:"limit"`
	Tick  time.Duration `yaml:"tick"`
}
