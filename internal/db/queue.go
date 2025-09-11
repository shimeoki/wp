package db

import (
	"context"
	"time"
)

type Queue struct {
	ID          int
	WallpaperID int
	StatusID    int
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueueRepo interface {
	GetAll(ctx context.Context) ([]*Queue, error)
	GetByID(ctx context.Context, ids ...int) ([]*Queue, error)

	Create(ctx context.Context, qs ...*Queue) error
	Update(ctx context.Context, qs ...*Queue) error
	Delete(ctx context.Context, ids ...int) error
}
