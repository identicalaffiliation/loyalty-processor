package initializer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/identicalaffiliation/loyalty-processor/migrator/internal/config"
	"github.com/pressly/goose/v3"
)

func MigrateUp(ctx context.Context, cfg *config.Config, db *sql.DB) error {
	if err := goose.UpContext(ctx, db, cfg.MigrationsPath); err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}

	return nil
}

func MigrateDown(ctx context.Context, cfg *config.Config, db *sql.DB) error {
	if err := goose.DownContext(ctx, db, cfg.MigrationsPath); err != nil {
		return fmt.Errorf("down migrations: %w", err)
	}

	return nil
}

func MigrateReset(ctx context.Context, cfg *config.Config, db *sql.DB) error {
	if err := goose.ResetContext(ctx, db, cfg.MigrationsPath); err != nil {
		return fmt.Errorf("reset migrations: %w", err)
	}

	return nil
}
