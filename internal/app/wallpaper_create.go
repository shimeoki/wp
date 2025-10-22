package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateWallpaperProviders struct {
	Store         domain.Store
	WallpaperRepo domain.WallpaperRepo
}

type CreateWallpaperWorker Worker[*CreateWallpaperProviders]

type CreateWallpaperHandler struct {
	worker CreateWallpaperWorker
}

func NewCreateWallpaperHandler(
	w CreateWallpaperWorker,
) *CreateWallpaperHandler {
	return &CreateWallpaperHandler{worker: w}
}

type CreateWallpaperCommand struct {
	Image  io.ReadCloser
	Format string
}

type CreateWallpaperResult struct {
	Hash string
}

func (h *CreateWallpaperHandler) Handle(
	ctx Ctx,
	cmd *CreateWallpaperCommand,
) (*CreateWallpaperResult, error) {
	var r CreateWallpaperResult

	if err := h.worker.Do(ctx, func(p *CreateWallpaperProviders) error {
		hash, err := p.Store.Create(ctx, cmd.Image)
		if err != nil {
			return err
		}

		if wall, _ := p.WallpaperRepo.FindByHash(ctx, hash); wall != nil {
			return errors.New("wallpaper already exists")
		}

		f, err := domain.ParseFormat(cmd.Format)
		if err != nil {
			return err
		}

		wall, err := domain.NewWallpaper(f, hash)
		if err != nil {
			return err
		}

		if err := p.WallpaperRepo.Save(ctx, wall); err != nil {
			return err
		}

		r.Hash = hash.String()

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
