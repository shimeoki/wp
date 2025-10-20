package config

import (
	"database/sql"

	"github.com/shimeoki/wp/internal/app"
	"github.com/shimeoki/wp/internal/domain"
	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type Provider struct {
	db    *sql.DB
	store *store.LocalStore
}

func (p *Provider) Provide(ctx app.Ctx) (*Worker, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Worker{
		tx:    tx,
		store: p.store,

		wallpapers: sqlite.NewWallpaperRepo(tx),
		tags:       sqlite.NewTagRepo(tx),
	}, nil
}

type Worker struct {
	tx    *sql.Tx
	store *store.LocalStore

	wallpapers *sqlite.WallpaperRepo
	tags       *sqlite.TagRepo
}

func (w *Worker) Commit() error {
	return w.tx.Commit()
}

func (w *Worker) Rollback() error {
	return w.tx.Rollback()
}

func (w *Worker) Store() domain.Store {
	return w.store
}

func (w *Worker) WallpaperRepo() domain.WallpaperRepo {
	return w.wallpapers
}

func (w *Worker) TagRepo() domain.TagRepo {
	return w.tags
}
