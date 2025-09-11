package db

import (
	"context"
	"time"
)

type Wallpaper struct {
	ID        int
	Hash      string
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WallpaperAlias struct {
	WallpaperID int
	AliasID     int
}

type WallpaperTag struct {
	WallpaperID int
	TagID       int
}

type WallpaperSource struct {
	WallpaperID int
	SourceID    int
}

type WallpaperRepo interface {
	GetAll(ctx context.Context) ([]*Wallpaper, error)
	GetByID(ctx context.Context, ids ...int) ([]*Wallpaper, error)
	GetByHash(ctx context.Context, hashes ...string) ([]*Wallpaper, error)

	Create(ctx context.Context, ws ...*Wallpaper) error
	Update(ctx context.Context, ws ...*Wallpaper) error
	Delete(ctx context.Context, ids ...int) error

	AddAlias(ctx context.Context, joins ...*WallpaperAlias) error
	RemoveAlias(ctx context.Context, joins ...*WallpaperAlias) error

	AddTag(ctx context.Context, joins ...*WallpaperTag) error
	RemoveTag(ctx context.Context, joins ...*WallpaperTag) error

	AddSource(ctx context.Context, joins ...*WallpaperSource) error
	RemoveSource(ctx context.Context, joins ...*WallpaperSource) error
}
