package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *SourceRepo) Count(ctx domain.Ctx) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, sourceCountQuery).Scan(&count)
	return count, err
}
