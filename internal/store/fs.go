package store

import (
	"io"
	"os"
)

type FS interface {
	Open(name string) (io.ReadCloser, error)
	Create(name string) (io.WriteCloser, error)
	Remove(name string) error
}

type LocalFS struct {
	root *os.Root
}

func NewLocalFS(path string) (*LocalFS, error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}

	return &LocalFS{root: root}, nil
}

func (fsys *LocalFS) Open(name string) (io.ReadCloser, error) {
	return fsys.root.Open(name)
}

func (fsys *LocalFS) Create(name string) (io.WriteCloser, error) {
	return fsys.root.Create(name)
}

func (fsys *LocalFS) Remove(name string) error {
	return fsys.root.Remove(name)
}
