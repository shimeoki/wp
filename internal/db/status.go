package db

import "context"

type Status struct {
	ID   int
	Name string
}

type StatusRepo interface {
	GetAll(ctx context.Context) ([]*Status, error)
	GetByID(ctx context.Context, ids ...int) ([]*Status, error)
	GetByName(ctx context.Context, names ...string) ([]*Status, error)

	Create(ctx context.Context, ss ...*Status) error
	Update(ctx context.Context, ss ...*Status) error
	Delete(ctx context.Context, ids ...int) error
}
