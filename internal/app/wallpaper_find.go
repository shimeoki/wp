package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type FindWallpaperProvider struct {
	WallpaperRepo domain.WallpaperRepo
}

type FindWallpaperWorker Worker[*FindWallpaperProvider]

type FindWallpaperHandler struct {
	worker FindWallpaperWorker
}

func NewFindWallpaperHandler(w FindWallpaperWorker) *FindWallpaperHandler {
	return &FindWallpaperHandler{worker: w}
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

	if err := h.worker.Do(ctx, func(p *FindWallpaperProvider) error {
		hash, err := domain.ParseHash(qry.Hash)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo.FindByHash(ctx, hash)
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
