package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type AddTagWorker struct {
	WallpaperRepo domain.WallpaperRepo
	TagRepo       domain.TagRepo
}

type AddTagProvider Provider[*AddTagWorker]

type AddTagHandler struct {
	provider AddTagProvider
}

func NewAddTagHandler(
	p AddTagProvider,
) *AddTagHandler {
	return &AddTagHandler{provider: p}
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

	if err := h.provider(ctx, func(w *AddTagWorker) error {
		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		name, err := domain.ParseName(cmd.TagName)
		if err != nil {
			return err
		}

		wall, _ := w.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		tag, _ := w.TagRepo.FindByName(ctx, name)
		if tag == nil {
			// automatically create tag?
			return errors.New("tag not found")
		}

		if _, ok := wall.Tags[tag.ID]; ok {
			return errors.New("tag already attached")
		}

		if err := wall.AddTag(tag); err != nil {
			return err
		}

		return w.WallpaperRepo.Save(ctx, wall)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
