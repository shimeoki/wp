package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type RemoveSourceProvider interface {
	WallpaperProvider
	SourceProvider
}

type RemoveSourceWorker Worker[RemoveSourceProvider]

type RemoveSourceHandler struct {
	worker RemoveSourceWorker
}

func NewRemoveSourceHandler(w RemoveSourceWorker) *RemoveSourceHandler {
	return &RemoveSourceHandler{worker: w}
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
	var r RemoveSourceResult

	if err := h.worker.Work(ctx, func(p RemoveSourceProvider) error {
		sources, wallpapers := p.SourceRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return errors.New("wallpaper not found")
		}

		id, err := domain.ParseID(cmd.SourceID)
		if err != nil {
			return err
		}

		source, _ := sources.FindByID(ctx, id)
		if source == nil {
			return errors.New("source not found")
		}

		if _, ok := w.Sources[source.ID]; !ok {
			return errors.New("source not attached")
		}

		if err := w.RemoveSource(source.ID); err != nil {
			return err
		}

		return wallpapers.Save(ctx, w)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
