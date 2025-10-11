package db

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/shimeoki/wp/internal/config"
	"github.com/shimeoki/wp/internal/v2/domain"
	_ "modernc.org/sqlite"
)

type DB interface {
	ExecContext(ctx domain.Ctx, query string, args ...any) (sql.Result, error)
	QueryContext(ctx domain.Ctx, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx domain.Ctx, query string, args ...any) *sql.Row
}

//go:embed sqlite.sql
var sqliteScheme string

func OpenSQLiteDB(ctx domain.Ctx, cfg *config.DB) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err := ping(db); err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, sqliteScheme); err != nil {
		return nil, err
	}

	return db, nil
}

func ping(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}
