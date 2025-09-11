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
	Create(ctx context.Context, a *Alias) error

	GetAll(ctx context.Context) ([]*Alias, error)
	GetByID(ctx context.Context, id int) (*Alias, error)

	Update(ctx context.Context, a *Alias) error

	Delete(ctx context.Context, id int) error
}
