package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/identicalaffiliation/loyalty-processor/migrator/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/migrator/internal/initializer"
	"github.com/identicalaffiliation/loyalty-processor/migrator/pkg/logger"
	"github.com/identicalaffiliation/loyalty-processor/migrator/pkg/psqlpool"
)

const (
	up    = "up"
	down  = "down"
	reset = "reset"
)

func main() {
	var (
		configPath string
		command    string
	)
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.StringVar(&command, "command", "", "command to migrations (up,down, reset)")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	slogger, err := logger.NewLogger(&cfg.LoggerConfig)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := psqlpool.NewPostgresPool(cfg)
	if err != nil {
		slogger.Error(
			"failed to create postgres pool",
			"error", err,
		)
		log.Fatal(err)
	}

	defer func() {
		if err := pool.Close(); err != nil {
			slogger.Error(
				"failed to close postgres pool",
				"error", err,
			)
		}
	}()

	ctx, stop := context.WithTimeout(context.Background(), time.Second*15)
	defer stop()

	switch command {
	case up:
		if err := initializer.MigrateUp(ctx, cfg, pool); err != nil {
			slogger.Error(
				"failed to migrate up",
				"error", err,
			)
			os.Exit(1)
		}
	case down:
		if err := initializer.MigrateDown(ctx, cfg, pool); err != nil {
			slogger.Error(
				"failed to migrate down",
				"error", err,
			)
			os.Exit(1)
		}
	case reset:
		if err := initializer.MigrateReset(ctx, cfg, pool); err != nil {
			slogger.Error(
				"failed to migrate reset",
				"error", err,
			)
			os.Exit(1)
		}
	default:
		slogger.Error("invalid migrate command")
		os.Exit(1)
	}

	slogger.Debug("migrations was init successfully!")
}
