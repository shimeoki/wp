package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type AddAliasProvider interface {
	AliasProvider
	WallpaperProvider
}

type AddAliasWorker Worker[AddAliasProvider]

type AddAliasHandler struct {
	worker AddAliasWorker
}

func NewAddAliasHandler(w AddAliasWorker) *AddAliasHandler {
	return &AddAliasHandler{worker: w}
}

type AddAliasCommand struct {
	WallpaperHash string
	AliasName     string
}

type AddAliasResult struct{}

func (h *AddAliasHandler) Handle(
	ctx Ctx,
	cmd *AddAliasCommand,
) (*AddAliasResult, error) {
	var r AddAliasResult

	if err := h.worker.Work(ctx, func(p AddAliasProvider) error {
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

		a, err := domain.NewAlias(name, w.ID)
		if err != nil {
			return err
		}

		return aliases.Save(ctx, a)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
