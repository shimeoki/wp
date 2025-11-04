package app

import "fmt"

type ListAliasesProvider interface {
	AliasProvider
	WallpaperProvider
}

type ListAliasesWorker Worker[ListAliasesProvider]

type ListAliasesHandler struct {
	worker ListAliasesWorker
}

func NewListAliasesHandler(w ListAliasesWorker) *ListAliasesHandler {
	return &ListAliasesHandler{worker: w}
}

type ListAliasesQuery struct{}

type ListAliasesResult struct {
	Aliases []string
}

func (h *ListAliasesHandler) Handle(
	ctx Ctx,
	qry *ListAliasesQuery,
) (*ListAliasesResult, error) {
	var r ListAliasesResult

	if err := h.worker.Work(ctx, func(p ListAliasesProvider) error {
		it, err := p.AliasRepo().All(ctx)
		if err != nil {
			return err
		}

		wallpapers := p.WallpaperRepo()

		for a := range it {
			w, err := wallpapers.FindByID(ctx, a.WallpaperID)
			if err != nil {
				return err
			}

			s := fmt.Sprintf("%s -> %s", a.Name.String(), w.Hash.String())
			r.Aliases = append(r.Aliases, s)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
