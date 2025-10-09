package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shimeoki/wp/internal/v2/domain"
)

type AddSourceCommand struct {
	wallpapers domain.WallpaperRepo
	sources    domain.SourceRepo
}

func NewAddSourceCommand(
	wallpapers domain.WallpaperRepo,
	sources domain.SourceRepo,
) *AddSourceCommand {
	return &AddSourceCommand{
		wallpapers: wallpapers,
		sources:    sources,
	}
}

type AddSourceData struct {
	WallpaperHash string
	SourceID      string
}

type AddSourceResult struct{}

func (cmd *AddSourceCommand) Execute(
	ctx Ctx,
	data *AddSourceData,
) (*AddSourceResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, domain.Hash(data.WallpaperHash))
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

	if _, ok := w.Sources[source.ID]; ok {
		return nil, errors.New("source already attached")
	}

	if err := w.AddSource(source); err != nil {
		return nil, err
	}

	if err := cmd.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &AddSourceResult{}, nil
}
