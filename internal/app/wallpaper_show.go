package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type ShowWallpaperWorker struct {
	Store         domain.Store
	WallpaperRepo domain.WallpaperRepo
}

type ShowWallpaperProvider Provider[*ShowWallpaperWorker]

type ShowWallpaperHandler struct {
	provider ShowWallpaperProvider
}

func NewShowWallpaperHandler(
	p ShowWallpaperProvider,
) *ShowWallpaperHandler {
	return &ShowWallpaperHandler{provider: p}
}

type ShowWallpaperQuery struct {
	Hash string
}

type ShowWallpaperResult struct {
	Image  io.ReadCloser
	Format string
}

func (h *ShowWallpaperHandler) Handle(
	ctx Ctx,
	qry *ShowWallpaperQuery,
) (*ShowWallpaperResult, error) {
	var r ShowWallpaperResult

	if err := h.provider(ctx, func(w *ShowWallpaperWorker) error {
		hash, err := domain.ParseHash(qry.Hash)
		if err != nil {
			return err
		}

		wall, _ := w.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		img, err := w.Store.Get(ctx, hash)
		if err != nil {
			return err
		}

		r.Image = img
		r.Format = wall.Format.String()

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
