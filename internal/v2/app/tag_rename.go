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
	before, _ := h.tags.ByName(ctx, cmd.Before)
	if before == nil {
		return nil, errors.New("tag not found")
	}

	if err := before.Rename(cmd.After); err != nil {
		return nil, err
	}

	after, _ := h.tags.ByName(ctx, cmd.After)
	if after != nil {
		walls, err := h.wallpapers.ByTagID(ctx, after.ID)
		if err != nil {
			return nil, err
		}

		for wall := range walls {
			if err := wall.RemoveTag(after.ID); err != nil {
				return nil, err
			}

			if err := wall.AddTag(before); err != nil {
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

	if err := h.tags.Save(ctx, before); err != nil {
		return nil, err
	}

	return &RenameTagResult{}, nil
}
