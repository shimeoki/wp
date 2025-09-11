package db

import "time"

type Tag struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Source struct {
	ID        int
	Name      string
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

type Queue struct {
	ID          int
	WallpaperID int
	StatusID    int
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
