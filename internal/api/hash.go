package api

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"regexp"
)

type Hash string

var InvalidHash = errors.New("invalid hash")

type Hasher interface {
	Compute(io.Reader) (Hash, error)
	Valid(Hash) bool
}

type SHA256Hasher struct{}

func (h *SHA256Hasher) Compute(r io.Reader) (Hash, error) {
	hash := sha256.New()

	if _, err := io.Copy(hash, r); err != nil {
		return "", err
	}

	return Hash(hex.EncodeToString(hash.Sum(nil))), nil
}

func (h *SHA256Hasher) Valid(hash Hash) bool {
	if len(hash) != 64 {
		return false
	}

	match, err := regexp.MatchString("^[a-fA-F0-9]{64}$", string(hash))
	if err != nil {
		return false
	}

	return match
}
