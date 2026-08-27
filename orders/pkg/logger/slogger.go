package logger

import (
	"log/slog"
	"os"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(cfg *config.LoggerConfig) (*Logger, error) {
	levels := map[string]slog.Level{
		LevelDebug: slog.LevelDebug,
		LevelInfo:  slog.LevelInfo,
		LevelError: slog.LevelError,
		LevelWarn:  slog.LevelWarn,
	}

	level, ok := levels[cfg.Level]
	if !ok {
		return nil, ErrInvalidLevel
	}

	handlers := map[string]slog.Handler{
		TextFormat: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),

		JsonFormat: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	}

	handler, ok := handlers[cfg.Format]
	if !ok {
		return nil, ErrInvalidFormat
	}

	logger := &Logger{
		logger: slog.New(handler),
	}

	return logger, nil
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}
