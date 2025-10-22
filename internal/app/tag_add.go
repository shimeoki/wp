package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type AddTagProviders struct {
	WallpaperRepo domain.WallpaperRepo
	TagRepo       domain.TagRepo
}

type AddTagWorker Worker[*AddTagProviders]

type AddTagHandler struct {
	worker AddTagWorker
}

func NewAddTagHandler(w AddTagWorker) *AddTagHandler {
	return &AddTagHandler{worker: w}
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

	if err := h.worker.Do(ctx, func(p *AddTagProviders) error {
		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		name, err := domain.ParseName(cmd.TagName)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo.FindByHash(ctx, hash)
		if wall == nil {
			return errors.New("wallpaper not found")
		}

		tag, _ := p.TagRepo.FindByName(ctx, name)
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

		return p.WallpaperRepo.Save(ctx, wall)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
