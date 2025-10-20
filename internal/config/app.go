package config

import (
	"context"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type App struct {
	store    *store.LocalStore
	cfg      *Config
	provider *Provider
	handlers *Handlers
}

func NewApp(cfg *Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Open(ctx context.Context) error {
	db, err := sqlite.Open(ctx, a.cfg.DB.DataSourceName)
	if err != nil {
		return err
	}

	var hasher store.SHA256Hasher

	store, err := store.NewLocalStore(a.cfg.Store.Path, &hasher)
	if err != nil {
		return err
	}

	a.provider = &Provider{db: db, store: store}
	a.handlers = a.provider.Handlers()
	return nil
}

func (a *App) Close() error {
	return a.store.Close()
}

func (a *App) Handlers() *Handlers {
	if a.provider == nil {
		return nil
	}

	return a.handlers
}

func (a *App) Config() *Config {
	return a.cfg
}
