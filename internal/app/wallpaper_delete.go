package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
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
	hash, err := domain.ParseHash(cmd.Hash)
	if err != nil {
		return nil, err
	}

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
