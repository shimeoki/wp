package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type RemoveAliasProvider interface {
	AliasProvider
	WallpaperProvider
}

type RemoveAliasWorker Worker[RemoveAliasProvider]

type RemoveAliasHandler struct {
	worker RemoveAliasWorker
}

func NewRemoveAliasHandler(w RemoveAliasWorker) *RemoveAliasHandler {
	return &RemoveAliasHandler{worker: w}
}

type RemoveAliasCommand struct {
	WallpaperHash string
	AliasName     string
}

type RemoveAliasResult struct{}

func (h *RemoveAliasHandler) Handle(
	ctx Ctx,
	cmd *RemoveAliasCommand,
) (*RemoveAliasResult, error) {
	var r RemoveAliasResult

	if err := h.worker.Work(ctx, func(p RemoveAliasProvider) error {
		aliases, wallpapers := p.AliasRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return errors.New("wallpaper not found")
		}

		name, err := domain.ParseName(cmd.AliasName)
		if err != nil {
			return err
		}

		// NOTE: probably there should be a constraint to make unique
		// aliases by name for a wallpaper, but works for now

		aa, err := aliases.FindByWallpaperID(ctx, w.ID)
		if err != nil {
			return err
		}

		for a := range aa {
			if a.Name.String() == name.String() {
				return aliases.Delete(ctx, a.ID)
			}
		}

		return errors.New("alias not found")
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
