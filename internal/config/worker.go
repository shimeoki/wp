package config

import (
	"database/sql"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/domain"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type Worker struct {
	db    *sql.DB
	store *store.LocalStore
}

type Provider struct {
	Store         domain.Store
	TagRepo       domain.TagRepo
	WallpaperRepo domain.WallpaperRepo
}

func (w *Worker) Do(
	ctx app.Ctx,
	fn func(*Provider) error,
) error {
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	if err := fn(&Provider{
		Store:         w.store,
		TagRepo:       sqlite.NewTagRepo(tx),
		WallpaperRepo: sqlite.NewWallpaperRepo(tx),
	}); err != nil {
		return err
	}

	return tx.Commit()
}
