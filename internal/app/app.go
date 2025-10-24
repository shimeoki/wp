package app

import (
	"context"
	"errors"
)

type Ctx = context.Context

type Providers any

type Version int

const VERSION Version = 1

var InvalidVersion = errors.New("unexpected version")
