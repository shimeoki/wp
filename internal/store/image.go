package store

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var InvalidFormat = errors.New("invalid format")

type Image interface {
	io.ReadCloser
	Format() Format
}

type LocalImage struct {
	file   *os.File
	format Format
}

func NewLocalImage(path string) (*LocalImage, error) {
	var format Format
	switch filepath.Ext(path) {
	case "jpg", "jpeg":
		format = JPEG
	case "png":
		format = PNG
	default:
		return nil, InvalidFormat
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return &LocalImage{file: file, format: format}, nil
}

func (i *LocalImage) Format() Format {
	return i.format
}

func (i *LocalImage) Read(p []byte) (int, error) {
	return i.file.Read(p)
}

func (i *LocalImage) Close() error {
	return i.file.Close()
}
