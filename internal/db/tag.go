package db

import (
	"context"
	"time"
)

type Tag struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TagRepo interface {
	GetAll(ctx context.Context) ([]*Tag, error)
	GetByID(ctx context.Context, ids ...int) ([]*Tag, error)
	GetByName(ctx context.Context, names ...string) ([]*Tag, error)

	Create(ctx context.Context, ts ...*Tag) error
	Update(ctx context.Context, ts ...*Tag) error
	Delete(ctx context.Context, ids ...int) error
}
