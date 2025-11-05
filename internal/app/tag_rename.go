package app

import "github.com/shimeoki/wp/internal/domain"

type RenameTagProvider interface {
	WallpaperProvider
	TagProvider
}

type RenameTagWorker Worker[RenameTagProvider]

type RenameTagHandler struct {
	worker RenameTagWorker
	logger Logger
}

func NewRenameTagHandler(w RenameTagWorker, l Logger) *RenameTagHandler {
	return &RenameTagHandler{worker: w, logger: l}
}

type RenameTagCommand struct {
	Before string
	After  string
}

type RenameTagResult struct{}

// TODO: refactor, split into small functions

func (h *RenameTagHandler) Handle(
	ctx Ctx,
	cmd *RenameTagCommand,
) (*RenameTagResult, error) {
	var r RenameTagResult

	Act(h.logger, ctx, "renaming tag", cmd)
	if err := h.worker.Work(ctx, func(p RenameTagProvider) error {
		wallpapers, tags := p.WallpaperRepo(), p.TagRepo()

		before, err := domain.ParseName(cmd.Before)
		if err != nil {
			return err
		}

		tag, _ := tags.FindByName(ctx, before)
		if tag == nil {
			return domain.NewNotFoundError("tag", "name", before.String())
		}

		after, err := domain.ParseName(cmd.After)
		if err != nil {
			return err
		}

		if err := tag.Rename(after); err != nil {
			return err
		}

		if existent, _ := tags.FindByName(ctx, after); existent != nil {
			walls, err := wallpapers.FindByTagID(ctx, existent.ID)
			if err != nil {
				return err
			}

			for wall := range walls {
				if err := wall.RemoveTag(existent.ID); err != nil {
					return err
				}

				if err := wall.AddTag(tag); err != nil {
					return err
				}

				if err := wallpapers.Save(ctx, wall); err != nil {
					return err
				}
			}

			if err := tags.Delete(ctx, existent.ID); err != nil {
				return err
			}
		}

		return tags.Save(ctx, tag)
	}); err != nil {
		Fail(h.logger, ctx, "failed to rename tag", err)
		return nil, err
	}

	return &r, nil
}
