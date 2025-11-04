package sqlite

import (
	"iter"
	"slices"

	"github.com/shimeoki/wp/internal/domain"
)

func (r *AliasRepo) FindByWallpaperID(
	ctx domain.Ctx,
	id domain.ID,
) (iter.Seq[*domain.Alias], error) {
	rows, err := r.db.QueryContext(ctx, aliasSelectQuery+" where w.uuid = ?",
		id.String())

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
