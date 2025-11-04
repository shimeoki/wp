package app

import "github.com/shimeoki/wp/internal/domain"

type AddAliasProvider interface {
	AliasProvider
	WallpaperProvider
}

type AddAliasWorker Worker[AddAliasProvider]

type AddAliasHandler struct {
	worker AddAliasWorker
	logger *logger
}

func NewAddAliasHandler(w AddAliasWorker, l Logger) *AddAliasHandler {
	return &AddAliasHandler{worker: w, logger: wrapLogger(l)}
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
		h.logger.Info(ctx, "adding alias", "command", cmd)
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

		a, err := domain.NewAlias(name, w.ID)
		if err != nil {
			return err
		}

		h.logger.Info(ctx, "saving alias", "entity", a)
		return aliases.Save(ctx, a)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
