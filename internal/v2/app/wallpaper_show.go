package app

import "errors"

type ShowWallpaperQuery struct {
	Hash
}

type ShowWallpaperResult struct {
	Image
}

func (cmd *WallpaperService) Show(
	ctx Ctx,
	qry *ShowWallpaperQuery,
) (*ShowWallpaperResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, string(qry.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	f, err := toAppFormat(w.Format)
	if err != nil {
		return nil, err
	}

	r, err := cmd.store.Get(qry.Hash)
	if err != nil {
		return nil, err
	}

	img := &image{ReadCloser: r, format: f}

	return &ShowWallpaperResult{Image: img}, nil
}
