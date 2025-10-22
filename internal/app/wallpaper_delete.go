package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteWallpaperWorker struct {
	Store         domain.Store
	WallpaperRepo domain.WallpaperRepo
}

type DeleteWallpaperProvider Provider[*DeleteWallpaperWorker]

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
	var r DeleteWallpaperResult

	if err := h.provider(ctx, func(w *DeleteWallpaperWorker) error {
		hash, err := domain.ParseHash(cmd.Hash)
		if err != nil {
			return err
		}

		wall, _ := w.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		if err := w.WallpaperRepo.Delete(ctx, wall.ID); err != nil {
			return err
		}

		return w.Store.Remove(ctx, hash)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
