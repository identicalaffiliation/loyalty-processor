package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	inventoryv1 "github.com/identicalaffiliation/loyalty-processor/gen/inventory/v1"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/adapters/rpc"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/application"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/invetory/pkg/logger"
	"github.com/identicalaffiliation/loyalty-processor/invetory/pkg/psqlpool"
	"google.golang.org/grpc"
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

	if err := postgresPool.Ping(ctx); err != nil {
		slogger.Error(
			"failed to ping postgres pool",
			"error", err,
		)
		os.Exit(1)
	}

	repo := postgres.NewInventoryRepository(postgresPool)
	cases := application.NewUsecase(repo, slogger)

	server := grpc.NewServer()
	handler := rpc.NewHandler(cases)
	inventoryv1.RegisterInventoryServiceServer(server, handler)

	listener, err := net.Listen(
		cfg.ServerConfig.ConnectionType,
		fmt.Sprintf(":%d", cfg.ServerConfig.Port),
	)
	if err != nil {
		slogger.Error(
			"failed to create listener",
			"error", err,
		)
		os.Exit(1)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			slogger.Error(
				"failed to close listener",
				"error", err,
			)
		}
	}()

	go func() {
		slogger.Debug("server started on port", cfg.ServerConfig.Port)
		if err := server.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			slogger.Error(
				"failed to server listener",
				"error", err,
			)
		}
	}()

	notifyCtx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-notifyCtx.Done()

	slogger.Debug("server is stopping gracefully!")
	server.GracefulStop()
	slogger.Debug("server was stopped gracefully successful!")
}
