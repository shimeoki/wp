package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Ctx = context.Context

type DB interface {
	ExecContext(ctx Ctx, query string, args ...any) (sql.Result, error)
	QueryContext(ctx Ctx, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx Ctx, query string, args ...any) *sql.Row
}

type (
	uid       = uuid.UUID
	text      = string
	integer   = int64
	timestamp = time.Time
)

//go:embed schema.sql
var schema string

func Open(ctx Ctx, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
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
