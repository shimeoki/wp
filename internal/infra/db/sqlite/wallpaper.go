package sqlite

import (
	"database/sql"
	"iter"
	"maps"

	"github.com/shimeoki/wp/internal/domain"
)

type WallpaperRepo struct {
	db DB
}

func NewWallpaperRepo(db DB) *WallpaperRepo {
	return &WallpaperRepo{db: db}
}

type wallpaperTable struct {
	WallpaperID        *integer
	WallpaperUUID      *uid
	WallpaperHash      *text
	WallpaperFormat    *text
	WallpaperCreatedAt *timestamp
	WallpaperUpdatedAt *timestamp

	TagID        *integer
	TagUUID      *uid
	TagName      *text
	TagCreatedAt *timestamp
	TagUpdatedAt *timestamp

	SourceID        *integer
	SourceUUID      *uid
	SourceName      *text
	SourceLink      *text
	SourceCreatedAt *timestamp
	SourceUpdatedAt *timestamp
}

func scanWallpaperTable(
	rows *sql.Rows,
) (*wallpaperTable, error) {
	var tbl wallpaperTable

	if err := rows.Scan(
		&tbl.WallpaperID,
		&tbl.WallpaperUUID,
		&tbl.WallpaperHash,
		&tbl.WallpaperFormat,
		&tbl.WallpaperCreatedAt,
		&tbl.WallpaperUpdatedAt,

		&tbl.TagID,
		&tbl.TagUUID,
		&tbl.TagName,
		&tbl.TagCreatedAt,
		&tbl.TagUpdatedAt,

		&tbl.SourceID,
		&tbl.SourceUUID,
		&tbl.SourceName,
		&tbl.SourceLink,
		&tbl.SourceCreatedAt,
		&tbl.SourceUpdatedAt,
	); err != nil {
		return nil, err
	}

	return &tbl, nil
}

func (t *wallpaperTable) toWallpaperDomain() (*domain.Wallpaper, error) {
	id, err := domain.ParseID(t.WallpaperUUID.String())
	if err != nil {
		return nil, err
	}

	f, err := domain.ParseFormat(*t.WallpaperFormat)
	if err != nil {
		return nil, err
	}

	h, err := domain.ParseHash(*t.WallpaperHash)
	if err != nil {
		return nil, err
	}

	return &domain.Wallpaper{
		ID:        id,
		Format:    f,
		Hash:      h,
		Sources:   make(map[domain.ID]*domain.Source),
		Tags:      make(map[domain.ID]*domain.Tag),
		CreatedAt: *t.WallpaperCreatedAt,
		UpdatedAt: *t.WallpaperUpdatedAt,
	}, nil
}

func (t *wallpaperTable) toSourceDomain() (*domain.Source, error) {
	id, err := domain.ParseID(t.SourceUUID.String())
	if err != nil {
		return nil, err
	}

	n, err := domain.ParseName(*t.SourceName)
	if err != nil {
		return nil, err
	}

	return &domain.Source{
		ID:        id,
		Name:      n,
		Link:      t.SourceLink,
		CreatedAt: *t.SourceCreatedAt,
		UpdatedAt: *t.SourceUpdatedAt,
	}, nil
}

func (t *wallpaperTable) toTagDomain() (*domain.Tag, error) {
	id, err := domain.ParseID(t.TagUUID.String())
	if err != nil {
		return nil, err
	}

	n, err := domain.ParseName(*t.TagName)
	if err != nil {
		return nil, err
	}

	return &domain.Tag{
		ID:        id,
		Name:      n,
		CreatedAt: *t.TagCreatedAt,
		UpdatedAt: *t.TagUpdatedAt,
	}, nil
}

var wallpaperQuery = `
	select
		w.id
		, w.uuid
		, w.hash
		, w.format
		, w.created_at
		, w.updated_at

		, t.id
		, t.uuid
		, t.name
		, t.created_at
		, t.updated_at

		, s.id
		, s.uuid
		, s.name
		, s.link
		, s.created_at
		, s.updated_at

	from wallpaper as w

	left join wallpaper_tag as wt on wt.wallpaper_id = w.id
	left join tag as t on wt.tag_id = t.id

	left join wallpaper_source as ws on ws.wallpaper_id = w.id
	left join source as s on ws.source_id = s.id
`

func scanWallpapers(rows *sql.Rows) (map[integer]*domain.Wallpaper, error) {
	wallpapers := make(map[integer]*domain.Wallpaper)
	sources := make(map[integer]*domain.Source)
	tags := make(map[integer]*domain.Tag)

	for rows.Next() {
		tbl, err := scanWallpaperTable(rows)
		if err != nil {
			return nil, err
		}

		wid := *tbl.WallpaperID

		w := wallpapers[wid]
		if w == nil {
			w, err := tbl.toWallpaperDomain()
			if err != nil {
				return nil, err
			}

			wallpapers[wid] = w
		}

		if tbl.SourceID != nil {
			sid := *tbl.SourceID

			s := sources[sid]
			if s == nil {
				s, err := tbl.toSourceDomain()
				if err != nil {
					return nil, err
				}

				sources[sid] = s
			}

			w.Sources[s.ID] = s
		}

		if tbl.TagID != nil {
			tid := *tbl.TagID

			t := tags[tid]
			if t == nil {
				t, err := tbl.toTagDomain()
				if err != nil {
					return nil, err
				}

				tags[tid] = t
			}

			w.Tags[t.ID] = t
		}
	}

	return wallpapers, rows.Err()
}

