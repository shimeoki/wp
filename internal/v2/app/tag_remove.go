package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type RemoveTagHandler struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewRemoveTagHandler(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *RemoveTagHandler {
	return &RemoveTagHandler{
		wallpapers: wallpapers,
		tags:       tags,
	}
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
	hash, err := domain.ParseHash(cmd.WallpaperHash)
	if err != nil {
		return nil, err
	}

	name, err := domain.ParseName(cmd.TagName)
	if err != nil {
		return nil, err
	}

	w, _ := h.wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	t, _ := h.tags.FindByName(ctx, name)
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if _, ok := w.Tags[t.ID]; !ok {
		return nil, errors.New("tag not attached")
	}

	if err := w.RemoveTag(t.ID); err != nil {
		return nil, err
	}

	if err := h.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &RemoveTagResult{}, nil
}
