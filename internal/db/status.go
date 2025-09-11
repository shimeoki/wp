package db

import "context"

type Status struct {
	ID   int
	Name string
}

type StatusRepo interface {
	GetAll(ctx context.Context) ([]*Status, error)
	GetByID(ctx context.Context, id int) (*Status, error)
	GetByName(ctx context.Context, name string) (*Status, error)

	Create(ctx context.Context, s *Status) error
	Update(ctx context.Context, s *Status) error
	Delete(ctx context.Context, id int) error
}
