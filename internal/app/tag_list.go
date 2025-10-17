package app

import "github.com/shimeoki/wp/internal/domain"

type ListTagsHandler struct {
	tags domain.TagRepo
}

func NewListTagsHandler(
	tags domain.TagRepo,
) *ListTagsHandler {
	return &ListTagsHandler{
		tags: tags,
	}
}

type ListTagsQuery struct{}

type ListTagsResult struct {
	Names []string
}

func (h *ListTagsHandler) Handle(
	ctx Ctx,
	qry *ListTagsQuery,
) (*ListTagsResult, error) {
	it, err := h.tags.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListTagsResult{}

	for tag := range it {
		result.Names = append(result.Names, tag.Name.String())
	}

	return result, nil
}
