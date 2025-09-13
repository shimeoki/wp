package db

import (
	"context"
	"database/sql"
	"time"
)

type Wallpaper struct {
	ID        int64  // primary key
	Hash      string // unique
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WallpaperAlias struct {
	WallpaperID int64 // foreign key
	AliasID     int64 // foreign key
}

type WallpaperTag struct {
	WallpaperID int64 // foreign key
	TagID       int64 // foreign key
}

type WallpaperSource struct {
	WallpaperID int64 // foreign key
	SourceID    int64 // foreign key
}

type WallpaperRepo interface {
	GetAll(ctx context.Context) ([]*Wallpaper, error)
	GetByID(ctx context.Context, id int64) (*Wallpaper, error)
	GetByHash(ctx context.Context, hash string) (*Wallpaper, error)

	Create(ctx context.Context, w *Wallpaper) error
	Update(ctx context.Context, w *Wallpaper) error
	Delete(ctx context.Context, id int64) error

	AddAlias(ctx context.Context, join *WallpaperAlias) error
	RemoveAlias(ctx context.Context, join *WallpaperAlias) error

	AddTag(ctx context.Context, join *WallpaperTag) error
	RemoveTag(ctx context.Context, join *WallpaperTag) error

	AddSource(ctx context.Context, join *WallpaperSource) error
	RemoveSource(ctx context.Context, join *WallpaperSource) error
}

type sqliteWallpaperRepo struct {
	db *sql.DB
}

func (r *sqliteWallpaperRepo) GetAll(
	ctx context.Context,
) ([]*Wallpaper, error) {
	sql := "select id, hash, extension, created_at, updated_at from wallpaper"

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ws []*Wallpaper

	for rows.Next() {
		var w Wallpaper

		err := rows.Scan(
			&w.ID,
			&w.Hash,
			&w.Extension,
			&w.CreatedAt,
			&w.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		ws = append(ws, &w)
	}

	return ws, rows.Err()
}

func (r *sqliteWallpaperRepo) GetByID(
	ctx context.Context,
	id int64,
) (*Wallpaper, error) {
	sql := `
		select
			id
			, hash
			, extension
			, created_at
			, updated_at
		from wallpaper where id = ?`

	var w Wallpaper
	err := r.db.QueryRowContext(ctx, sql, id).
		Scan(
			&w.ID,
			&w.Hash,
			&w.Extension,
			&w.CreatedAt,
			&w.UpdatedAt,
		)

	return &w, err
}

func (r *sqliteWallpaperRepo) GetByHash(
	ctx context.Context,
	hash string,
) (*Wallpaper, error) {
	sql := `
		select
			id
			, hash
			, extension
			, created_at
			, updated_at
		from wallpaper where hash = ?`

	var w Wallpaper
	err := r.db.QueryRowContext(ctx, sql, hash).
		Scan(
			&w.ID,
			&w.Hash,
			&w.Extension,
			&w.CreatedAt,
			&w.UpdatedAt,
		)

	return &w, err
}

func (r *sqliteWallpaperRepo) Create(ctx context.Context, w *Wallpaper) error {
	sql := "insert into wallpaper(hash, extension) values (?, ?)"

	result, err := r.db.ExecContext(ctx, sql, w.Hash, w.Extension)
	if err != nil {
		return err
	}

	w.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteWallpaperRepo) Update(ctx context.Context, w *Wallpaper) error {
	sql := "update wallpaper set hash = ?, extension = ? where id = ?"

	_, err := r.db.ExecContext(ctx, sql, w.Hash, w.Extension, w.ID)

	return err
}

func (r *sqliteWallpaperRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from wallpaper where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}

func (r *sqliteWallpaperRepo) AddAlias(
	ctx context.Context,
	join *WallpaperAlias,
) error {
	sql := `
		insert into wallpaper_alias(wallpaper_id, alias_id)
		values(?, ?)`

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.AliasID)

	return err
}

func (r *sqliteWallpaperRepo) RemoveAlias(
	ctx context.Context,
	join *WallpaperAlias,
) error {
	sql := "delete from wallpaper_alias where wallpaper_id = ? and alias_id = ?"

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.AliasID)

	return err
}

func (r *sqliteWallpaperRepo) AddTag(
	ctx context.Context,
	join *WallpaperTag,
) error {
	sql := `
		insert into wallpaper_tag(wallpaper_id, tag_id)
		values(?, ?)`

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.TagID)

	return err
}

func (r *sqliteWallpaperRepo) RemoveTag(
	ctx context.Context,
	join *WallpaperTag,
) error {
	sql := "delete from wallpaper_tag where wallpaper_id = ? and tag_id = ?"

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.TagID)

	return err
}

func (r *sqliteWallpaperRepo) AddSource(
	ctx context.Context,
	join *WallpaperSource,
) error {
	sql := `
		insert into wallpaper_source(wallpaper_id, source_id)
		values(?, ?)`

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.SourceID)

	return err
}

func (r *sqliteWallpaperRepo) RemoveSource(
	ctx context.Context,
	join *WallpaperSource,
) error {
	sql := "delete from wallpaper_source where wallpaper_id = ? and source_id = ?"

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.SourceID)

	return err
}
