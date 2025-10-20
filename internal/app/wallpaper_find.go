package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type FindWallpaperWorker interface {
	Worker
	WallpaperRepo() domain.WallpaperRepo
}

type FindWallpaperProvider interface {
	Provider[FindWallpaperWorker]
}

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
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	repo := worker.WallpaperRepo()

	hash, err := domain.ParseHash(qry.Hash)
	if err != nil {
		return nil, err
	}

	w, _ := repo.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	return &FindWallpaperResult{Format: w.Format.String()}, nil
}
