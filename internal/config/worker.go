package config

import (
	"database/sql"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type Worker struct {
	db    *sql.DB
	store *store.LocalStore
}

func (w *Worker) Work(
	ctx app.Ctx,
	j app.Job[*Provider],
) error {
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	if err := j(&Provider{
		store:      w.store,
		tags:       sqlite.NewTagRepo(tx),
		wallpapers: sqlite.NewWallpaperRepo(tx),
	}); err != nil {
		return err
	}

	return tx.Commit()
}
