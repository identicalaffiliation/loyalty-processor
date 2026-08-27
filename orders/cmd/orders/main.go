package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/identicalaffiliation/loyalty-processor/orders/internal/adapters/kafka"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/adapters/rpc"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/application"
	"github.com/identicalaffiliation/loyalty-processor/orders/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/orders/pkg/httpserver"
	"github.com/identicalaffiliation/loyalty-processor/orders/pkg/logger"
	"github.com/identicalaffiliation/loyalty-processor/orders/pkg/outboxworker"
	"github.com/identicalaffiliation/loyalty-processor/orders/pkg/psqlpool"
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

	postgresPool, err := psqlpool.NewPostgresPool(context.Background(), &cfg.PostgresConfig)
	if err != nil {
		slogger.Error(
			"failed to create postgres pool",
			"error", err,
		)
		os.Exit(1)
	}

	defer postgresPool.Close()

	if err := postgresPool.Ping(context.Background()); err != nil {
		slogger.Error(
			"failed to ping postgres pool",
			"error", err,
		)
		os.Exit(1)
	}

	sender := kafka.NewProducer(&cfg.KafkaConfig)
	defer func() {
		if err := sender.Close(); err != nil {
			slogger.Error(
				"failed to close sender",
				"error", err,
			)
		}
	}()

	wg := &sync.WaitGroup{}
	outbox := postgres.NewOutboxRepository(postgresPool)

	txManager := postgres.NewTxManager(postgresPool)

	client, cleanup, err := rpc.NewInventoryClient(&cfg.InventoryServiceConfig)
	if err != nil {
		slogger.Error(
			"failed to create client",
			"error", err,
		)
		os.Exit(1)
	}

	defer cleanup()

	create := application.NewCreateOrderUsecase(txManager, slogger, client)

	server := httpserver.RegisterRoutes(&cfg.ServerConfig, create)

	go func() {
		slogger.Debug("server started listen port", cfg.ServerConfig.Port)
		if err := server.Start(server.Server.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slogger.Error(
				"failed to start server",
				"error", err,
			)
		}
	}()

	notify, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := outboxworker.Run(
			notify,
			cfg.OutboxConfig.Limit,
			cfg.OutboxConfig.Tick,
			outbox,
			sender,
			slogger,
		)
		if err != nil {
			slogger.Debug("worker was stopped by notify signal")
		}
	}()

	<-notify.Done()
	wg.Wait()
	
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ServerConfig.ShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slogger.Error(
			"failed to shutdown server",
			"error", err,
		)
		os.Exit(1)
	}
}
