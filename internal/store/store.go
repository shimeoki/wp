package store

import (
	"io"
	"os"
)

type Store interface {
	Get(hash string) (img io.ReadCloser, err error)
	Create(img io.Reader) (hash string, err error)
	Remove(hash string) error
}

type LocalStore struct {
	root   *os.Root
	hasher Hasher
}

func NewLocalStore(path string, hasher Hasher) (*LocalStore, error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	return &LocalStore{hasher: hasher, root: root}, nil
}

func (s *LocalStore) Create(img io.Reader) (string, error) {
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

	file, err := s.root.Create(hash)
	if err != nil {
		return "", err
	}

	defer file.Close()
	if _, err := io.Copy(file, tmp); err != nil {
		return "", err
	}

	return hash, nil
}

func (s *LocalStore) Remove(hash string) error {
	return s.root.Remove(hash)
}

func (s *LocalStore) Get(hash string) (io.ReadCloser, error) {
	if !s.hasher.Valid(hash) {
		return nil, InvalidHash
	}

	return s.root.Open(hash)
}
