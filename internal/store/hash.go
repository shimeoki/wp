package store

import (
	"crypto/sha256"
	"errors"
	"hash"
	"io"
	"regexp"
)

var InvalidHash = errors.New("invalid hash")

type Hasher interface {
	Compute(r io.Reader) (hash string, err error)
	Valid(hash string) bool
}

type SHA256Hasher struct {
	hsh hash.Hash
}

func NewSHA256Hasher() *SHA256Hasher {
	return &SHA256Hasher{hsh: sha256.New()}
}

func (h *SHA256Hasher) Compute(r io.Reader) (string, error) {
	h.hsh.Reset()

	if _, err := io.Copy(h.hsh, r); err != nil {
		return "", err
	}

	return string(h.hsh.Sum(nil)), nil
}

func (h *SHA256Hasher) Valid(hash string) bool {
	match, err := regexp.MatchString("^[a-fA-F0-9]{64}$", hash)
	if err != nil {
		return false
	}

	return match
}
