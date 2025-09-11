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
	GetByID(ctx context.Context, id int) (*Wallpaper, error)
	GetByHash(ctx context.Context, hash string) (*Wallpaper, error)

	Create(ctx context.Context, w *Wallpaper) error
	Update(ctx context.Context, w *Wallpaper) error
	Delete(ctx context.Context, id int) error

	AddAlias(ctx context.Context, ids *WallpaperAlias) error
	RemoveAlias(ctx context.Context, ids *WallpaperAlias) error

	AddTag(ctx context.Context, ids *WallpaperTag) error
	RemoveTag(ctx context.Context, ids *WallpaperTag) error

	AddSource(ctx context.Context, ids *WallpaperSource) error
	RemoveSource(ctx context.Context, ids *WallpaperSource) error
}
