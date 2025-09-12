package db

import (
	"context"
	"time"
)

type Alias struct {
	ID        int64 // primary key
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AliasRepo interface {
	GetAll(ctx context.Context) ([]*Alias, error)
	GetByID(ctx context.Context, id int64) (*Alias, error)

	Create(ctx context.Context, a *Alias) error
	Update(ctx context.Context, a *Alias) error
	Delete(ctx context.Context, id int64) error
}
