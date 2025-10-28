package sqlite

import (
	"iter"
	"maps"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *WallpaperRepo) All(
	ctx domain.Ctx,
) (iter.Seq[*domain.Wallpaper], error) {
	rows, err := r.db.QueryContext(ctx, wallpaperSelectQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	return maps.Values(wallpapers), rows.Err()
}
