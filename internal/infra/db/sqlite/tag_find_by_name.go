package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *TagRepo) FindByName(
	ctx domain.Ctx,
	n domain.Name,
) (*domain.Tag, error) {
	var tbl tagTable

	row := r.db.QueryRowContext(ctx, tagSelectQuery+" where name = ?",
		n.String())

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
