package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shimeoki/wp/internal/v2/domain"
)

type RemoveSourceCommand struct {
	wallpapers domain.WallpaperRepo
	sources    domain.SourceRepo
}

func NewRemoveSourceCommand(
	wallpapers domain.WallpaperRepo,
	sources domain.SourceRepo,
) *RemoveSourceCommand {
	return &RemoveSourceCommand{
		wallpapers: wallpapers,
		sources:    sources,
	}
}

type RemoveSourceData struct {
	Hash     string
	SourceID string
}

type RemoveSourceResult struct{}

func (cmd *RemoveSourceCommand) Execute(
	ctx Ctx,
	data *RemoveSourceData,
) (*RemoveSourceResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, string(data.Hash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	id, err := uuid.Parse(data.SourceID)
	if err != nil {
		return nil, err
	}

	source, _ := cmd.sources.ByID(ctx, domain.ID(id))
	if source == nil {
		return nil, errors.New("source not found")
	}

	if _, ok := w.Sources[source.ID]; !ok {
		return nil, errors.New("source not attached")
	}

	if err := w.RemoveSource(source.ID); err != nil {
		return nil, err
	}

	if err := cmd.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &RemoveSourceResult{}, nil
}
