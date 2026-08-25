package config

type Config struct {
	LoggerConfig LoggerConfig `yaml:"logger"`
	KafkaConfig  KafkaConfig  `yaml:"kafka"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type KafkaConfig struct {
	ConnectionType string   `yaml:"conn_type"`
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	Topics         []string `yaml:"topics"`
	Partitions     int      `yaml:"partitions"`
	Replicas       int      `yaml:"replicas"`
}
