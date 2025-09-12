package db

import (
	"context"
	"database/sql"
	"time"
)

type Alias struct {
	ID        int64 // primary key
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AliasRepo interface {
	GetAll(ctx context.Context) ([]*Alias, error)
	GetByID(ctx context.Context, id int64) (*Alias, error)

	Create(ctx context.Context, a *Alias) error
	Update(ctx context.Context, a *Alias) error
	Delete(ctx context.Context, id int64) error
}

type sqliteAliasRepo struct {
	db *sql.DB
}

func (r *sqliteAliasRepo) GetAll(ctx context.Context) ([]*Alias, error) {
	sql := "select id, name, created_at, updated_at"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var as []*Alias

	for rows.Next() {
		var a Alias

		err := rows.Scan(&a.ID, &a.Name, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		as = append(as, &a)
	}

	return as, rows.Err()
}

func (r *sqliteAliasRepo) GetByID(
	ctx context.Context,
	id int64,
) (*Alias, error) {
	sql := "select id, name, created_at, updated_at from alias where id = ?"

	var a Alias
	err := r.db.QueryRowContext(ctx, sql, id).
		Scan(&a.ID, &a.Name, &a.CreatedAt, &a.UpdatedAt)

	return &a, err
}

func (r *sqliteAliasRepo) Create(ctx context.Context, a *Alias) error {
	sql := "insert into alias(name) values(?)"

	result, err := r.db.ExecContext(ctx, sql, a.Name)
	if err != nil {
		return err
	}

	a.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteAliasRepo) Update(ctx context.Context, a *Alias) error {
	sql := "update alias set name = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, a.Name, a.ID)

	return err
}

func (r *sqliteAliasRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from alias where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
