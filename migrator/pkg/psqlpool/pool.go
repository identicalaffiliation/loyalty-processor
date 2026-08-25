package psqlpool

import (
	"database/sql"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/migrator/internal/config"
	_ "github.com/lib/pq"
)

const driver string = "postgres"

func NewPostgresPool(cfg *config.Config) (*sql.DB, error) {
	pool, err := sql.Open(driver, cfg.DbURL)
	if err != nil {
		return nil, fmt.Errorf("open psql pool: %w", err)
	}

	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("ping psql pool: %w", err)
	}

	return pool, nil
}
