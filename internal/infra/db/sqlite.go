package db

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

type (
	uid       = uuid.UUID
	text      = string
	integer   = int64
	timestamp = time.Time
)

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
