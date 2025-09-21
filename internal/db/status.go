package db

import (
	"database/sql"
)

type Status struct {
	ID
	Name
}

type StatusCreate struct {
	Name
}

type StatusUpdate struct {
	ID
	Name
}

type StatusRepo interface {
	GetAll(ctx Ctx) ([]*Status, error)
	GetByID(ctx Ctx, id ID) (*Status, error)
	GetByName(ctx Ctx, name Name) (*Status, error)

	Create(ctx Ctx, s *StatusCreate) (ID, error)
	Update(ctx Ctx, s *StatusUpdate) error
	Delete(ctx Ctx, id ID) error
}

type sqliteStatusRepo struct {
	db *sql.DB
}

func (r *sqliteStatusRepo) GetAll(ctx Ctx) ([]*Status, error) {
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

func (r *sqliteStatusRepo) GetByID(ctx Ctx, id ID) (*Status, error) {
	sql := "select id, name from status where id = ?"

	row := r.db.QueryRowContext(ctx, sql, id)
	var s Status

	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *sqliteStatusRepo) GetByName(ctx Ctx, name Name) (*Status, error) {
	sql := "select id, name from status where name = ?"

	row := r.db.QueryRowContext(ctx, sql, name)
	var s Status

	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *sqliteStatusRepo) Create(ctx Ctx, s *StatusCreate) (ID, error) {
	sql := "insert into status(name) values (?)"

	result, err := r.db.ExecContext(ctx, sql, s.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID(id), nil
}

func (r *sqliteStatusRepo) Update(ctx Ctx, s *StatusUpdate) error {
	sql := "update status set name = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, s.Name, s.ID)

	return err
}

func (r *sqliteStatusRepo) Delete(ctx Ctx, id ID) error {
	sql := "delete from status where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
