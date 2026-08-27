package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/identicalaffiliation/loyalty-processor/payments/internal/adapters/kafka"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/application"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/payments/pkg/logger"
	"github.com/identicalaffiliation/loyalty-processor/payments/pkg/outboxworker"
	"github.com/identicalaffiliation/loyalty-processor/payments/pkg/psqlpool"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	slogger, err := logger.NewLogger(&cfg.LoggerConfig)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	postgresPool, err := psqlpool.NewPostgresPool(ctx, &cfg.PostgresConfig)
	if err != nil {
		slogger.Error(
			"failed to create postgres pool",
			"error", err,
		)
		os.Exit(1)
	}

	defer postgresPool.Close()

	txManager := postgres.NewTxManager(postgresPool)

	usecase := application.NewPayOrderUsecase(slogger, txManager)

	reader := kafka.NewConsumer(&cfg.KafkaConfig, slogger)
	writer := kafka.NewProducer(&cfg.KafkaConfig)

	wg := &sync.WaitGroup{}

	defer func(writer *kafka.Producer) {
		err := writer.Close()
		if err != nil {
			slogger.Error(
				"failed to close kafka producer",
				"error", err,
			)
		}
	}(writer)

	notify, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := outboxworker.Run(
			notify,
			cfg.OutboxConfig.Limit,
			cfg.OutboxConfig.Tick,
			postgres.NewOutboxRepository(postgresPool),
			writer,
			slogger,
		)
		if err != nil || !errors.Is(err, context.Canceled) {
			slogger.Error("worker fail", "error", err)
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := reader.ReadMessages(notify, usecase)
		if err != nil || !errors.Is(err, context.Canceled) {
			slogger.Error("reader fail", "error", err)
		}
	}()

	wg.Wait()
}
