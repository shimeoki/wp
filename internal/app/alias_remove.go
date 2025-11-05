package app

import "github.com/shimeoki/wp/internal/domain"

type RemoveAliasProvider interface {
	AliasProvider
	WallpaperProvider
}

type RemoveAliasWorker Worker[RemoveAliasProvider]

type RemoveAliasHandler struct {
	worker RemoveAliasWorker
	logger Logger
}

func NewRemoveAliasHandler(w RemoveAliasWorker, l Logger) *RemoveAliasHandler {
	return &RemoveAliasHandler{worker: w, logger: l}
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

	Act(h.logger, ctx, "removing alias", cmd)
	if err := h.worker.Work(ctx, func(p RemoveAliasProvider) error {
		aliases, wallpapers := p.AliasRepo(), p.WallpaperRepo()

		hash, err := domain.ParseHash(cmd.WallpaperHash)
		if err != nil {
			return err
		}

		w, _ := wallpapers.FindByHash(ctx, hash)
		if w == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
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

		return domain.NewNotFoundError("alias", "name", name.String())
	}); err != nil {
		Fail(h.logger, ctx, "failed to remove alias", err)
		return nil, err
	}

	return &r, nil
}
