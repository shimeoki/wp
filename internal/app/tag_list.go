package app

import "github.com/shimeoki/wp/internal/domain"

type ListTagsWorker interface {
	Worker
	TagRepo() domain.TagRepo
}

type ListTagsProvider interface {
	Provider[ListTagsWorker]
}

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
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	defer worker.Rollback()
	repo := worker.TagRepo()

	it, err := repo.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListTagsResult{}

	for tag := range it {
		result.Names = append(result.Names, tag.Name.String())
	}

	if err := worker.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
