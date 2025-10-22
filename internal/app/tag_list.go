package app

import "github.com/shimeoki/wp/internal/domain"

type ListTagsWorker struct {
	TagRepo domain.TagRepo
}

type ListTagsProvider Provider[*ListTagsWorker]

type ListTagsHandler struct {
	provider ListTagsProvider
}

func NewListTagsHandler(
	p ListTagsProvider,
) *ListTagsHandler {
	return &ListTagsHandler{provider: p}
}

type ListTagsQuery struct{}

type ListTagsResult struct {
	Names []string
}

func (h *ListTagsHandler) Handle(
	ctx Ctx,
	qry *ListTagsQuery,
) (*ListTagsResult, error) {
	var r ListTagsResult

	if err := h.provider(ctx, func(w *ListTagsWorker) error {
		it, err := w.TagRepo.All(ctx)
		if err != nil {
			return err
		}

		for tag := range it {
			r.Names = append(r.Names, tag.Name.String())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
