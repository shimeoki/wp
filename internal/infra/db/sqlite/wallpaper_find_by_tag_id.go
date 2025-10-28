package sqlite

import (
	"iter"
	"maps"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *WallpaperRepo) FindByTagID(
	ctx domain.Ctx,
	id domain.ID,
) (iter.Seq[*domain.Wallpaper], error) {
	rows, err := r.db.QueryContext(
		ctx,
		wallpaperSelectQuery+" where t.uuid = ?",
		id.String(),
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	return maps.Values(wallpapers), nil
}
