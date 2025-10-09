package domain

import (
	"errors"
	"io"
)

type Store interface {
	Get(Hash) (io.ReadCloser, error)
	Create(io.Reader) (Hash, error)
	Remove(Hash) error
}

type Hash string

var InvalidHash = errors.New("invalid hash")

type Hasher interface {
	Compute(io.Reader) (Hash, error)
	Valid(Hash) bool
}
