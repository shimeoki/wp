package db

import "context"

type Status struct {
	ID   int
	Name string
}

type StatusRepo interface {
	Create(ctx context.Context, s *Status) error

	GetAll(ctx context.Context) ([]*Status, error)
	GetByID(ctx context.Context, id int) (*Status, error)
	GetByName(ctx context.Context, name string) (*Status, error)

	Update(ctx context.Context, s *Status) error

	Delete(ctx context.Context, s *Status) error
}
