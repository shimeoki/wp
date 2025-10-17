package domain

import "errors"

type Status string

const (
	QUEUED  Status = "queued"
	USED    Status = "used"
	SKIPPED Status = "skipped"
)

var InvalidStatus = errors.New("invalid status")
