package app

import (
	"context"
	"errors"
)

type Ctx = context.Context

type Provider[W any] func(Ctx, func(W) error) error

type Version int

const VERSION Version = 1

var InvalidVersion = errors.New("unexpected version")
