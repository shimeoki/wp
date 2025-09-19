package img

import (
	"errors"
	"io"
)

type Extension string

const (
	JPEG Extension = "jpg"
	PNG  Extension = "png"
)

var InvalidExtension = errors.New("invalid extension")

type Image interface {
	io.ReadCloser
	Extension() Extension
}
