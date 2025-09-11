package db

import "time"

type Queue struct {
	ID          int
	WallpaperID int
	StatusID    int
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
