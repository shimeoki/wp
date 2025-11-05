package app

import "github.com/shimeoki/wp/internal/domain"

type RemoveTagProvider interface {
	WallpaperProvider
	TagProvider
}

type RemoveTagWorker Worker[RemoveTagProvider]

type RemoveTagHandler struct {
	worker RemoveTagWorker
	logger Logger
}

func NewRemoveTagHandler(w RemoveTagWorker, l Logger) *RemoveTagHandler {
	return &RemoveTagHandler{worker: w, logger: l}
}

type RemoveTagCommand struct {
	WallpaperHash string
	TagName       string
}

type RemoveTagResult struct{}

func (h *RemoveTagHandler) Handle(
	ctx Ctx,
	cmd *RemoveTagCommand,
) (*RemoveTagResult, error) {
	var r RemoveTagResult

	Act(h.logger, ctx, "removing tag", cmd)
	if err := h.worker.Work(ctx, func(p RemoveTagProvider) error {
		tags, wallpapers := p.TagRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		name, err := domain.ParseName(cmd.TagName)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		t, _ := tags.FindByName(ctx, name)
		if t == nil {
			return domain.NewNotFoundError("tag", "name", name.String())
		}

		if err := w.RemoveTag(t.ID); err != nil {
			return err
		}

		return wallpapers.Save(ctx, w)
	}); err != nil {
		Fail(h.logger, ctx, "failed to remove tag", err)
		return nil, err
	}

	return &r, nil
}
