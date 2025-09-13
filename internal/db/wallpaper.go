package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Wallpaper struct {
	ID        int64
	Hash      string
	Extension string
	CreatedAt time.Time
	Aliases   []*Alias
	Tags      []*Tag
	Sources   []*Source
}

type WallpaperCreate struct {
	Hash      string
	Extension string
}

type WallpaperTag struct {
	WallpaperID int64
	TagID       int64
}

type WallpaperSource struct {
	WallpaperID int64
	SourceID    int64
}

type WallpaperRepo interface {
	GetAll(ctx context.Context) ([]*Wallpaper, error)
	GetByID(ctx context.Context, id int64) (*Wallpaper, error)
	GetByHash(ctx context.Context, hash string) (*Wallpaper, error)

	Create(ctx context.Context, w *WallpaperCreate) (int64, error)
	Delete(ctx context.Context, id int64) error

	AddTag(ctx context.Context, w *WallpaperTag) error
	RemoveTag(ctx context.Context, w *WallpaperTag) error

	AddSource(ctx context.Context, w *WallpaperSource) error
	RemoveSource(ctx context.Context, w *WallpaperSource) error
}

type sqliteWallpaperRepo struct {
	db *sql.DB
}

func (r *sqliteWallpaperRepo) query() string {
	return `
		select
			w.id
			, w.hash
			, w.extension
			, w.created_at
			, a.id
			, a.wallpaper_id
			, a.name
			, a.created_at
			, a.updated_at
			, t.id
			, t.name
			, t.created_at
			, t.updated_at
			, s.id
			, s.name
			, s.link
			, s.created_at
			, s.updated_at
		from wallpaper as w
		left join alias as a on a.wallpaper_id = w.id
		left join wallpaper_tag as wt on wt.wallpaper_id = w.id
		left join tag as t on wt.tag_id = t.id
		left join wallpaper_source as ws on ws.wallpaper_id = w.id
		left join source as s on ws.source_id = s.id`
}

type wallpaperRow struct {
	wallpaper Wallpaper
	alias     Alias
	tag       Tag
	source    Source
}

func (r *sqliteWallpaperRepo) scanRow(rows *sql.Rows, row *wallpaperRow) error {
	err := rows.Scan(
		row.wallpaper.ID,
		row.wallpaper.Hash,
		row.wallpaper.Extension,
		row.wallpaper.CreatedAt,
		row.alias.ID,
		row.alias.WallpaperID,
		row.alias.Name,
		row.alias.CreatedAt,
		row.alias.UpdatedAt,
		row.tag.ID,
		row.tag.Name,
		row.tag.CreatedAt,
		row.tag.UpdatedAt,
		row.source.ID,
		row.source.Name,
		row.source.Link,
		row.source.CreatedAt,
		row.source.UpdatedAt,
	)

	return err
}

func (r *sqliteWallpaperRepo) GetAll(
	ctx context.Context,
) ([]*Wallpaper, error) {
	sql := r.query()

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var ws []*Wallpaper

	wallpapers := make(map[int64]*Wallpaper)
	tags := make(map[int64]*Tag)
	wallpaperTags := make(map[WallpaperTag]bool)
	sources := make(map[int64]*Source)
	wallpaperSources := make(map[WallpaperSource]bool)
	aliases := make(map[int64]bool) // is many-to-one, above are many-to-many

	for rows.Next() {
		var row wallpaperRow

		err := r.scanRow(rows, &row)
		if err != nil {
			return nil, err
		}

		wid := row.wallpaper.ID
		if wallpapers[wid] == nil {
			wallpapers[wid] = &row.wallpaper
			ws = append(ws, &row.wallpaper)
		}

		w := wallpapers[wid]

		aid := row.alias.ID
		if aid != 0 && !aliases[aid] {
			aliases[aid] = true
			w.Aliases = append(w.Aliases, &row.alias)
		}

		tid := row.tag.ID
		if tid != 0 {
			if tags[tid] == nil {
				tags[tid] = &row.tag
			}

			wt := WallpaperTag{WallpaperID: wid, TagID: tid}
			if !wallpaperTags[wt] {
				wallpaperTags[wt] = true
				w.Tags = append(w.Tags, tags[tid])
			}
		}

		sid := row.source.ID
		if sid != 0 {
			if sources[sid] == nil {
				sources[sid] = &row.source
			}

			ws := WallpaperSource{WallpaperID: wid, SourceID: sid}
			if !wallpaperSources[ws] {
				wallpaperSources[ws] = true
				w.Sources = append(w.Sources, sources[sid])
			}
		}
	}

	return ws, rows.Err()
}

func (r *sqliteWallpaperRepo) GetByID(
	ctx context.Context,
	id int64,
) (*Wallpaper, error) {
	sql := fmt.Sprintf("%s where id = ?", r.query())

	rows, err := r.db.QueryContext(ctx, sql, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var w Wallpaper

	as := make(map[int64]bool)
	ts := make(map[int64]bool)
	ss := make(map[int64]bool)

	for rows.Next() {
		var row wallpaperRow

		err := r.scanRow(rows, &row)
		if err != nil {
			return nil, err
		}

		if w.ID == 0 {
			w = row.wallpaper
		}

		if row.alias.ID != 0 && !as[row.alias.ID] {
			as[row.alias.ID] = true
			w.Aliases = append(w.Aliases, &row.alias)
		}

		if row.tag.ID != 0 && !ts[row.tag.ID] {
			ts[row.tag.ID] = true
			w.Tags = append(w.Tags, &row.tag)
		}

		if row.source.ID != 0 && !ss[row.source.ID] {
			ss[row.source.ID] = true
			w.Sources = append(w.Sources, &row.source)
		}
	}

	return &w, rows.Err()
}

func (r *sqliteWallpaperRepo) GetByHash(
	ctx context.Context,
	hash string,
) (*Wallpaper, error) {
	sql := fmt.Sprintf("%s where hash = ?", r.query())

	rows, err := r.db.QueryContext(ctx, sql, hash)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var w Wallpaper

	as := make(map[int64]bool)
	ts := make(map[int64]bool)
	ss := make(map[int64]bool)

	for rows.Next() {
		var row wallpaperRow

		err := r.scanRow(rows, &row)
		if err != nil {
			return nil, err
		}

		if w.ID == 0 {
			w = row.wallpaper
		}

		if row.alias.ID != 0 && !as[row.alias.ID] {
			as[row.alias.ID] = true
			w.Aliases = append(w.Aliases, &row.alias)
		}

		if row.tag.ID != 0 && !ts[row.tag.ID] {
			ts[row.tag.ID] = true
			w.Tags = append(w.Tags, &row.tag)
		}

		if row.source.ID != 0 && !ss[row.source.ID] {
			ss[row.source.ID] = true
			w.Sources = append(w.Sources, &row.source)
		}
	}

	return &w, rows.Err()
}

func (r *sqliteWallpaperRepo) Create(
	ctx context.Context,
	w *WallpaperCreate,
) (int64, error) {
	sql := "insert into wallpaper(hash, extension) values (?, ?)"

	result, err := r.db.ExecContext(ctx, sql, w.Hash, w.Extension)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *sqliteWallpaperRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from wallpaper where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

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
	sql := `
		delete from wallpaper_source
		where wallpaper_id = ? and source_id = ?`

	_, err := r.db.ExecContext(ctx, sql, join.WallpaperID, join.SourceID)

	return err
}
