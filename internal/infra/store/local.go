package store

import (
	"io"
	"os"

	"github.com/shimeoki/wp/internal/domain"
)

type LocalStore struct {
	root   *os.Root
	hasher Hasher
}

func NewLocalStore(path string, h Hasher) (*LocalStore, error) {
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, err
	}

	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	return &LocalStore{hasher: h, root: root}, nil
}

func (s *LocalStore) Create(
	ctx domain.Ctx,
	img io.Reader,
) (domain.Hash, error) {
	tmp, err := os.CreateTemp("", "wp-local-store")
	if err != nil {
		return domain.Hash{}, nil
	}

	defer os.Remove(tmp.Name())
	defer tmp.Close()

	r := io.TeeReader(img, tmp)

	hash, err := s.hasher.Compute(r)
	if err != nil {
		return domain.Hash{}, err
	}

	stored, err := s.Get(ctx, hash)
	if stored != nil {
		stored.Close()
		return hash, nil
	}

	if _, err := tmp.Seek(0, 0); err != nil {
		return domain.Hash{}, err
	}

	file, err := s.root.Create(hash.String())
	if err != nil {
		return domain.Hash{}, err
	}

	defer file.Close()
	if _, err := io.Copy(file, tmp); err != nil {
		return domain.Hash{}, err
	}

	return hash, nil
}

func (s *LocalStore) Delete(ctx domain.Ctx, h domain.Hash) error {
	return s.root.Remove(h.String())
}

func (s *LocalStore) Get(ctx domain.Ctx, h domain.Hash) (io.ReadCloser, error) {
	r, err := s.root.Open(h.String())
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *LocalStore) Count(ctx domain.Ctx) (int, error) {
	f, err := s.root.Open(".")
	if err != nil {
		return 0, err
	}

	defer f.Close()

	entries, err := f.ReadDir(0)
	if err != nil {
		return 0, err
	}

	return len(entries), nil
}

func (s *LocalStore) Close() error {
	return s.root.Close()
}
