package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type CreateWallpaperCommand struct {
	store      domain.Store
	wallpapers domain.WallpaperRepo
}

func NewCreateWallpaperCommand(
	store domain.Store,
	wallpapers domain.WallpaperRepo,
) *CreateWallpaperCommand {
	return &CreateWallpaperCommand{
		store:      store,
		wallpapers: wallpapers,
	}
}

type CreateWallpaperData struct {
	Image io.ReadCloser
	Format
}

type CreateWallpaperResult struct {
	Hash string
}

func (s *CreateWallpaperCommand) Execute(
	ctx Ctx,
	data *CreateWallpaperData,
) (*CreateWallpaperResult, error) {
	hash, err := s.store.Create(data.Image)
	if err != nil {
		return nil, err
	}

	h := domain.Hash(hash)

	w, _ := s.wallpapers.ByHash(ctx, h)
	if w != nil {
		return nil, errors.New("wallpaper already exists")
	}

	f, err := toDomainFormat(data.Format)
	if err != nil {
		return nil, err
	}

	wall := domain.NewWallpaper(f, h)

	if err := s.wallpapers.Save(ctx, wall); err != nil {
		return nil, err
	}

	return &CreateWallpaperResult{Hash: string(hash)}, nil
}
