package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateWallpaperWorker struct {
	Store         domain.Store
	WallpaperRepo domain.WallpaperRepo
}

type CreateWallpaperProvider Provider[*CreateWallpaperWorker]

type CreateWallpaperHandler struct {
	provider CreateWallpaperProvider
}

func NewCreateWallpaperHandler(
	p CreateWallpaperProvider,
) *CreateWallpaperHandler {
	return &CreateWallpaperHandler{provider: p}
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

	if err := h.provider(ctx, func(w *CreateWallpaperWorker) error {
		hash, err := w.Store.Create(ctx, cmd.Image)
		if err != nil {
			return err
		}

		if wall, _ := w.WallpaperRepo.FindByHash(ctx, hash); wall != nil {
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

		if err := w.WallpaperRepo.Save(ctx, wall); err != nil {
			return err
		}

		r.Hash = hash.String()

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
