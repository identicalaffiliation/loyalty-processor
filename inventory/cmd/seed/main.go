package main

import (
	"context"
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/adapters/postgres"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/domain"
	"github.com/identicalaffiliation/loyalty-processor/invetory/pkg/psqlpool"
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

	repo := postgres.NewInventoryRepository(pool)
	seed := []*domain.Product{
		{
			ID:    uuid.New(),
			Title: "Iphone 17",
			Stock: 5,
		},
		{
			ID:          uuid.New(),
			Title:       "Black T-shirt",
			Description: getPtr("simple black t"),
			Stock:       1,
		},
		{
			ID:    uuid.New(),
			Title: "Пуховик",
			Stock: 100,
		},
		{
			ID:          uuid.New(),
			Title:       "Reno logan",
			Description: getPtr("100km"),
			Stock:       2,
		},
	}

	for _, data := range seed {
		err = repo.CreateProduct(ctx, data)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed data was add successfully!")
}

func getPtr(str string) *string {
	return &str
}
