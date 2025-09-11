package db

import (
	"context"
	"time"
)

type Source struct {
	ID        int
	Name      string
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SourceRepo interface {
	GetAll(ctx context.Context) ([]*Source, error)
	GetByID(ctx context.Context, ids ...int) ([]*Source, error)

	Create(ctx context.Context, ss ...*Source) error
	Update(ctx context.Context, ss ...*Source) error
	Delete(ctx context.Context, ids ...int) error
}
