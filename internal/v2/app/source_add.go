package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type AddSourceHandler struct {
	wallpapers domain.WallpaperRepo
	sources    domain.SourceRepo
}

func NewAddSourceHandler(
	wallpapers domain.WallpaperRepo,
	sources domain.SourceRepo,
) *AddSourceHandler {
	return &AddSourceHandler{
		wallpapers: wallpapers,
		sources:    sources,
	}
}

type AddSourceCommand struct {
	WallpaperHash string
	SourceID      string
}

type AddSourceResult struct{}

func (h *AddSourceHandler) Handle(
	ctx Ctx,
	cmd *AddSourceCommand,
) (*AddSourceResult, error) {
	hash, err := domain.ParseHash(cmd.WallpaperHash)
	if err != nil {
		return nil, err
	}

	w, _ := h.wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	id, err := domain.ParseID(cmd.SourceID)
	if err != nil {
		return nil, err
	}

	source, _ := h.sources.FindByID(ctx, id)
	if source == nil {
		return nil, errors.New("source not found")
	}

	if _, ok := w.Sources[source.ID]; ok {
		return nil, errors.New("source already attached")
	}

	if err := w.AddSource(source); err != nil {
		return nil, err
	}

	if err := h.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &AddSourceResult{}, nil
}
