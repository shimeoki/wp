package domain

import (
	"context"
	"iter"

	"github.com/google/uuid"
)

type Ctx = context.Context

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.Must(uuid.NewV7()))
}

func (id ID) String() string {
	return (uuid.UUID)(id).String()
}

type Repo[V any] interface {
	Save(Ctx, V) error
	Delete(Ctx, ID) error

	ByID(Ctx, ID) (V, error)

	All(Ctx) (iter.Seq[V], error)
	Count(Ctx) (int, error)
}
