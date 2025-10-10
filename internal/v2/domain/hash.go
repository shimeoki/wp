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

func (h Hash) String() string {
	return string(h)
}

func ParseHash(value string) (Hash, error) {
	if len(value) == 0 {
		return "", EmptyHash
	}

	return Hash(value), nil
}

var (
	InvalidHash = errors.New("invalid hash")
	EmptyHash   = errors.New("hash is empty")
)

type Hasher interface {
	Compute(io.Reader) (Hash, error)
	Valid(Hash) bool
}
