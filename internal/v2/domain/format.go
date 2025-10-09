package domain

import "errors"

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var InvalidFormat = errors.New("invalid format")
