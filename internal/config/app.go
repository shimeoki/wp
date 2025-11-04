package config

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/shimeoki/wp/internal/infra/db/sqlite"
	"github.com/shimeoki/wp/internal/infra/store"
)

type App struct {
	cfg      *Config
	worker   *LocalSQLiteWorker
	handlers *Handlers

	logger *slog.Logger
	output io.WriteCloser
}

func NewApp(cfg *Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Open(ctx context.Context) error {
	db, err := sqlite.Open(ctx, a.cfg.DB.DataSourceName)
	if err != nil {
		return err
	}

	if a.cfg.Log.Path == "" {
		a.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		a.output, err = os.OpenFile(a.cfg.Log.Path,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			return err
		}

		a.logger = slog.New(slog.NewTextHandler(a.output, nil))
	}

	a.logger = a.logger.With("layer", "app")

	store, err := store.NewLocalStore(a.cfg.Store.Path, store.SHA256Hasher())
	if err != nil {
		return err
	}

	a.worker = &LocalSQLiteWorker{db: db, store: store}
	a.handlers = NewHandlers(a.worker, a.logger)
	return nil
}

func (a *App) Close() error {
	if a.worker == nil {
		return nil
	}

	if a.output != nil {
		a.output.Close()
	}

	err := a.worker.Close()

	a.logger = nil
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

func (a *App) Logger() *slog.Logger {
	return a.logger
}
