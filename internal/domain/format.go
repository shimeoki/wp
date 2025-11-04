package domain

import (
	"errors"
	"fmt"
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

func NewInvalidFormatError(value string) error {
	return fmt.Errorf("'%s' is an %w",
		value, ErrInvalidFormat)
}

func ParseFormat(extension string) (Format, error) {
	switch strings.ToLower(extension) {
	case "jpg", "jpeg":
		return JPEG, nil
	case "png":
		return PNG, nil
	}

	return "", NewInvalidFormatError(extension)
}
