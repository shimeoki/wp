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
	List []struct {
		ID   string
		Name string
		Link *string
	}
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
		result.List = append(result.List, struct {
			ID   string
			Name string
			Link *string
		}{
			ID:   source.ID.String(),
			Name: source.Name.String(),
			Link: source.Link,
		})
	}

	return result, nil
}
