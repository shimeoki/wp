package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type ShowWallpaperQuery struct {
	store      Store
	wallpapers domain.WallpaperRepo
}

func NewShowWallpaperQuery(
	store Store,
	wallpapers domain.WallpaperRepo,
) *ShowWallpaperQuery {
	return &ShowWallpaperQuery{
		store:      store,
		wallpapers: wallpapers,
	}
}

type ShowWallpaperData struct {
	Hash
}

type ShowWallpaperResult struct {
	Image
}

func (qry *ShowWallpaperQuery) Execute(
	ctx Ctx,
	data *ShowWallpaperData,
) (*ShowWallpaperResult, error) {
	w, _ := qry.wallpapers.ByHash(ctx, string(data.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	f, err := toAppFormat(w.Format)
	if err != nil {
		return nil, err
	}

	r, err := qry.store.Get(data.Hash)
	if err != nil {
		return nil, err
	}

	img := &image{ReadCloser: r, format: f}

	return &ShowWallpaperResult{Image: img}, nil
}
