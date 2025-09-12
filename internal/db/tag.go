package db

import (
	"context"
	"time"
)

type Tag struct {
	ID        int64  // primary key
	Name      string // unique
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
	GetAll(ctx context.Context) ([]*Tag, error)
	GetByID(ctx context.Context, id int64) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)

	Create(ctx context.Context, t *Tag) error
	Update(ctx context.Context, t *Tag) error
	Delete(ctx context.Context, id int64) error
}
