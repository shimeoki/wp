package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type RenameTagHandler struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewRenameTagHandler(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *RenameTagHandler {
	return &RenameTagHandler{
		wallpapers: wallpapers,
		tags:       tags,
	}
}

type RenameTagCommand struct {
	Before string
	After  string
}

type RenameTagResult struct{}

func (h *RenameTagHandler) Handle(
	ctx Ctx,
	cmd *RenameTagCommand,
) (*RenameTagResult, error) {
	before, err := domain.ParseName(cmd.Before)
	if err != nil {
		return nil, err
	}

	tag, _ := h.tags.FindByName(ctx, before)
	if tag == nil {
		return nil, errors.New("tag not found")
	}

	after, err := domain.ParseName(cmd.After)
	if err != nil {
		return nil, err
	}

	if err := tag.Rename(after); err != nil {
		return nil, err
	}

	if after, _ := h.tags.FindByName(ctx, after); after != nil {
		walls, err := h.wallpapers.FindByTagID(ctx, after.ID)
		if err != nil {
			return nil, err
		}

		for wall := range walls {
			if err := wall.RemoveTag(after.ID); err != nil {
				return nil, err
			}

			if err := wall.AddTag(tag); err != nil {
				return nil, err
			}

			if err := h.wallpapers.Save(ctx, wall); err != nil {
				return nil, err
			}
		}

		if err := h.tags.Delete(ctx, after.ID); err != nil {
			return nil, err
		}
	}

	if err := h.tags.Save(ctx, tag); err != nil {
		return nil, err
	}

	return &RenameTagResult{}, nil
}
