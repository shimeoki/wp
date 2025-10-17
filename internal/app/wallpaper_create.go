package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateWallpaperHandler struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewCreateWallpaperHandler(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *CreateWallpaperHandler {
	return &CreateWallpaperHandler{
		store:      store,
		wallpapers: wallpapers,
	}
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
	hash, err := h.store.Create(ctx, cmd.Image)
	if err != nil {
		return nil, err
	}

	w, _ := h.wallpapers.FindByHash(ctx, hash)
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

	if err := h.wallpapers.Save(ctx, wall); err != nil {
		return nil, err
	}

	return &CreateWallpaperResult{Hash: string(hash)}, nil
}
