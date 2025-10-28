package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *WallpaperRepo) FindByID(
	ctx domain.Ctx,
	id domain.ID,
) (*domain.Wallpaper, error) {
	rows, err := r.db.QueryContext(
		ctx,
		wallpaperSelectQuery+" where w.uuid = ?",
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

	for _, w := range wallpapers {
		return w, nil
	}

	return nil, nil
}
