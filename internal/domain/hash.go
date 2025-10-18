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
	Remove(Ctx, Hash) error
}

var (
	InvalidDigest = errors.New("invalid hash digest")
	InvalidAlgo   = errors.New("invalid hash algorithm")
	InvalidHash   = errors.New("invalid hash value")
)

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

	return "", InvalidAlgo
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
		return Hash{}, InvalidHash
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

	return InvalidAlgo
}

func (h Hash) validateSHA256() error {
	if len(h.Digest) != 64 {
		return InvalidDigest
	}

	match, err := regexp.MatchString("^[a-fA-F0-9]{64}$", h.Digest)
	if err != nil || !match {
		return InvalidDigest
	}

	return nil
}

func (h Hash) validateMD5() error {
	if len(h.Digest) != 32 {
		return InvalidDigest
	}

	match, err := regexp.MatchString("^[a-fA-F0-9]{32}$", h.Digest)
	if err != nil || !match {
		return InvalidDigest
	}

	return nil
}

func (h Hash) String() string {
	return fmt.Sprintf("%s-%s", h.Algo, h.Digest)
}
