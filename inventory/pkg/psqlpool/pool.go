package psqlpool

import (
	"context"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/invetory/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(cfg.DbURL)
	if err != nil {
		return nil, fmt.Errorf("parse dsn: %w", err)
	}

	conf.MaxConns = cfg.MaxConnections
	conf.MaxConnLifetime = cfg.MaxLifetime
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	return pool, nil
}
