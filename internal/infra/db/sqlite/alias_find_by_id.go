package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *AliasRepo) FindByID(
	ctx domain.Ctx,
	id domain.ID,
) (*domain.Alias, error) {
	var tbl aliasTable

	row := r.db.QueryRowContext(ctx, aliasSelectQuery+" where uuid = ?",
		id.String())

	if err := row.Scan(
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

	return tbl.toDomain(), nil
}
