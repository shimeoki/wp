package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/google/uuid"
	"github.com/shimeoki/wp/internal/config"
	"github.com/shimeoki/wp/internal/domain"
	_ "modernc.org/sqlite"
)

type DB interface {
	ExecContext(ctx domain.Ctx, query string, args ...any) (sql.Result, error)
	QueryContext(ctx domain.Ctx, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx domain.Ctx, query string, args ...any) *sql.Row
}

type (
	uid       = uuid.UUID
	text      = string
	integer   = int64
	timestamp = time.Time
)

//go:embed schema.sql
var schema string

func Open(ctx domain.Ctx, cfg *config.DB) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err := ping(db); err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	return db, nil
}

func ping(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}
