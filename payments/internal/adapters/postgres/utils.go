package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	constraitCode       = "23514"
	uniqueViolationCode = "23505"
)

func checkConstraits(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == constraitCode {
			return true
		}
	}

	return false
}

func checkUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == uniqueViolationCode {
			return true
		}
	}

	return false
}
