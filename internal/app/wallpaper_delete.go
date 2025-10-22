package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteWallpaperProviders struct {
	Store         domain.Store
	WallpaperRepo domain.WallpaperRepo
}

type DeleteWallpaperWorker Worker[*DeleteWallpaperProviders]

type DeleteWallpaperHandler struct {
	worker DeleteWallpaperWorker
}

func NewDeleteWallpaperHandler(
	w DeleteWallpaperWorker,
) *DeleteWallpaperHandler {
	return &DeleteWallpaperHandler{worker: w}
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

	if err := h.worker.Do(ctx, func(p *DeleteWallpaperProviders) error {
		hash, err := domain.ParseHash(cmd.Hash)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		if err := p.WallpaperRepo.Delete(ctx, wall.ID); err != nil {
			return err
		}

		return p.Store.Remove(ctx, hash)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
