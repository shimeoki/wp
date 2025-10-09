package domain

import (
	"errors"
	"strings"
)

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var InvalidFormat = errors.New("invalid format")

func ParseFormat(extension string) (Format, error) {
	switch strings.ToLower(extension) {
	case "jpeg":
	case "jpg":
		return JPEG, nil
	case "png":
		return PNG, nil
	}

	return "", InvalidFormat
}
