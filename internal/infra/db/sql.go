package db

import (
	"database/sql"

	"github.com/shimeoki/wp/internal/domain"
)

type DB interface {
	ExecContext(ctx domain.Ctx, query string, args ...any) (sql.Result, error)
	QueryContext(ctx domain.Ctx, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx domain.Ctx, query string, args ...any) *sql.Row
}
