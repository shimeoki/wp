package app

import "github.com/shimeoki/wp/internal/domain"

type RemoveSourceProvider interface {
	WallpaperProvider
	SourceProvider
}

type RemoveSourceWorker Worker[RemoveSourceProvider]

type RemoveSourceHandler struct {
	worker RemoveSourceWorker
	logger Logger
}

func NewRemoveSourceHandler(
	w RemoveSourceWorker,
	l Logger,
) *RemoveSourceHandler {
	return &RemoveSourceHandler{worker: w, logger: l}
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

	Act(h.logger, ctx, "removing source", cmd)
	if err := h.worker.Work(ctx, func(p RemoveSourceProvider) error {
		sources, wallpapers := p.SourceRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		id, err := domain.ParseID(cmd.SourceID)
		if err != nil {
			return err
		}

		source, _ := sources.FindByID(ctx, id)
		if source == nil {
			return domain.NewNotFoundError("source", "id", id.String())
		}

		if err := w.RemoveSource(source.ID); err != nil {
			return err
		}

		return wallpapers.Save(ctx, w)
	}); err != nil {
		Fail(h.logger, ctx, "failed to remove source", err)
		return nil, err
	}

	return &r, nil
}
