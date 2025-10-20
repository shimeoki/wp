package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateWallpaperWorker interface {
	Worker
	Store() domain.Store
	WallpaperRepo() domain.WallpaperRepo
}

type CreateWallpaperProvider interface {
	Provider[CreateWallpaperWorker]
}

type CreateWallpaperHandler struct {
	provider CreateWallpaperProvider
}

func NewCreateWallpaperHandler(
	p CreateWallpaperProvider,
) *CreateWallpaperHandler {
	return &CreateWallpaperHandler{provider: p}
}

type CreateWallpaperCommand struct {
	Image  io.ReadCloser
	Format string
}

type CreateWallpaperResult struct {
	Hash string
}

func (h *CreateWallpaperHandler) Handle(
	ctx Ctx,
	cmd *CreateWallpaperCommand,
) (*CreateWallpaperResult, error) {
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	defer worker.Rollback()
	store, repo := worker.Store(), worker.WallpaperRepo()

	hash, err := store.Create(ctx, cmd.Image)
	if err != nil {
		return nil, err
	}

	w, _ := repo.FindByHash(ctx, hash)
	if w != nil {
		return nil, errors.New("wallpaper already exists")
	}

	f, err := domain.ParseFormat(cmd.Format)
	if err != nil {
		return nil, err
	}

	wall, err := domain.NewWallpaper(f, hash)
	if err != nil {
		return nil, err
	}

	if err := repo.Save(ctx, wall); err != nil {
		return nil, err
	}

	if err := worker.Commit(); err != nil {
		return nil, err
	}

	return &CreateWallpaperResult{Hash: hash.String()}, nil
}
