package store

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type Hasher interface {
	Compute(io.Reader) (domain.Hash, error)
}

type SHA256Hasher struct{}

func (h *SHA256Hasher) Compute(r io.Reader) (domain.Hash, error) {
	hasher := sha256.New()

	if _, err := io.Copy(hasher, r); err != nil {
		return domain.Hash{}, err
	}

	return domain.MakeHash(domain.SHA256, hex.EncodeToString(hasher.Sum(nil)))
}
