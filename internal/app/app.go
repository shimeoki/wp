package app

import (
	"context"
	"errors"
)

type Ctx = context.Context

type Worker interface {
	Commit() error
	Rollback() error
}

type Provider[W Worker] interface {
	Provide(Ctx) (W, error)
}

type Version int

const VERSION Version = 1

var InvalidVersion = errors.New("unexpected version")
