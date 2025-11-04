package domain

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Store interface {
	Get(Ctx, Hash) (io.ReadCloser, error)
	Create(Ctx, io.Reader) (Hash, error)
	Delete(Ctx, Hash) error
	Count(Ctx) (int, error)
}

var (
	ErrInvalidDigest = errors.New("invalid hash digest")
	ErrInvalidAlgo   = errors.New("invalid hash algorithm")
	ErrInvalidHash   = errors.New("invalid hash value")
)

func NewInvalidAlgoError(algo string) error {
	return fmt.Errorf("'%s' is an %w",
		algo, ErrInvalidAlgo)
}

func NewInvalidDigestError(algo, digest, regex string) error {
	return fmt.Errorf("%w: expected %s for %s, got '%s'",
		ErrInvalidDigest, regex, algo, digest)
}

func NewInvalidHashError(hash string) error {
	return fmt.Errorf("%w: expected '<algo>-<digest>', got '%s'",
		ErrInvalidHash, hash)
}

type Algo string

const (
	SHA256 Algo = "sha256"
	MD5    Algo = "md5"
)

func ParseAlgo(name string) (Algo, error) {
	switch strings.ToLower(name) {
	case "sha256":
		return SHA256, nil
	case "md5":
		return MD5, nil
	}

	return "", NewInvalidAlgoError(name)
}

type Hash struct {
	Algo
	Digest string
}

func MakeHash(a Algo, digest string) (Hash, error) {
	h := Hash{Algo: a, Digest: strings.ToLower(digest)}

	if err := h.validate(); err != nil {
		return Hash{}, err
	}

	return h, nil
}

func ParseHash(value string) (Hash, error) {
	parts := strings.Split(strings.ToLower(value), "-")
	if len(parts) != 2 {
		return Hash{}, NewInvalidHashError(value)
	}

	a, err := ParseAlgo(parts[0])
	if err != nil {
		return Hash{}, err
	}

	return MakeHash(a, parts[1])
}

func (h Hash) validate() error {
	switch h.Algo {
	case SHA256:
		return h.validateSHA256()
	case MD5:
		return h.validateMD5()
	}

	return NewInvalidAlgoError(string(h.Algo))
}

func (h Hash) validateSHA256() error {
	re := "^[a-fA-F0-9]{64}$"

	match, err := regexp.MatchString(re, h.Digest)
	if err != nil || !match {
		return NewInvalidDigestError(string(h.Algo), h.Digest, re)
	}

	return nil
}

func (h Hash) validateMD5() error {
	re := "^[a-fA-F0-9]{32}$"

	match, err := regexp.MatchString(re, h.Digest)
	if err != nil || !match {
		return NewInvalidDigestError(string(h.Algo), h.Digest, re)
	}

	return nil
}

func (h Hash) String() string {
	return fmt.Sprintf("%s-%s", h.Algo, h.Digest)
}
