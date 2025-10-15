package store

import (
	"io"
	"os"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type LocalStore struct {
	root   *os.Root
	hasher domain.Hasher
}

func NewLocalStore(path string, hasher domain.Hasher) (*LocalStore, error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	return &LocalStore{hasher: hasher, root: root}, nil
}

func (s *LocalStore) Create(img io.Reader) (domain.Hash, error) {
	tmp, err := os.CreateTemp("", "wp-local-store")
	if err != nil {
		return "", err
	}

	defer os.Remove(tmp.Name())
	defer tmp.Close()

	r := io.TeeReader(img, tmp)

	hash, err := s.hasher.Compute(r)
	if err != nil {
		return "", err
	}

	stored, err := s.Get(hash)
	if stored != nil {
		stored.Close()
		return hash, nil
	}

	file, err := s.root.Create(hash.String())
	if err != nil {
		return "", err
	}

	defer file.Close()
	if _, err := io.Copy(file, tmp); err != nil {
		return "", err
	}

	return hash, nil
}

func (s *LocalStore) Remove(h domain.Hash) error {
	return s.root.Remove(h.String())
}

func (s *LocalStore) Get(h domain.Hash) (io.ReadCloser, error) {
	if !s.hasher.Valid(h) {
		return nil, domain.InvalidHash
	}

	r, err := s.root.Open(h.String())
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *LocalStore) Close() error {
	return s.root.Close()
}
