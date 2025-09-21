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
	Format    string
	CreatedAt time.Time
	Aliases   []*Alias
	Tags      []*Tag
	Sources   []*Source
}

type WallpaperCreate struct {
	Hash   string
	Format string
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
			, w.format
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

	aliasID          *int64
	aliasWallpaperID *int64
	aliasName        *string
	aliasCreatedAt   *time.Time
	aliasUpdatedAt   *time.Time

	tagID        *int64
	tagName      *string
	tagCreatedAt *time.Time
	tagUpdatedAt *time.Time

	sourceID        *int64
	sourceName      *string
	sourceLink      *string
	sourceCreatedAt *time.Time
	sourceUpdatedAt *time.Time
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
		// wallpaper fields are guaranteed to be not null
		&row.wallpaper.ID,
		&row.wallpaper.Hash,
		&row.wallpaper.Format,
		&row.wallpaper.CreatedAt,

		&row.aliasID,
		&row.aliasWallpaperID,
		&row.aliasName,
		&row.aliasCreatedAt,
		&row.aliasUpdatedAt,

		&row.tagID,
		&row.tagName,
		&row.tagCreatedAt,
		&row.tagUpdatedAt,

		&row.sourceID,
		&row.sourceName,
		&row.sourceLink,
		&row.sourceCreatedAt,
		&row.sourceUpdatedAt,
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

		s.scanAlias(wid, &row)
		s.scanTag(wid, &row)
		s.scanSource(wid, &row)
	}

	return ws, nil
}

func (s *sqliteWallpaperScanner) scanAlias(wid int64, row *wallpaperRow) {
	if row.aliasID == nil {
		return
	}

	aid := *row.aliasID
	if aid == 0 || s.aliases[aid] {
		return
	}

	w := s.wallpapers[wid]
	alias := &Alias{
		ID:          ID(aid),
		WallpaperID: ID(*row.aliasWallpaperID),
		Name:        *row.aliasName,
		CreatedAt:   *row.aliasCreatedAt,
		UpdatedAt:   *row.aliasUpdatedAt,
	}

	s.aliases[aid] = true
	w.Aliases = append(w.Aliases, alias)
}

func (s *sqliteWallpaperScanner) scanTag(wid int64, row *wallpaperRow) {
	if row.tagID == nil {
		return
	}

	tid := *row.tagID
	if tid == 0 {
		return
	}

	w := s.wallpapers[wid]
	tag := &Tag{
		ID:        tid,
		Name:      *row.tagName,
		CreatedAt: *row.tagCreatedAt,
		UpdatedAt: *row.tagUpdatedAt,
	}

	if s.tags[tid] == nil {
		s.tags[tid] = tag
	}

	wt := WallpaperTag{WallpaperID: wid, TagID: tid}
	if !s.wallpaperTags[wt] {
		s.wallpaperTags[wt] = true
		w.Tags = append(w.Tags, s.tags[tid])
	}
}

func (s *sqliteWallpaperScanner) scanSource(wid int64, row *wallpaperRow) {
	if row.sourceID == nil {
		return
	}

	sid := *row.sourceID
	if sid == 0 {
		return
	}

	w := s.wallpapers[wid]
	source := &Source{
		ID:        sid,
		Name:      *row.sourceName,
		Link:      row.sourceLink,
		CreatedAt: *row.sourceCreatedAt,
		UpdatedAt: *row.sourceUpdatedAt,
	}

	if s.sources[sid] == nil {
		s.sources[sid] = source
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
	sql := fmt.Sprintf("%s where w.id = ?", r.query())

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

	switch len(ws) {
	case 0:
		return nil, rows.Err()
	case 1:
		return ws[0], rows.Err()
	default:
		return nil, errors.New("no or multiple wallpapers with provided id")
	}
}

func (r *sqliteWallpaperRepo) GetByHash(
	ctx context.Context,
	hash string,
) (*Wallpaper, error) {
	sql := fmt.Sprintf("%s where w.hash = ?", r.query())

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

	switch len(ws) {
	case 0:
		return nil, rows.Err()
	case 1:
		return ws[0], rows.Err()
	default:
		return nil, errors.New("no or multiple wallpapers with provided hash")
	}
}

func (r *sqliteWallpaperRepo) Create(
	ctx context.Context,
	w *WallpaperCreate,
) (int64, error) {
	sql := "insert into wallpaper(hash, format) values (?, ?)"

	result, err := r.db.ExecContext(ctx, sql, w.Hash, w.Format)
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
