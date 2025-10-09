package app

import "errors"

type FindWallpaperQuery struct {
	Hash
}

type FindWallpaperResult struct {
	Wallpaper WallpaperResult
}

func (cmd *WallpaperService) Find(
	ctx Ctx,
	qry *FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, string(qry.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	f, err := toAppFormat(w.Format)
	if err != nil {
		return nil, err
	}

	wall := WallpaperResult{Format: f, Hash: qry.Hash}

	return &FindWallpaperResult{Wallpaper: wall}, nil
}
