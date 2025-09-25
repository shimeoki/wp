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
	// Add a wallpaper entry with the specified image.
	CreateWallpaper(Ctx, Image) (Hash, error)

	// Remove a wallpaper entry with the image.
	DeleteWallpaper(Ctx, Hash) error

	// Get all information about the wallpaper.
	GetWallpaper(Ctx, Hash) (*Wallpaper, error)

	// Get the readable image for the specified wallpaper.
	ShowWallpaper(Ctx, Hash) (Image, error)

	// Add a tag. Can be added to some wallpapers later.
	CreateTag(Ctx, Name) error

	// Remove a tag. Also removes it from all wallpapers where it's used.
	DeleteTag(Ctx, Name) error

	// Rename a tag "before" to "after", if "before" exists.
	RenameTag(ctx Ctx, before, after Name) error

	// Get a list of all tags.
	ListTags(Ctx) ([]Tag, error)

	// Attach a tag to a specified wallpaper.
	AddTag(Ctx, Hash, Name) error

	// Unattach a tag from a specified wallpaper.
	RemoveTag(Ctx, Hash, Name) error

	// Add a source. Can be added to some wallpapers later.
	CreateSource(Ctx, *SourceCreate) (ID, error)

	// Remove a source. Also removes it from all wallpapers where it's used.
	DeleteSource(Ctx, ID) error

	// Attach a source to specified wallpaper.
	AddSource(Ctx, Hash, ID) error

	// Unattach a source from a specified wallpaper.
	RemoveSource(Ctx, Hash, ID) error

	// Get a list of all sources.
	ListSources(Ctx) ([]Source, error)

	// Add an alias to a wallpaper.
	AddAlias(Ctx, Hash, Name) error

	// Remove an alias from a wallpaper.
	RemoveAlias(Ctx, Hash, Name) error
}
