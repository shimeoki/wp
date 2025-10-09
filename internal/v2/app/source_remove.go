package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shimeoki/wp/internal/v2/domain"
)

type RemoveSourceHandler struct {
	wallpapers domain.WallpaperRepo
	sources    domain.SourceRepo
}

func NewRemoveSourceHandler(
	wallpapers domain.WallpaperRepo,
	sources domain.SourceRepo,
) *RemoveSourceHandler {
	return &RemoveSourceHandler{
		wallpapers: wallpapers,
		sources:    sources,
	}
}

type RemoveSourceCommand struct {
	WallpaperHash string
	SourceID      string
}

type RemoveSourceResult struct{}

func (h *RemoveSourceHandler) Handle(
	ctx Ctx,
	cmd *RemoveSourceCommand,
) (*RemoveSourceResult, error) {
	w, _ := h.wallpapers.FindByHash(ctx, domain.Hash(cmd.WallpaperHash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	id, err := uuid.Parse(cmd.SourceID)
	if err != nil {
		return nil, err
	}

	source, _ := h.sources.FindByID(ctx, domain.ID(id))
	if source == nil {
		return nil, errors.New("source not found")
	}

	if _, ok := w.Sources[source.ID]; !ok {
		return nil, errors.New("source not attached")
	}

	if err := w.RemoveSource(source.ID); err != nil {
		return nil, err
	}

	if err := h.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &RemoveSourceResult{}, nil
}
