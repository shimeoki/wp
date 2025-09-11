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
	GetByID(ctx context.Context, id int) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)

	Create(ctx context.Context, t *Tag) error
	Update(ctx context.Context, t *Tag) error
	Delete(ctx context.Context, id int) error
}
