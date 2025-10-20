package app

import (
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type ShowWallpaperWorker interface {
	Worker
	Store() domain.Store
	WallpaperRepo() domain.WallpaperRepo
}

type ShowWallpaperProvider interface {
	Provider[ShowWallpaperWorker]
}

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
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	defer worker.Rollback()
	store, repo := worker.Store(), worker.WallpaperRepo()

	hash, err := domain.ParseHash(qry.Hash)
	if err != nil {
		return nil, err
	}

	w, _ := repo.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	r, err := store.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	if err := worker.Commit(); err != nil {
		return nil, err
	}

	return &ShowWallpaperResult{Image: r, Format: w.Format.String()}, nil
}
