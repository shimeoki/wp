package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *TagRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	_, err := r.db.ExecContext(ctx, tagDeleteQuery, id.String())
	return err
}
