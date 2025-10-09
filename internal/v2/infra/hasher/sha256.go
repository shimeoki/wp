package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"regexp"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type SHA256Hasher struct{}

func (h *SHA256Hasher) Compute(r io.Reader) (domain.Hash, error) {
	hash := sha256.New()

	if _, err := io.Copy(hash, r); err != nil {
		return "", err
	}

	return domain.Hash(hex.EncodeToString(hash.Sum(nil))), nil
}

func (h *SHA256Hasher) Valid(hash domain.Hash) bool {
	if len(hash) != 64 {
		return false
	}

	match, err := regexp.MatchString("^[a-fA-F0-9]{64}$", string(hash))
	if err != nil {
		return false
	}

	return match
}
