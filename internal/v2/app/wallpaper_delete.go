package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type DeleteWallpaperCommand struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewDeleteWallpaperCommand(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *DeleteWallpaperCommand {
	return &DeleteWallpaperCommand{
		store:      store,
		wallpapers: wallpapers,
	}
}

type DeleteWallpaperData struct {
	Hash string
}

type DeleteWallpaperResult struct{}

func (cmd *DeleteWallpaperCommand) Execute(
	ctx Ctx,
	data *DeleteWallpaperData,
) (*DeleteWallpaperResult, error) {
	h := domain.Hash(data.Hash)

	w, _ := cmd.wallpapers.ByHash(ctx, h)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	if err := cmd.wallpapers.Delete(ctx, w.ID); err != nil {
		return nil, err
	}

	if err := cmd.store.Remove(h); err != nil {
		return nil, err
	}

	return &DeleteWallpaperResult{}, nil
}
