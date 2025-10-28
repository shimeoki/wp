package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *SourceRepo) FindByID(
	ctx domain.Ctx,
	id domain.ID,
) (*domain.Source, error) {
	var tbl sourceTable

	row := r.db.QueryRowContext(ctx, sourceSelectQuery+" where uuid = ?",
		id.String())

	if err := row.Scan(
		&tbl.ID,
		&tbl.UUID,
		&tbl.Name,
		&tbl.Link,
		&tbl.CreatedAt,
		&tbl.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return tbl.toDomain(), nil
}
