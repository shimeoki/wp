package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type RemoveTagProvider interface {
	WallpaperProvider
	TagProvider
}

type RemoveTagWorker Worker[RemoveTagProvider]

type RemoveTagHandler struct {
	worker RemoveTagWorker
}

func NewRemoveTagHandler(w RemoveTagWorker) *RemoveTagHandler {
	return &RemoveTagHandler{worker: w}
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
	var r RemoveTagResult

	if err := h.worker.Work(ctx, func(p RemoveTagProvider) error {
		tags, wallpapers := p.TagRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		name, err := domain.ParseName(cmd.TagName)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return errors.New("wallpaper not found")
		}

		t, _ := tags.FindByName(ctx, name)
		if t == nil {
			return errors.New("tag not found")
		}

		if _, ok := w.Tags[t.ID]; !ok {
			return errors.New("tag not attached")
		}

		if err := w.RemoveTag(t.ID); err != nil {
			return err
		}

		return wallpapers.Save(ctx, w)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
