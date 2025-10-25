package config

import (
	"context"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type App struct {
	store    *store.LocalStore
	cfg      *Config
	worker   *LocalSQLiteWorker
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

	store, err := store.NewLocalStore(a.cfg.Store.Path, store.SHA256Hasher())
	if err != nil {
		return err
	}

	a.worker = &LocalSQLiteWorker{db: db, store: store}
	a.handlers = a.worker.Handlers()
	return nil
}

func (a *App) Close() error {
	if a.store == nil {
		return nil
	}

	return a.store.Close()
}

func (a *App) Handlers() *Handlers {
	if a.worker == nil {
		return nil
	}

	return a.handlers
}

func (a *App) Config() *Config {
	return a.cfg
}
