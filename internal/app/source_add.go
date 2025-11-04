package app

import "github.com/shimeoki/wp/internal/domain"

type AddSourceProvider interface {
	WallpaperProvider
	SourceProvider
}

type AddSourceWorker Worker[AddSourceProvider]

type AddSourceHandler struct {
	worker AddSourceWorker
}

func NewAddSourceHandler(w AddSourceWorker) *AddSourceHandler {
	return &AddSourceHandler{worker: w}
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
		return nil, err
	}

	return &r, nil
}
