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
		store:      w.store,
		tags:       sqlite.NewTagRepo(tx),
		wallpapers: sqlite.NewWallpaperRepo(tx),
	}); err != nil {
		return err
	}

	return tx.Commit()
}

type Provider struct {
	store      domain.Store
	tags       domain.TagRepo
	wallpapers domain.WallpaperRepo
}

func (p *Provider) Store() domain.Store {
	return p.store
}

func (p *Provider) TagRepo() domain.TagRepo {
	return p.tags
}

func (p *Provider) WallpaperRepo() domain.WallpaperRepo {
	return p.wallpapers
}
