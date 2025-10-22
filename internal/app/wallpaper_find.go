package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type FindWallpaperWorker struct {
	WallpaperRepo domain.WallpaperRepo
}

type FindWallpaperProvider Provider[*FindWallpaperWorker]

type FindWallpaperHandler struct {
	provider FindWallpaperProvider
}

func NewFindWallpaperHandler(
	p FindWallpaperProvider,
) *FindWallpaperHandler {
	return &FindWallpaperHandler{provider: p}
}

type FindWallpaperQuery struct {
	Hash string
}

type FindWallpaperResult struct {
	Format string
}

func (h *FindWallpaperHandler) Handle(
	ctx Ctx,
	qry *FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	var r FindWallpaperResult

	if err := h.provider(ctx, func(w *FindWallpaperWorker) error {
		hash, err := domain.ParseHash(qry.Hash)
		if err != nil {
			return err
		}

		wall, _ := w.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		r.Format = wall.Format.String()

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