// keep-sorted start block=yes newline_separated=yes skip_lines=1

func (r *WallpaperRepo) All(
	ctx domain.Ctx,
) (iter.Seq[*domain.Wallpaper], error) {
	rows, err := r.db.QueryContext(ctx, wallpaperQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	return maps.Values(wallpapers), rows.Err()
}

func (r *WallpaperRepo) Count(ctx domain.Ctx) (int, error) {
	sql := `select count(*) from wallpaper`

	var count int
	err := r.db.QueryRowContext(ctx, sql).Scan(&count)

	return count, err
}

func (r *WallpaperRepo) Delete(ctx domain.Ctx, id domain.ID) error {
	sql := `delete from wallpaper where uuid = ?`

	_, err := r.db.ExecContext(ctx, sql, id.String())

	return err
}

func (r *WallpaperRepo) FindByID(
	ctx domain.Ctx,
	id domain.ID,
) (*domain.Wallpaper, error) {
	query := wallpaperQuery + ` where w.uuid = ?`

	rows, err := r.db.QueryContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	for _, w := range wallpapers {
		return w, nil
	}

	return nil, nil
}

func (r *WallpaperRepo) FindByHash(
	ctx domain.Ctx,
	hash domain.Hash,
) (*domain.Wallpaper, error) {
	query := wallpaperQuery + ` where w.hash = ?`

	rows, err := r.db.QueryContext(ctx, query, hash.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	for _, w := range wallpapers {
		return w, nil
	}

	return nil, nil
}

func (r *WallpaperRepo) FindByTagID(
	ctx domain.Ctx,
	id domain.ID,
) (iter.Seq[*domain.Wallpaper], error) {
	query := wallpaperQuery + ` where t.uuid = ?`

	rows, err := r.db.QueryContext(ctx, query, id.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	wallpapers, err := scanWallpapers(rows)
	if err != nil {
		return nil, err
	}

	return maps.Values(wallpapers), nil
}

func (r *WallpaperRepo) Save(ctx domain.Ctx, w *domain.Wallpaper) error {
	wallpaper, _ := r.FindByID(ctx, w.ID)
	if wallpaper == nil {
		return r.create(ctx, w)
	} else {
		return r.update(ctx, w)
	}
}

func (r *WallpaperRepo) create(
	ctx domain.Ctx,
	w *domain.Wallpaper,
) error {
	sql := `
		insert into wallpaper(
			uuid
			, format
			, hash
			, created_at
			, updated_at
		) values (?, ?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(
		ctx,
		sql,
		w.ID.String(),
		w.Format.String(),
		w.Hash.String(),
		w.CreatedAt,
		w.UpdatedAt,
	); err != nil {
		return err
	}

	return r.updateJoins(ctx, w)
}

func (r *WallpaperRepo) update(
	ctx domain.Ctx,
	w *domain.Wallpaper,
) error {
	sql := `
		update wallpaper set format = ?, hash = ?, updated_at = ? where uuid = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		sql,
		w.Format.String(),
		w.Hash.String(),
		w.UpdatedAt.String(),
		w.ID.String(),
	)

	return err
}

func (r *WallpaperRepo) updateJoins(
	ctx domain.Ctx,
	w *domain.Wallpaper,
) error {
	now, err := r.FindByID(ctx, w.ID)
	if err != nil {
		return err
	}

	wid := w.ID.String()

	if err := r.updateTags(ctx, wid, now.Tags, w.Tags); err != nil {
		return err
	}

	if err := r.updateSources(ctx, wid, now.Sources, w.Sources); err != nil {
		return err
	}

	return nil
}

func (r *WallpaperRepo) updateTags(
	ctx domain.Ctx,
	wid string,
	current, target map[domain.ID]*domain.Tag,
) error {
	var newTags []domain.ID

	for id := range target {
		if _, ok := current[id]; ok {
			delete(current, id)
		} else {
			newTags = append(newTags, id)
		}
	}

	for _, id := range newTags {
		sql := `
			insert into wallpaper_tag(wallpaper_id, tag_id) values (?, ?)
		`

		_, err := r.db.ExecContext(ctx, sql, wid, id.String())
		if err != nil {
			return err
		}
	}

	for id := range current {
		sql := `
			delete from wallpaper_tag where wallpaper_id = ? and tag_id = ?
		`

		_, err := r.db.ExecContext(ctx, sql, wid, id.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *WallpaperRepo) updateSources(
	ctx domain.Ctx,
	wid string,
	current, target map[domain.ID]*domain.Source,
) error {
	var newSources []domain.ID

	for id := range target {
		if _, ok := current[id]; ok {
			delete(current, id)
		} else {
			newSources = append(newSources, id)
		}
	}

	for _, id := range newSources {
		sql := `
			insert into wallpaper_source(wallpaper_id, source_id) values (?, ?)
		`

		_, err := r.db.ExecContext(ctx, sql, wid, id.String())
		if err != nil {
			return err
		}
	}

	for id := range current {
		sql := `
			delete from wallpaper_source where wallpaper_id = ? and source_id = ?
		`

		_, err := r.db.ExecContext(ctx, sql, wid, id.String())
		if err != nil {
			return err
		}
	}

	return nil
}

// keep-sorted end
