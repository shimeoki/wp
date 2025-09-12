package db

import (
	"context"
	"time"
)

type Queue struct {
	ID          int64 // primary key
	WallpaperID int64 // foreign key
	StatusID    int64 // foreign key
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueueRepo interface {
	GetAll(ctx context.Context) ([]*Queue, error)
	GetByID(ctx context.Context, id int64) (*Queue, error)

	Create(ctx context.Context, q *Queue) error
	Update(ctx context.Context, q *Queue) error
	Delete(ctx context.Context, id int64) error
}
