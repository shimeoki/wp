package sqlite

import (
	"iter"
	"slices"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *SourceRepo) All(ctx domain.Ctx) (iter.Seq[*domain.Source], error) {
	rows, err := r.db.QueryContext(ctx, sourceSelectQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var sources []*domain.Source

	for rows.Next() {
		var tbl sourceTable

		if err := rows.Scan(
			&tbl.ID,
			&tbl.UUID,
			&tbl.Name,
			&tbl.Link,
			&tbl.CreatedAt,
			&tbl.UpdatedAt,
		); err != nil {
			return nil, err
		}

		sources = append(sources, tbl.toDomain())
	}

	return slices.Values(sources), rows.Err()
}
