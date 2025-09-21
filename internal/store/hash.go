package store

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
)

var InvalidHash = errors.New("invalid hash")

type Hasher interface {
	Compute(r io.Reader) (hash string, err error)
	Valid(hash string) bool
}

type SHA256Hasher struct{}

func (h *SHA256Hasher) Compute(r io.Reader) (string, error) {
	hash := sha256.New()

	if _, err := io.Copy(hash, r); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (h *SHA256Hasher) Valid(hash string) bool {
	if len(hash) != 64 {
		return false
	}

	match, err := regexp.MatchString("^[a-fA-F0-9]{64}$", hash)
	if err != nil {
		return false
	}

	return match
}
