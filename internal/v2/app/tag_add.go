package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type AddTagHandler struct {
	wallpapers domain.WallpaperRepo
	tags       domain.TagRepo
}

func NewAddTagHandler(
	wallpapers domain.WallpaperRepo,
	tags domain.TagRepo,
) *AddTagHandler {
	return &AddTagHandler{
		wallpapers: wallpapers,
		tags:       tags,
	}
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
	w, _ := h.wallpapers.ByHash(ctx, domain.Hash(cmd.WallpaperHash))
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	t, _ := h.tags.ByName(ctx, cmd.TagName)
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

	if err := h.wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	return &AddTagResult{}, nil
}
