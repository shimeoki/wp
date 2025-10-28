package sqlite

import (
	"iter"
	"slices"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *TagRepo) All(ctx domain.Ctx) (iter.Seq[*domain.Tag], error) {
	rows, err := r.db.QueryContext(ctx, tagSelectQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tags []*domain.Tag

	for rows.Next() {
		var tbl tagTable

		if err := rows.Scan(
			&tbl.ID,
			&tbl.UUID,
			&tbl.Name,
			&tbl.CreatedAt,
			&tbl.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tags = append(tags, tbl.toDomain())
	}

	return slices.Values(tags), rows.Err()
}
