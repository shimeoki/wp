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

type Service interface {
	CreateWallpaper(Ctx, Image) (Hash, error)

	DeleteWallpaper(Ctx, Hash) error

	GetWallpaper(Ctx, Hash) (*Wallpaper, error)

	ShowWallpaper(Ctx, Hash) (Image, error)

	CreateTag(Ctx, Name) error

	DeleteTag(Ctx, Name) error

	RenameTag(ctx Ctx, before, after Name) error

	ListTags(Ctx) ([]Tag, error)

	AddTag(Ctx, Hash, Name) error

	RemoveTag(Ctx, Hash, Name) error

	CreateSource(Ctx, *SourceCreate) (ID, error)

	DeleteSource(Ctx, ID) error

	AddSource(Ctx, Hash, ID) error

	RemoveSource(Ctx, Hash, ID) error

	ListSources(Ctx) ([]Source, error)

	AddAlias(Ctx, Hash, Name) error

	RemoveAlias(Ctx, Hash, Name) error
}
