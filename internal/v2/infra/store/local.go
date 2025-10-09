package store

import (
	"io"
	"os"

	"github.com/shimeoki/wp/internal/v2/app"
)

type LocalStore struct {
	root   *os.Root
	hasher app.Hasher
}

func NewLocalStore(path string, hasher app.Hasher) (*LocalStore, error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	return &LocalStore{hasher: hasher, root: root}, nil
}

func (s *LocalStore) Create(img io.Reader) (app.Hash, error) {
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

func (s *LocalStore) Remove(h app.Hash) error {
	return s.root.Remove(string(h))
}

func (s *LocalStore) Get(h app.Hash) (io.ReadCloser, error) {
	if !s.hasher.Valid(h) {
		return nil, app.InvalidHash
	}

	return s.root.Open(string(h))
}
