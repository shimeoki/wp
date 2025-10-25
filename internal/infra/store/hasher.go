package store

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type Hasher interface {
	Hash(io.Reader) (domain.Hash, error)
}

type HasherFunc func(io.Reader) (domain.Hash, error)

func (f HasherFunc) Hash(r io.Reader) (domain.Hash, error) {
	return f(r)
}

func SHA256Hasher() Hasher {
	return HasherFunc(func(r io.Reader) (domain.Hash, error) {
		hasher := sha256.New()

		if _, err := io.Copy(hasher, r); err != nil {
			return domain.Hash{}, err
		}

		return domain.MakeHash(
			domain.SHA256,
			hex.EncodeToString(hasher.Sum(nil)),
		)
	})
}
