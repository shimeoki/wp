package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type AddTagCommand struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewAddTagCommand(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *AddTagCommand {
	return &AddTagCommand{
		wallpapers: wallpapers,
		tags:       tags,
	}
}

type AddTagData struct {
	WallpaperHash string
	TagName       string
}

type AddTagResult struct{}

func (cmd *AddTagCommand) Execute(
	ctx Ctx,
	data *AddTagData,
) (*AddTagResult, error) {
	w, _ := cmd.wallpapers.ByHash(ctx, string(data.WallpaperHash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	t, _ := cmd.tags.ByName(ctx, data.TagName)
	if t == nil {
		// automatically create tag?
		return nil, errors.New("tag not found")
	}

	if _, ok := w.Tags[t.ID]; ok {
		return nil, errors.New("tag already attached")
	}

	if err := w.AddTag(t); err != nil {
		return nil, err
	}

	if err := cmd.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &AddTagResult{}, nil
}
