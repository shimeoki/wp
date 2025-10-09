package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type WallpaperService struct {
	store      Store
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
	sources    domain.SourceRepo
}

func NewWallpaperService(
	store Store,
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
	sources domain.SourceRepo,
) *WallpaperService {
	return &WallpaperService{
		store:      store,
		wallpapers: wallpapers,
		tags:       tags,
		sources:    sources,
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

func toDomainFormat(f Format) (domain.Format, error) {
	switch f {
	case JPEG:
		return domain.JPEG, nil
	case PNG:
		return domain.PNG, nil
	}

	return "", InvalidFormat
}

func toAppFormat(f domain.Format) (Format, error) {
	switch f {
	case domain.JPEG:
		return JPEG, nil
	case domain.PNG:
		return PNG, nil
	}

	return "", InvalidFormat
}

type Image interface {
	io.ReadCloser
	Format() Format
}

type image struct {
	io.ReadCloser
	format Format
}

func (i *image) Format() Format {
	return i.format
}

type WallpaperResult struct {
	Format
	Hash
}
