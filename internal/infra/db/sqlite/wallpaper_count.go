package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *WallpaperRepo) Count(ctx domain.Ctx) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, wallpaperCountQuery).Scan(&count)
	return count, err
}
