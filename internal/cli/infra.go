package cli

import (
	"context"
	"database/sql"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

func openStore() *store.LocalStore {
	hasher := &store.SHA256Hasher{}

	s, err := store.NewLocalStore(cfg.Store.Path, hasher)
	if err != nil {
		fatal(err)
	}

	return s
}

func openDB(ctx context.Context) *sql.DB {
	db, err := sqlite.Open(ctx, cfg.DB.DataSourceName)
	if err != nil {
		fatal(err)
	}

	return db
}
