package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *SourceRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	_, err := r.db.ExecContext(ctx, sourceDeleteQuery, id.String())
	return err
}
