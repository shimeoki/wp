package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type ShowWallpaperQuery struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewShowWallpaperQuery(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *ShowWallpaperQuery {
	return &ShowWallpaperQuery{
		store:      store,
		wallpapers: wallpapers,
	}
}

type ShowWallpaperData struct {
	Hash string
}

type ShowWallpaperResult struct {
	Image io.ReadCloser
	Format
}

func (qry *ShowWallpaperQuery) Execute(
	ctx Ctx,
	data *ShowWallpaperData,
) (*ShowWallpaperResult, error) {
	h := domain.Hash(data.Hash)

	w, _ := qry.wallpapers.ByHash(ctx, h)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	f, err := toAppFormat(w.Format)
	if err != nil {
		return nil, err
	}

	r, err := qry.store.Get(h)
	if err != nil {
		return nil, err
	}

	return &ShowWallpaperResult{Image: r, Format: f}, nil
}
