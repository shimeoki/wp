package app

import "github.com/shimeoki/wp/internal/domain"

type AddTagProvider interface {
	WallpaperProvider
	TagProvider
}

type AddTagWorker Worker[AddTagProvider]

type AddTagHandler struct {
	worker AddTagWorker
	logger Logger
}

func NewAddTagHandler(w AddTagWorker, l Logger) *AddTagHandler {
	return &AddTagHandler{worker: w, logger: l}
}

type AddTagCommand struct {
	WallpaperHash string
	TagName       string
}

type AddTagResult struct{}

func (h *AddTagHandler) Handle(
	ctx Ctx,
	cmd *AddTagCommand,
) (*AddTagResult, error) {
	var r AddTagResult

	Act(h.logger, ctx, "adding tag", cmd)
	if err := h.worker.Work(ctx, func(p AddTagProvider) error {
		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		name, err := domain.ParseName(cmd.TagName)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo().FindByHash(ctx, hash)
		if wall == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		tag, _ := p.TagRepo().FindByName(ctx, name)
		if tag == nil {
			// automatically create tag?
			return domain.NewNotFoundError("tag", "name", name.String())
		}

		if err := wall.AddTag(tag); err != nil {
			return err
		}

		return p.WallpaperRepo().Save(ctx, wall)
	}); err != nil {
		Fail(h.logger, ctx, "failed to add tag", err)
		return nil, err
	}

	return &r, nil
}
