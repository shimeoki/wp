package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type AddTagWorker interface {
	Worker
	WallpaperRepo() domain.WallpaperRepo
	TagRepo() domain.TagRepo
}

type AddTagProvider interface {
	Provider[AddTagWorker]
}

type AddTagHandler struct {
	provider AddTagProvider
}

func NewAddTagHandler(
	p AddTagProvider,
) *AddTagHandler {
	return &AddTagHandler{provider: p}
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
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	defer worker.Rollback()
	wallpapers, tags := worker.WallpaperRepo(), worker.TagRepo()

	hash, err := domain.ParseHash(cmd.WallpaperHash)
	if err != nil {
		return nil, err
	}

	name, err := domain.ParseName(cmd.TagName)
	if err != nil {
		return nil, err
	}

	w, _ := wallpapers.FindByHash(ctx, hash)
	if w == nil {
		return nil, errors.New("wallpaper not found")
	}

	t, _ := tags.FindByName(ctx, name)
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

	if err := wallpapers.Save(ctx, w); err != nil {
		return nil, err
	}

	if err := worker.Commit(); err != nil {
		return nil, err
	}

	return &AddTagResult{}, nil
}
