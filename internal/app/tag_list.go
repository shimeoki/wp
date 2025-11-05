package app

type ListTagsProvider interface {
	TagProvider
}

type ListTagsWorker Worker[ListTagsProvider]

type ListTagsHandler struct {
	worker ListTagsWorker
	logger Logger
}

func NewListTagsHandler(w ListTagsWorker, l Logger) *ListTagsHandler {
	return &ListTagsHandler{worker: w, logger: l}
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

	Act(h.logger, ctx, "listing tags", qry)
	if err := h.worker.Work(ctx, func(p ListTagsProvider) error {
		it, err := p.TagRepo().All(ctx)
		if err != nil {
			return err
		}

		for tag := range it {
			r.Names = append(r.Names, tag.Name.String())
		}

		return nil
	}); err != nil {
		Fail(h.logger, ctx, "failed to list tags", err)
		return nil, err
	}

	return &r, nil
}
