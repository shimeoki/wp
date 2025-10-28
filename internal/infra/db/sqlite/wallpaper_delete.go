package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *WallpaperRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	_, err := r.db.ExecContext(ctx, wallpaperDeleteQuery, id.String())
	return err
}
