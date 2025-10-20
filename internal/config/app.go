package config

import (
	"context"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type App struct {
	store    *store.LocalStore
	cfg      *Config
	handlers *Handlers
}

func NewApp(ctx context.Context, cfg *Config) (*App, error) {
	db, err := sqlite.Open(ctx, cfg.DB.DataSourceName)
	if err != nil {
		return nil, err
	}

	var hasher store.SHA256Hasher

	store, err := store.NewLocalStore(cfg.Store.Path, &hasher)
	if err != nil {
		return nil, err
	}

	p := &Provider{db: db, store: store}

	return &App{
		store:    store,
		cfg:      cfg,
		handlers: p.Handlers(),
	}, nil
}

func (a *App) Handlers() *Handlers {
	return a.handlers
}

func (a *App) Config() *Config {
	return a.cfg
}

func (a *App) Close() error {
	return a.store.Close()
}
