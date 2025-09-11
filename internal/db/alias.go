package db

import (
	"context"
	"time"
)

type Alias struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AliasRepo interface {
	GetAll(ctx context.Context) ([]*Alias, error)
	GetByID(ctx context.Context, ids ...int) ([]*Alias, error)

	Create(ctx context.Context, as ...*Alias) error
	Update(ctx context.Context, as ...*Alias) error
	Delete(ctx context.Context, ids ...int) error
}
