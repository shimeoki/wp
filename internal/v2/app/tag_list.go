package app

import "github.com/shimeoki/wp/internal/v2/domain"

type ListTagsQuery struct {
	tags domain.TagRepo
}

func NewListTagsQuery(
	tags domain.TagRepo,
) *ListTagsQuery {
	return &ListTagsQuery{
		tags: tags,
	}
}

type ListTagsData struct{}

type ListTagsResult struct {
	Map map[string]TagResult
}

func (qry *ListTagsQuery) Execute(
	ctx Ctx,
	data *ListTagsData,
) (*ListTagsResult, error) {
	it, err := qry.tags.All(ctx)
	if err != nil {
		return nil, err
	}

	result := &ListTagsResult{Map: make(map[string]TagResult)}

	for tag := range it {
		result.Map[tag.Name] = TagResult{Name: tag.Name}
	}

	return result, nil
}
