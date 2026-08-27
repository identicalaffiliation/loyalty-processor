package main

import (
	"context"
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/payments/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/payments/pkg/psqlpool"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	ctx := context.Background()
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := psqlpool.NewPostgresPool(ctx, &cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	repo := postgres.NewBalancesRepository(pool)
	seed := []*domain.Balance{
		{
			UserID:  uuid.New(),
			Bonuses: 10000,
		},
	}

	for _, data := range seed {
		if err := repo.CreateBalance(ctx, data); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed data was add successfully!")
}
