package domain

import (
	"context"
	"errors"
	"iter"
)

type Ctx = context.Context

type Repo[E any] interface {
	Save(Ctx, E) error
	Delete(Ctx, ID) error
	FindByID(Ctx, ID) (E, error)

	All(Ctx) (iter.Seq[E], error)
	Count(Ctx) (int, error)
}

var (
	NotFound error = errors.New("not found")
)
