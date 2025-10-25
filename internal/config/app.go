package config

import (
	"context"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type App struct {
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
	a.handlers = NewHandlers(a.worker)
	return nil
}

func (a *App) Close() error {
	if a.worker == nil {
		return nil
	}

	err := a.worker.Close()

	a.worker = nil
	a.handlers = nil

	return err
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
