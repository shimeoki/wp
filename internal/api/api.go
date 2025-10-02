package api

import (
	"github.com/shimeoki/wp/internal/db"
)

type Ctx = db.Ctx

type ID int64
type Name string

type Tag string
type Alias string

type Source struct {
	ID   `        json:"id"`
	Name `        json:"name"`
	Link *string `json:"link"`
}

type SourceCreate struct {
	Name `        json:"name"`
	Link *string `json:"link"`
}

type Wallpaper struct {
	Format  `         json:"format"`
	Aliases []Alias  `json:"aliases"`
	Tags    []Tag    `json:"tags"`
	Sources []Source `json:"sources"`
}

type Wallpapers map[Hash]*Wallpaper

// TODO: queue service

type TagService interface {
	Create(Ctx, Name) error
	Delete(Ctx, Name) error

	Rename(ctx Ctx, before, after Name) error

	List(Ctx) ([]Tag, error)
}

type SourceService interface {
	Create(Ctx, *SourceCreate) (ID, error)
	Delete(Ctx, ID) error

	// TODO: update

	List(Ctx) ([]Source, error)
}

type WallpaperService interface {
	Create(Ctx, Image) (Hash, error)
	Delete(Ctx, Hash) error
	Get(Ctx, Hash) (*Wallpaper, error)
	Show(Ctx, Hash) (Image, error)

	AddTag(Ctx, Hash, Name) error
	RemoveTag(Ctx, Hash, Name) error

	AddSource(Ctx, Hash, ID) error
	RemoveSource(Ctx, Hash, ID) error

	AddAlias(Ctx, Hash, Name) error
	RemoveAlias(Ctx, Hash, Name) error
}

type Service interface {
	Wallpapers() WallpaperService
	Sources() SourceService
	Tags() TagService
}
