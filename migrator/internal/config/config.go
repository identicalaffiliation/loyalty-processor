package config

type Config struct {
	LoggerConfig   LoggerConfig `yaml:"logger"`
	DbURL          string       `env:"DB_URL"`
	MigrationsPath string       `env:"MIGRATIONS_PATH"`
}

type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}
