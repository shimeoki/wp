package app

import (
	"context"
	"errors"
)

type Ctx = context.Context

type Providers any

type Worker[P Providers] interface {
	Do(Ctx, func(P) error) error
}

type Version int

const VERSION Version = 1

var InvalidVersion = errors.New("unexpected version")
