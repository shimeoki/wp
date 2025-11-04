package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *AliasRepo) Count(ctx domain.Ctx) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, aliasCountQuery).Scan(&count)
	return count, err
}
