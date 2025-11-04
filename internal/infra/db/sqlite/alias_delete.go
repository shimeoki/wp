package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *AliasRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	_, err := r.db.ExecContext(ctx, aliasDeleteQuery, id.String())
	return err
}
