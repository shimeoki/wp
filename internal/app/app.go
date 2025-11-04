package app

import "github.com/shimeoki/wp/internal/domain"

type Ctx = domain.Ctx

var (
	ErrNotFound      = domain.ErrNotFound
	ErrAlreadyExists = domain.ErrAlreadyExists
)
