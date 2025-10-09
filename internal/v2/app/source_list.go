package app

import "github.com/shimeoki/wp/internal/v2/domain"

type ListSourcesQuery struct {
	sources domain.SourceRepo
}

func NewListSourcesQuery(
	sources domain.SourceRepo,
) *ListSourcesQuery {
	return &ListSourcesQuery{
		sources: sources,
	}
}

type ListSourcesData struct{}

type ListSourcesResult struct {
	List []SourceResult
}

func (qry *ListSourcesQuery) Execute(
	ctx Ctx,
	data *ListSourcesData,
) (*ListSourcesResult, error) {
	it, err := qry.sources.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListSourcesResult{}

	for source := range it {
		res := SourceResult{
			ID:   source.ID.String(),
			Name: source.Name,
			Link: source.Link,
		}

		result.List = append(result.List, res)
	}

	return result, nil
}
