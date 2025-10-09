package image

import (
	"os"
	"path/filepath"

	"github.com/shimeoki/wp/internal/v2/app"
)

type LocalImage struct {
	file   *os.File
	format app.Format
}

func NewLocalImage(path string) (*LocalImage, error) {
	var format app.Format
	switch filepath.Ext(path) {
	case "jpg", "jpeg":
		format = app.JPEG
	case "png":
		format = app.PNG
	default:
		return nil, app.InvalidFormat
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return &LocalImage{file: file, format: format}, nil
}

func (i *LocalImage) Format() app.Format {
	return i.format
}

func (i *LocalImage) Read(p []byte) (int, error) {
	return i.file.Read(p)
}

func (i *LocalImage) Close() error {
	return i.file.Close()
}
