package db

import (
	"context"
	"database/sql"
	"time"
)

type Alias struct {
	ID
	WallpaperID ID
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AliasCreate struct {
	WallpaperID ID
	Name        string
}

type AliasUpdate struct {
	ID
	Name string
}

type AliasRepo interface {
	GetAll(ctx context.Context) ([]*Alias, error)
	GetByID(ctx context.Context, id ID) (*Alias, error)
	GetByName(ctx context.Context, name string, wid ID) (*Alias, error)

	Create(ctx context.Context, a *AliasCreate) (ID, error)
	Update(ctx context.Context, a *AliasUpdate) error
	Delete(ctx context.Context, id ID) error
}

type sqliteAliasRepo struct {
	db *sql.DB
}

func (r *sqliteAliasRepo) GetAll(ctx context.Context) ([]*Alias, error) {
	sql := "select id, wallpaper_id, name, created_at, updated_at from alias"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var as []*Alias

	for rows.Next() {
		var a Alias

		err := rows.Scan(
			&a.ID,
			&a.WallpaperID,
			&a.Name,
			&a.CreatedAt,
			&a.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		as = append(as, &a)
	}

	return as, rows.Err()
}

func (r *sqliteAliasRepo) GetByID(ctx context.Context, id ID) (*Alias, error) {
	sql := `
		select
			id
			, wallpaper_id
			, name
			, created_at
			, updated_at
		from alias where id = ?`

	var a Alias
	row := r.db.QueryRowContext(ctx, sql, id)

	err := row.Scan(&a.ID, &a.WallpaperID, &a.Name, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *sqliteAliasRepo) GetByName(
	ctx context.Context,
	name string,
	wid ID,
) (*Alias, error) {
	sql := `
		select
			id
			, wallpaper_id
			, name
			, created_at
			, updated_at
		from alias where name = ? and wallpaper_id = ?`

	var a Alias
	row := r.db.QueryRowContext(ctx, sql, name, wid)

	err := row.Scan(&a.ID, &a.WallpaperID, &a.Name, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *sqliteAliasRepo) Create(
	ctx context.Context,
	a *AliasCreate,
) (ID, error) {
	sql := "insert into alias(name, wallpaper_id) values(?, ?)"

	result, err := r.db.ExecContext(ctx, sql, a.Name, a.WallpaperID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID(id), nil
}

func (r *sqliteAliasRepo) Update(ctx context.Context, a *AliasUpdate) error {
	sql := "update alias set name = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, a.Name, a.ID)

	return err
}

func (r *sqliteAliasRepo) Delete(ctx context.Context, id ID) error {
	sql := "delete from alias where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
