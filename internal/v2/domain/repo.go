package domain

import (
	"context"
	"iter"
)

type Ctx = context.Context

type Repo[V any] interface {
	Save(Ctx, V) error
	Delete(Ctx, ID) error

	ByID(Ctx, ID) (V, error)

	All(Ctx) (iter.Seq[V], error)
	Count(Ctx) (int, error)
}
