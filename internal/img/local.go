package img

import (
	"os"
	"path/filepath"
)

type LocalImage struct {
	file      *os.File
	extension Extension
}

func NewLocalImage(path string) (*LocalImage, error) {
	var extension Extension
	switch filepath.Ext(path) {
	case "jpg", "jpeg":
		extension = JPEG
	case "png":
		extension = PNG
	default:
		return nil, InvalidExtension
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return &LocalImage{file: file, extension: extension}, nil
}

func (i *LocalImage) Extension() Extension {
	return i.extension
}

func (i *LocalImage) Read(p []byte) (int, error) {
	return i.file.Read(p)
}

func (i *LocalImage) Close() error {
	return i.file.Close()
}
