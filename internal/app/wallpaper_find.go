package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
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
	Format string
}

func (h *FindWallpaperHandler) Handle(
	ctx Ctx,
	qry *FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	hash, err := domain.ParseHash(qry.Hash)
	if err != nil {
		return nil, err
	}

	w, _ := h.wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	return &FindWallpaperResult{Format: w.Format.String()}, nil
}
