package sqlite

import (
	"iter"
	"slices"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *AliasRepo) All(ctx domain.Ctx) (iter.Seq[*domain.Alias], error) {
	rows, err := r.db.QueryContext(ctx, aliasSelectQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var aliases []*domain.Alias

	for rows.Next() {
		var tbl aliasTable

		if err := rows.Scan(
			&tbl.ID,
			&tbl.UUID,
			&tbl.WallpaperID,
			&tbl.WallpaperUUID,
			&tbl.Name,
			&tbl.CreatedAt,
			&tbl.UpdatedAt,
		); err != nil {
			return nil, err
		}

		aliases = append(aliases, tbl.toDomain())
	}

	return slices.Values(aliases), rows.Err()
}
