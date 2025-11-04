package domain

import (
	"errors"
	"fmt"
)

type Status string

const (
	QUEUED  Status = "queued"
	USED    Status = "used"
	SKIPPED Status = "skipped"
)

var (
	ErrInvalidStatus = errors.New("invalid status")
)

func NewInvalidStatusError(value string) error {
	return fmt.Errorf("'%s' is an %w",
		value, ErrInvalidStatus)
}
