package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type FindWallpaperHandler struct {
	wallpapers domain.WallpaperRepo
}

func NewFindWallpaperHandler(
	wallpapers domain.WallpaperRepo,
) *FindWallpaperHandler {
	return &FindWallpaperHandler{
		wallpapers: wallpapers,
	}
}

type FindWallpaperQuery struct {
	Hash string
}

type FindWallpaperResult struct {
	Wallpaper WallpaperResult
}

func (h *FindWallpaperHandler) Handle(
	ctx Ctx,
	qry *FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	w, _ := h.wallpapers.ByHash(ctx, domain.Hash(qry.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	wall := WallpaperResult{Format: w.Format.String(), Hash: qry.Hash}

	return &FindWallpaperResult{Wallpaper: wall}, nil
}
