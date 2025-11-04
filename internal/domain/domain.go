package domain

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"time"
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
	ErrNotFound          = errors.New("not found")
	ErrAlreadyExists     = errors.New("already exists")
	ErrInvalidRelation   = errors.New("invalid relation")
	ErrInvalidTimestamps = errors.New("invalid timestamps")
)

func NewNotFoundError(entity, with, value string) error {
	return fmt.Errorf("%s with %s '%s' %w",
		entity, with, value, ErrNotFound)
}

func NewAlreadyExistsError(entity, with, value string) error {
	return fmt.Errorf("%s with %s '%s' %w",
		entity, with, value, ErrAlreadyExists,
	)
}

func NewInvalidRelationError(entity, with, value, msg string) error {
	return fmt.Errorf("%w: %s with %s '%s' %s",
		ErrInvalidRelation, entity, with, value, msg)
}

func ValidateTimestamps(created, updated time.Time) error {
	if !created.After(updated) {
		return nil
	}

	return fmt.Errorf("%w: 'created' is '%s' and 'updated' is '%s'",
		ErrInvalidTimestamps, created.String(), updated.String())
}
