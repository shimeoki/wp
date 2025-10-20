package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteWallpaperWorker interface {
	Worker
	Store() domain.Store
	WallpaperRepo() domain.WallpaperRepo
}

type DeleteWallpaperProvider interface {
	Provider[DeleteWallpaperWorker]
}

type DeleteWallpaperHandler struct {
	provider DeleteWallpaperProvider
}

func NewDeleteWallpaperHandler(
	p DeleteWallpaperProvider,
) *DeleteWallpaperHandler {
	return &DeleteWallpaperHandler{provider: p}
}

type DeleteWallpaperCommand struct {
	Hash string
}

type DeleteWallpaperResult struct{}

func (h *DeleteWallpaperHandler) Handle(
	ctx Ctx,
	cmd *DeleteWallpaperCommand,
) (*DeleteWallpaperResult, error) {
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	defer worker.Rollback()
	store, repo := worker.Store(), worker.WallpaperRepo()

	hash, err := domain.ParseHash(cmd.Hash)
	if err != nil {
		return nil, err
	}

	w, _ := repo.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	if err := repo.Delete(ctx, w.ID); err != nil {
		return nil, err
	}

	if err := store.Remove(ctx, hash); err != nil {
		return nil, err
	}

	if err := worker.Commit(); err != nil {
		return nil, err
	}

	return &DeleteWallpaperResult{}, nil
}
