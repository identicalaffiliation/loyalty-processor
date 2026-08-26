package config

import "time"

type Config struct {
	LoggerConfig   LoggerConfig   `yaml:"logger"`
	PostgresConfig PostgresConfig `yaml:"postgres"`
	ServerConfig   ServerConfig   `yaml:"server"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type PostgresConfig struct {
	DbURL          string        `env:"DB_URL"`
	MaxLifetime    time.Duration `yaml:"max_lifetime"`
	MaxConnections int32         `yaml:"max_conns"`
}

type ServerConfig struct {
	Port           int    `yaml:"port"`
	ConnectionType string `yaml:"connection_type"`
}
