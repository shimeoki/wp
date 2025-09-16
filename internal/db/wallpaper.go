package db

import (
	"context"
	"database/sql"
	"errors"
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

type sqliteWallpaperScanner struct {
	rows             *sql.Rows
	wallpapers       map[int64]*Wallpaper
	tags             map[int64]*Tag
	sources          map[int64]*Source
	aliases          map[int64]bool
	wallpaperTags    map[WallpaperTag]bool
	wallpaperSources map[WallpaperSource]bool
}

func (r *sqliteWallpaperRepo) newScanner(
	rows *sql.Rows,
) *sqliteWallpaperScanner {
	return &sqliteWallpaperScanner{
		rows:             rows,
		wallpapers:       make(map[int64]*Wallpaper),
		tags:             make(map[int64]*Tag),
		sources:          make(map[int64]*Source),
		aliases:          make(map[int64]bool),
		wallpaperTags:    make(map[WallpaperTag]bool),
		wallpaperSources: make(map[WallpaperSource]bool),
	}
}

func (s *sqliteWallpaperScanner) scanRow(row *wallpaperRow) error {
	err := s.rows.Scan(
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

func (s *sqliteWallpaperScanner) scanRows() ([]*Wallpaper, error) {
	var ws []*Wallpaper

	for s.rows.Next() {
		var row wallpaperRow

		err := s.scanRow(&row)
		if err != nil {
			return nil, err
		}

		wid := row.wallpaper.ID

		if s.wallpapers[wid] == nil {
			s.wallpapers[wid] = &row.wallpaper
			ws = append(ws, &row.wallpaper)
		}

		s.scanAliases(wid, &row)
		s.scanTags(wid, &row)
		s.scanSources(wid, &row)
	}

	return ws, nil
}

func (s *sqliteWallpaperScanner) scanAliases(wid int64, row *wallpaperRow) {
	aid := row.alias.ID
	if aid == 0 || s.aliases[aid] {
		return
	}

	w := s.wallpapers[wid]

	s.aliases[aid] = true
	w.Aliases = append(w.Aliases, &row.alias)
}

func (s *sqliteWallpaperScanner) scanTags(wid int64, row *wallpaperRow) {
	tid := row.tag.ID
	if tid == 0 {
		return
	}

	w := s.wallpapers[wid]

	if s.tags[tid] == nil {
		s.tags[tid] = &row.tag
	}

	wt := WallpaperTag{WallpaperID: wid, TagID: tid}
	if !s.wallpaperTags[wt] {
		s.wallpaperTags[wt] = true
		w.Tags = append(w.Tags, s.tags[tid])
	}
}

func (s *sqliteWallpaperScanner) scanSources(wid int64, row *wallpaperRow) {
	sid := row.source.ID
	if sid == 0 {
		return
	}

	w := s.wallpapers[wid]

	if s.sources[sid] == nil {
		s.sources[sid] = &row.source
	}

	ws := WallpaperSource{WallpaperID: wid, SourceID: sid}
	if !s.wallpaperSources[ws] {
		s.wallpaperSources[ws] = true
		w.Sources = append(w.Sources, s.sources[sid])
	}
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
	s := r.newScanner(rows)

	ws, err := s.scanRows()
	if err != nil {
		return nil, err
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
	s := r.newScanner(rows)

	ws, err := s.scanRows()
	if err != nil {
		return nil, err
	}

	if len(ws) != 1 {
		return nil, errors.New("no or multiple wallpapers with provided id")
	}

	return ws[0], rows.Err()
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
	s := r.newScanner(rows)

	ws, err := s.scanRows()
	if err != nil {
		return nil, err
	}

	if len(ws) != 1 {
		return nil, errors.New("no or multiple wallpapers with provided hash")
	}

	return ws[0], rows.Err()
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
