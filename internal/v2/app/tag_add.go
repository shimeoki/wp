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
