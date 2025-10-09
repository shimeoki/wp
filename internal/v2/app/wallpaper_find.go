package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type FindWallpaperQuery struct {
	wallpapers domain.WallpaperRepo
}

func NewFindWallpaperQuery(
	wallpapers domain.WallpaperRepo,
) *FindWallpaperQuery {
	return &FindWallpaperQuery{
		wallpapers: wallpapers,
	}
}

type FindWallpaperData struct {
	Hash string
}

type FindWallpaperResult struct {
	Wallpaper WallpaperResult
}

func (qry *FindWallpaperQuery) Execute(
	ctx Ctx,
	data *FindWallpaperData,
) (*FindWallpaperResult, error) {
	w, _ := qry.wallpapers.ByHash(ctx, domain.Hash(data.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	f, err := toAppFormat(w.Format)
	if err != nil {
		return nil, err
	}

	wall := WallpaperResult{Format: f, Hash: data.Hash}

	return &FindWallpaperResult{Wallpaper: wall}, nil
}
