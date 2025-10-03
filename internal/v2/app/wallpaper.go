package app

import (
	"errors"
	"io"
	"time"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type WallpaperService struct {
	store  Store
	hasher Hasher
	repo   domain.WallpaperRepo
}

func NewWallpaperService(s Store, h Hasher) *WallpaperService {
	return &WallpaperService{
		store:  s,
		hasher: h,
	}
}

type Store interface {
	Get(Hash) (io.ReadCloser, error)
	Create(io.Reader) (Hash, error)
	Remove(Hash) error
}

type Hash string

var InvalidHash = errors.New("invalid hash")

type Hasher interface {
	Compute(io.Reader) (Hash, error)
	Valid(Hash) bool
}

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var InvalidFormat = errors.New("invalid format")

type Image interface {
	io.ReadCloser
	Format() Format
}

type WallpaperResult struct {
	Format
	Hash

	CreatedAt time.Time
	UpdatedAt time.Time
}

type AliasResult struct {
	Name string
}
