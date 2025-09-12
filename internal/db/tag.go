package db

import (
	"context"
	"database/sql"
	"time"
)

type Tag struct {
	ID        int64  // primary key
	Name      string // unique
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
	GetAll(ctx context.Context) ([]*Tag, error)
	GetByID(ctx context.Context, id int64) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)

	Create(ctx context.Context, t *Tag) error
	Update(ctx context.Context, t *Tag) error
	Delete(ctx context.Context, id int64) error
}

type sqliteTagRepo struct {
	db *sql.DB
}

func (r *sqliteTagRepo) GetAll(ctx context.Context) ([]*Tag, error) {
	sql := "select id, name, created_at, updated_at from tag"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ts []*Tag

	for rows.Next() {
		var t Tag

		err := rows.Scan(&t.ID, &t.Name, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ts = append(ts, &t)
	}

	return ts, rows.Err()
}

func (r *sqliteTagRepo) GetByID(ctx context.Context, id int64) (*Tag, error) {
	sql := "select id, name, created_at, updated_at from tag where id = ?"

	var t Tag
	err := r.db.QueryRowContext(ctx, sql, id).
		Scan(&t.ID, &t.Name, &t.CreatedAt, &t.UpdatedAt)

	return &t, err
}

func (r *sqliteTagRepo) GetByName(
	ctx context.Context,
	name string,
) (*Tag, error) {
	sql := "select id, name, created_at, updated_at from tag where name = ?"

	var t Tag
	err := r.db.QueryRowContext(ctx, sql, name).
		Scan(&t.ID, &t.Name, &t.CreatedAt, &t.UpdatedAt)

	return &t, err
}

func (r *sqliteTagRepo) Create(ctx context.Context, t *Tag) error {
	sql := "insert into tag(name) values (?)"

	result, err := r.db.ExecContext(ctx, sql, t.Name)
	if err != nil {
		return err
	}

	t.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteTagRepo) Update(ctx context.Context, t *Tag) error {
	sql := "update tag set name = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, t.Name, t.ID)

	return err
}

func (r *sqliteTagRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from tag where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
