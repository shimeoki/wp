package db

import (
	"context"
	"database/sql"
	"time"
)

type Tag struct {
	ID
	Name
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagCreate struct {
	Name
}

type TagUpdate struct {
	ID
	Name
}

type TagRepo interface {
	GetAll(ctx context.Context) ([]*Tag, error)
	GetByID(ctx context.Context, id ID) (*Tag, error)
	GetByName(ctx context.Context, name Name) (*Tag, error)

	Create(ctx context.Context, t *TagCreate) (ID, error)
	Update(ctx context.Context, t *TagUpdate) error
	Delete(ctx context.Context, id ID) error
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

func (r *sqliteTagRepo) GetByID(ctx context.Context, id ID) (*Tag, error) {
	sql := "select id, name, created_at, updated_at from tag where id = ?"

	row := r.db.QueryRowContext(ctx, sql, id)
	var t Tag

	if err := row.Scan(&t.ID, &t.Name, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *sqliteTagRepo) GetByName(
	ctx context.Context,
	name Name,
) (*Tag, error) {
	sql := "select id, name, created_at, updated_at from tag where name = ?"

	row := r.db.QueryRowContext(ctx, sql, name)
	var t Tag

	if err := row.Scan(&t.ID, &t.Name, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *sqliteTagRepo) Create(
	ctx context.Context,
	t *TagCreate,
) (ID, error) {
	sql := "insert into tag(name) values (?)"

	result, err := r.db.ExecContext(ctx, sql, t.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID(id), nil
}

func (r *sqliteTagRepo) Update(ctx context.Context, t *TagUpdate) error {
	sql := "update tag set name = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, t.Name, t.ID)

	return err
}

func (r *sqliteTagRepo) Delete(ctx context.Context, id ID) error {
	sql := "delete from tag where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
