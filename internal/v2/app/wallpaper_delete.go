package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type DeleteWallpaperHandler struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewDeleteWallpaperHandler(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *DeleteWallpaperHandler {
	return &DeleteWallpaperHandler{
		store:      store,
		wallpapers: wallpapers,
	}
}

type DeleteWallpaperCommand struct {
	Hash string
}

type DeleteWallpaperResult struct{}

func (h *DeleteWallpaperHandler) Handle(
	ctx Ctx,
	cmd *DeleteWallpaperCommand,
) (*DeleteWallpaperResult, error) {
	hash := domain.Hash(cmd.Hash)

	w, _ := h.wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	if err := h.wallpapers.Delete(ctx, w.ID); err != nil {
		return nil, err
	}

	if err := h.store.Remove(hash); err != nil {
		return nil, err
	}

	return &DeleteWallpaperResult{}, nil
}
