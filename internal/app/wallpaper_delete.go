package app

import "github.com/shimeoki/wp/internal/domain"

type DeleteWallpaperProvider interface {
	StoreProvider
	WallpaperProvider
}

type DeleteWallpaperWorker Worker[DeleteWallpaperProvider]

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

	if err := h.worker.Work(ctx, func(p DeleteWallpaperProvider) error {
		hash, err := domain.ParseHash(cmd.Hash)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo().FindByHash(ctx, hash)
		if wall == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		if err := p.WallpaperRepo().Delete(ctx, wall.ID); err != nil {
			return err
		}

		return p.Store().Delete(ctx, hash)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
