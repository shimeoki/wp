package db

import (
	"context"
	"time"
)

type Wallpaper struct {
	ID        int64  // primary key
	Hash      string // unique
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WallpaperAlias struct {
	WallpaperID int64 // foreign key
	AliasID     int64 // foreign key
}

type WallpaperTag struct {
	WallpaperID int64 // foreign key
	TagID       int64 // foreign key
}

type WallpaperSource struct {
	WallpaperID int64 // foreign key
	SourceID    int64 // foreign key
}

type WallpaperRepo interface {
	GetAll(ctx context.Context) ([]*Wallpaper, error)
	GetByID(ctx context.Context, id int64) (*Wallpaper, error)
	GetByHash(ctx context.Context, hash string) (*Wallpaper, error)

	Create(ctx context.Context, w *Wallpaper) error
	Update(ctx context.Context, w *Wallpaper) error
	Delete(ctx context.Context, id int64) error

	AddAlias(ctx context.Context, join *WallpaperAlias) error
	RemoveAlias(ctx context.Context, join *WallpaperAlias) error

	AddTag(ctx context.Context, join *WallpaperTag) error
	RemoveTag(ctx context.Context, join *WallpaperTag) error

	AddSource(ctx context.Context, join *WallpaperSource) error
	RemoveSource(ctx context.Context, join *WallpaperSource) error
}
