package app

import "github.com/shimeoki/wp/internal/domain"

type ListTagsProviders struct {
	TagRepo domain.TagRepo
}

type ListTagsWorker Worker[*ListTagsProviders]

type ListTagsHandler struct {
	worker ListTagsWorker
}

func NewListTagsHandler(w ListTagsWorker) *ListTagsHandler {
	return &ListTagsHandler{worker: w}
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

	if err := h.worker.Do(ctx, func(p *ListTagsProviders) error {
		it, err := p.TagRepo.All(ctx)
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
