package app

import "github.com/shimeoki/wp/internal/domain"

type AddSourceProvider interface {
	WallpaperProvider
	SourceProvider
}

type AddSourceWorker Worker[AddSourceProvider]

type AddSourceHandler struct {
	worker AddSourceWorker
	logger Logger
}

func NewAddSourceHandler(w AddSourceWorker, l Logger) *AddSourceHandler {
	return &AddSourceHandler{worker: w, logger: l}
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
	var r AddSourceResult

	Act(h.logger, ctx, "adding source", cmd)
	if err := h.worker.Work(ctx, func(p AddSourceProvider) error {
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

		if err := w.AddSource(source); err != nil {
			return err
		}

		return wallpapers.Save(ctx, w)
	}); err != nil {
		Fail(h.logger, ctx, "failed to add source", err)
		return nil, err
	}

	return &r, nil
}
