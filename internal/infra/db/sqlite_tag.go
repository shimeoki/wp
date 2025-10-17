package db

import (
	"iter"
	"slices"

	"github.com/shimeoki/wp/internal/domain"
)

type sqliteTagTable struct {
	ID        integer
	UUID      uid
	Name      text
	CreatedAt timestamp
	UpdatedAt timestamp
}

func (t *sqliteTagTable) toDomain() *domain.Tag {
	return &domain.Tag{
		ID:        domain.ID(t.UUID),
		Name:      domain.Name(t.Name),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

type SQLiteTagRepo struct {
	db DB
}

func NewSQLiteTagRepo(db DB) *SQLiteTagRepo {
	return &SQLiteTagRepo{db: db}
}

// keep-sorted start block=yes newline_separated=yes skip_lines=1

func (r *SQLiteTagRepo) All(ctx domain.Ctx) (iter.Seq[*domain.Tag], error) {
	sql := `select id, uuid, name, created_at, updated_at from tag`

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var tags []*domain.Tag

	for rows.Next() {
		var tbl sqliteTagTable

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

func (r *SQLiteTagRepo) Count(ctx domain.Ctx) (int, error) {
	sql := `select count(*) from tag`

	var count int
	err := r.db.QueryRowContext(ctx, sql).Scan(&count)

	return count, err
}

func (r *SQLiteTagRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	sql := `delete from tag where uuid = ?`

	_, err := r.db.ExecContext(ctx, sql, id.String())

	return err
}

func (r *SQLiteTagRepo) FindByID(
	ctx domain.Ctx,
	id domain.ID,
) (*domain.Tag, error) {
	sql := `
		select id, uuid, name, created_at, updated_at from tag where uuid = ?
	`

	row := r.db.QueryRowContext(ctx, sql, id.String())
	var tbl sqliteTagTable

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

func (r *SQLiteTagRepo) FindByName(
	ctx domain.Ctx,
	n domain.Name,
) (*domain.Tag, error) {
	sql := `
		select id, uuid, name, created_at, updated_at from tag where name = ?
	`

	row := r.db.QueryRowContext(ctx, sql, n.String())
	var tbl sqliteTagTable

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

func (r *SQLiteTagRepo) Save(ctx domain.Ctx, t *domain.Tag) error {
	tag, _ := r.FindByID(ctx, t.ID)
	if tag == nil {
		return r.create(ctx, t)
	} else {
		return r.update(ctx, t)
	}
}

func (r *SQLiteTagRepo) create(ctx domain.Ctx, t *domain.Tag) error {
	sql := `
		insert into tag(uuid, name, created_at, updated_at) values (?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(
		ctx,
		sql,
		t.ID.String(),
		t.Name.String(),
		t.CreatedAt,
		t.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (r *SQLiteTagRepo) update(ctx domain.Ctx, t *domain.Tag) error {
	sql := `update tag set name = ?, updated_at = ? where uuid = ?`

	_, err := r.db.ExecContext(
		ctx,
		sql,
		t.Name.String(),
		t.UpdatedAt,
		t.ID.String(),
	)

	return err
}

// keep-sorted end
