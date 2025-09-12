package db

import (
	"context"
	"time"
)

type Source struct {
	ID        int64  // primary key
	Name      string // unique
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SourceRepo interface {
	GetAll(ctx context.Context) ([]*Source, error)
	GetByID(ctx context.Context, id int) (*Source, error)

	Create(ctx context.Context, s *Source) error
	Update(ctx context.Context, s *Source) error
	Delete(ctx context.Context, id int) error
}
