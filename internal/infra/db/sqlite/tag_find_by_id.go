package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *TagRepo) FindByID(ctx domain.Ctx, id domain.ID) (*domain.Tag, error) {
	var tbl tagTable

	row := r.db.QueryRowContext(ctx, tagSelectQuery+" where uuid = ?",
		id.String())

	if err := row.Scan(
		&tbl.ID,
		&tbl.UUID,
		&tbl.Name,
		&tbl.CreatedAt,
		&tbl.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return tbl.toDomain(), nil
}
