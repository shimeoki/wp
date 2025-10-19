package app

import (
	"context"
	"errors"
)

type Ctx = context.Context

type Tx interface {
	Commit() error
	Rollback() error
}

type Version int

const VERSION Version = 1

var InvalidVersion = errors.New("unexpected version")
