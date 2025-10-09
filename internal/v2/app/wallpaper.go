package app

import (
	"github.com/shimeoki/wp/internal/v2/domain"
)

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

func toDomainFormat(f Format) (domain.Format, error) {
	switch f {
	case JPEG:
		return domain.JPEG, nil
	case PNG:
		return domain.PNG, nil
	}

	return "", domain.InvalidFormat
}

func toAppFormat(f domain.Format) (Format, error) {
	switch f {
	case domain.JPEG:
		return JPEG, nil
	case domain.PNG:
		return PNG, nil
	}

	return "", domain.InvalidFormat
}

type WallpaperResult struct {
	Format
	Hash string
}
