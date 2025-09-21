package db

import (
	"context"
	"database/sql"
	"time"
)

type Source struct {
	ID
	Name      string
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SourceCreate struct {
	Name string
	Link *string
}

type SourceUpdate struct {
	ID
	Name string
	Link *string
}

type SourceRepo interface {
	GetAll(ctx context.Context) ([]*Source, error)
	GetByID(ctx context.Context, id ID) (*Source, error)

	Create(ctx context.Context, s *SourceCreate) (ID, error)
	Update(ctx context.Context, s *SourceUpdate) error
	Delete(ctx context.Context, id ID) error
}

type sqliteSourceRepo struct {
	db *sql.DB
}

func (r *sqliteSourceRepo) GetAll(ctx context.Context) ([]*Source, error) {
	sql := "select id, name, link, created_at, updated_at from source"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ss []*Source

	for rows.Next() {
		var s Source

		err := rows.Scan(&s.ID, &s.Name, &s.Link, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ss = append(ss, &s)
	}

	return ss, rows.Err()
}

func (r *sqliteSourceRepo) GetByID(
	ctx context.Context,
	id ID,
) (*Source, error) {
	sql := `
		select
			id
			, name
			, link
			, created_at
			, updated_at
		from source where id = ?`

	row := r.db.QueryRowContext(ctx, sql, id)
	var s Source

	err := row.Scan(&s.ID, &s.Name, &s.Link, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *sqliteSourceRepo) Create(
	ctx context.Context,
	s *SourceCreate,
) (ID, error) {
	sql := "insert into source(name, link) values(?, ?)"

	result, err := r.db.ExecContext(ctx, sql, s.Name, s.Link)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID(id), nil
}

func (r *sqliteSourceRepo) Update(ctx context.Context, s *SourceUpdate) error {
	sql := "update source set name = ?, link = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, s.Name, s.Link, s.ID)

	return err
}

func (r *sqliteSourceRepo) Delete(ctx context.Context, id ID) error {
	sql := "delete from source where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
