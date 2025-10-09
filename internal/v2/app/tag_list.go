package app

import "github.com/shimeoki/wp/internal/v2/domain"

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
	Map map[string]TagResult
}

func (h *ListTagsHandler) Handle(
	ctx Ctx,
	qry *ListTagsQuery,
) (*ListTagsResult, error) {
	it, err := h.tags.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListTagsResult{Map: make(map[string]TagResult)}

	for tag := range it {
		result.Map[tag.Name] = TagResult{Name: tag.Name}
	}

	return result, nil
}
