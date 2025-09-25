package api

import (
	"io"
	"os"
)

type Store interface {
	Get(Hash) (io.ReadCloser, error)
	Create(io.Reader) (Hash, error)
	Remove(Hash) error
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

func (s *LocalStore) Create(img io.Reader) (Hash, error) {
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

	file, err := s.root.Create(string(hash))
	if err != nil {
		return "", err
	}

	defer file.Close()
	if _, err := io.Copy(file, tmp); err != nil {
		return "", err
	}

	return hash, nil
}

func (s *LocalStore) Remove(h Hash) error {
	return s.root.Remove(string(h))
}

func (s *LocalStore) Get(h Hash) (io.ReadCloser, error) {
	if !s.hasher.Valid(h) {
		return nil, InvalidHash
	}

	return s.root.Open(string(h))
}
