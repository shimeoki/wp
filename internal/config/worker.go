package config

import (
	"database/sql"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type LocalSQLiteWorker struct {
	db    *sql.DB
	store *store.LocalStore
}

func (w *LocalSQLiteWorker) Work(ctx app.Ctx, j app.Job[*Provider]) error {
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	if err := j(&Provider{
		store:      w.store,
		tags:       sqlite.NewTagRepo(tx),
		wallpapers: sqlite.NewWallpaperRepo(tx),
		sources:    sqlite.NewSourceRepo(tx),
		aliases:    sqlite.NewAliasRepo(tx),
	}); err != nil {
		return err
	}

	return tx.Commit()
}

func (w *LocalSQLiteWorker) Close() error {
	return w.store.Close()
}
