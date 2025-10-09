package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type ShowWallpaperHandler struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewShowWallpaperHandler(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *ShowWallpaperHandler {
	return &ShowWallpaperHandler{
		store:      store,
		wallpapers: wallpapers,
	}
}

type ShowWallpaperQuery struct {
	Hash string
}

type ShowWallpaperResult struct {
	Image  io.ReadCloser
	Format string
}

func (h *ShowWallpaperHandler) Handle(
	ctx Ctx,
	qry *ShowWallpaperQuery,
) (*ShowWallpaperResult, error) {
	hash := domain.Hash(qry.Hash)

	w, _ := h.wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	r, err := h.store.Get(hash)
	if err != nil {
		return nil, err
	}

	return &ShowWallpaperResult{Image: r, Format: w.Format.String()}, nil
}
