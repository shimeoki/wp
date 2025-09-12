package db

import (
	"context"
	"database/sql"
)

type Status struct {
	ID   int    // primary key
	Name string // unique
}

type StatusRepo interface {
	GetAll(ctx context.Context) ([]*Status, error)
	GetByID(ctx context.Context, id int) (*Status, error)
	GetByName(ctx context.Context, name string) (*Status, error)

	Create(ctx context.Context, s *Status) error
	Update(ctx context.Context, s *Status) error
	Delete(ctx context.Context, id int) error
}

type sqliteStatusRepo struct {
	db *sql.DB
}

func (r *sqliteStatusRepo) GetAll(ctx context.Context) ([]*Status, error) {
	sql := "select id, name from status"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ss []*Status

	for rows.Next() {
		var s Status

		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}

		ss = append(ss, &s)
	}

	return ss, rows.Err()
}

func (r *sqliteStatusRepo) GetByID(
	ctx context.Context,
	id int,
) (*Status, error) {
	sql := "select id, name from status where id = ?"

	var s Status
	err := r.db.QueryRowContext(ctx, sql, id).Scan(&s.ID, &s.Name)

	return &s, err
}

func (r *sqliteStatusRepo) GetByName(
	ctx context.Context,
	name string,
) (*Status, error) {
	sql := "select id, name from status where name = ?"

	var s Status
	err := r.db.QueryRowContext(ctx, sql, name).Scan(&s.ID, &s.Name)

	return &s, err
}
