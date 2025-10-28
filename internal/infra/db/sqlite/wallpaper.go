package sqlite

import (
	"database/sql"
	_ "embed"

	"github.com/shimeoki/wp/internal/domain"
)

var (
	//go:embed sql/wallpaper_select.sql
	wallpaperSelectQuery string

	//go:embed sql/wallpaper_count.sql
	wallpaperCountQuery string

	//go:embed sql/wallpaper_delete.sql
	wallpaperDeleteQuery string

	//go:embed sql/wallpaper_create.sql
	wallpaperCreateQuery string

	//go:embed sql/wallpaper_update.sql
	wallpaperUpdateQuery string

	//go:embed sql/wallpaper_create_tag.sql
	wallpaperTagCreateQuery string

	//go:embed sql/wallpaper_delete_tag.sql
	wallpaperTagDeleteQuery string

	//go:embed sql/wallpaper_create_source.sql
	wallpaperSourceCreateQuery string

	//go:embed sql/wallpaper_delete.sql
	wallpaperSourceDeleteQuery string
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

func scanWallpaperTable(rows *sql.Rows) (*wallpaperTable, error) {
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
