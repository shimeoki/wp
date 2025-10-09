package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type RemoveTagCommand struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewRemoveTagCommand(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *RemoveTagCommand {
	return &RemoveTagCommand{
		wallpapers: wallpapers,
		tags:       tags,
	}
}

type RemoveTagData struct {
	WallpaperHash string
	TagName       string
}

type RemoveTagResult struct{}

func (cmd *RemoveTagCommand) Execute(
	ctx Ctx,
	data *RemoveTagData,
) (*RemoveTagResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, domain.Hash(data.WallpaperHash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	t, _ := cmd.tags.ByName(ctx, data.TagName)
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if _, ok := w.Tags[t.ID]; !ok {
		return nil, errors.New("tag not attached")
	}

	if err := w.RemoveTag(t.ID); err != nil {
		return nil, err
	}

	if err := cmd.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &RemoveTagResult{}, nil
}
