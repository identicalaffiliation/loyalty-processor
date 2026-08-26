package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const constraitCode = "23514"

func checkConstraits(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == constraitCode {
			return true
		}
	}

	return false
}
