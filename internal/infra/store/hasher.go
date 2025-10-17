package store

import (
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type Hasher interface {
	Compute(io.Reader) (domain.Hash, error)
	Valid(domain.Hash) bool
}
