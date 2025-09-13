package db

import (
	"context"
	"database/sql"
	"time"
)

type Source struct {
	ID        int64 // primary key
	Name      string
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SourceRepo interface {
	GetAll(ctx context.Context) ([]*Source, error)
	GetByID(ctx context.Context, id int64) (*Source, error)

	Create(ctx context.Context, s *Source) error
	Update(ctx context.Context, s *Source) error
	Delete(ctx context.Context, id int64) error
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
	id int64,
) (*Source, error) {
	sql := `
		select
			id
			, name
			, link
			, created_at
			, updated_at
		from source where id = ?`

	var s Source
	err := r.db.QueryRowContext(ctx, sql, id).
		Scan(&s.ID, &s.Name, &s.Link, &s.CreatedAt, &s.UpdatedAt)

	return &s, err
}

func (r *sqliteSourceRepo) Create(ctx context.Context, s *Source) error {
	sql := "insert into source(name, link) values(?, ?)"

	result, err := r.db.ExecContext(ctx, sql, s.Name, s.Link)
	if err != nil {
		return err
	}

	s.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteSourceRepo) Update(ctx context.Context, s *Source) error {
	sql := "update source set name = ?, link = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, s.Name, s.Link, s.ID)

	return err
}

func (r *sqliteSourceRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from source where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
