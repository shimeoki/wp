package app

type ListSourcesProvider interface {
	SourceProvider
}

type ListSourcesWorker Worker[ListSourcesProvider]

type ListSourcesHandler struct {
	worker ListSourcesWorker
}

func NewListSourcesHandler(w ListSourcesWorker) *ListSourcesHandler {
	return &ListSourcesHandler{worker: w}
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
	var r ListSourcesResult

	if err := h.worker.Work(ctx, func(p ListSourcesProvider) error {
		it, err := p.SourceRepo().All(ctx)
		if err != nil {
			return err
		}

		for source := range it {
			r.List = append(r.List, struct {
				ID   string
				Name string
				Link *string
			}{
				ID:   source.ID.String(),
				Name: source.Name.String(),
				Link: source.Link,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
