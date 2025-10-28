package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *TagRepo) Count(ctx domain.Ctx) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, tagCountQuery).Scan(&count)
	return count, err
}
