package app

import "github.com/shimeoki/wp/internal/v2/domain"

type ListSourcesHandler struct {
	sources domain.SourceRepo
}

func NewListSourcesHandler(
	sources domain.SourceRepo,
) *ListSourcesHandler {
	return &ListSourcesHandler{
		sources: sources,
	}
}

type ListSourcesQuery struct{}

type ListSourcesResult struct {
	List []SourceResult
}

func (h *ListSourcesHandler) Handle(
	ctx Ctx,
	qry *ListSourcesQuery,
) (*ListSourcesResult, error) {
	it, err := h.sources.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListSourcesResult{}

	for source := range it {
		res := SourceResult{
			ID:   source.ID.String(),
			Name: source.Name.String(),
			Link: source.Link,
		}

		result.List = append(result.List, res)
	}

	return result, nil
}
