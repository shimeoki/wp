package domain

import (
	"errors"
	"strings"
)

type Format string

func (f Format) String() string {
	return string(f)
}

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var (
	ErrInvalidFormat = errors.New("invalid format")
)

func ParseFormat(extension string) (Format, error) {
	switch strings.ToLower(extension) {
	case "jpeg":
	case "jpg":
		return JPEG, nil
	case "png":
		return PNG, nil
	}

	return "", ErrInvalidFormat
}
