package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type RenameTagCommand struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewRenameTagCommand(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *RenameTagCommand {
	return &RenameTagCommand{
		wallpapers: wallpapers,
		tags:       tags,
	}
}

type RenameTagData struct {
	Before string
	After  string
}

type RenameTagResult struct{}

func (cmd *RenameTagCommand) Execute(
	ctx Ctx,
	data *RenameTagData,
) (*RenameTagResult, error) {
	before, _ := cmd.tags.ByName(ctx, data.Before)
	if before == nil {
		return nil, errors.New("tag not found")
	}

	if err := before.Rename(data.After); err != nil {
		return nil, err
	}

	after, _ := cmd.tags.ByName(ctx, data.After)
	if after != nil {
		walls, err := cmd.wallpapers.ByTagID(ctx, after.ID)
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

			if err := cmd.wallpapers.Save(ctx, wall); err != nil {
				return nil, err
			}
		}

		if err := cmd.tags.Delete(ctx, after.ID); err != nil {
			return nil, err
		}
	}

	if err := cmd.tags.Save(ctx, before); err != nil {
		return nil, err
	}

	return &RenameTagResult{}, nil
}
