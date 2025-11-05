package app

import "github.com/shimeoki/wp/internal/domain"

type DeleteTagProvider interface {
	TagProvider
}

type DeleteTagWorker Worker[DeleteTagProvider]

type DeleteTagHandler struct {
	worker DeleteTagWorker
	logger Logger
}

func NewDeleteTagHandler(w DeleteTagWorker, l Logger) *DeleteTagHandler {
	return &DeleteTagHandler{worker: w, logger: l}
}

type DeleteTagCommand struct {
	Name string
}

type DeleteTagResult struct{}

func (h *DeleteTagHandler) Handle(
	ctx Ctx,
	cmd *DeleteTagCommand,
) (*DeleteTagResult, error) {
	var r DeleteTagResult

	Act(h.logger, ctx, "deleting tag", cmd)
	if err := h.worker.Work(ctx, func(p DeleteTagProvider) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := p.TagRepo().FindByName(ctx, name)
		if t == nil {
			return domain.NewNotFoundError("tag", "name", name.String())
		}

		return p.TagRepo().Delete(ctx, t.ID)
	}); err != nil {
		Fail(h.logger, ctx, "failed to delete tag", err)
		return nil, err
	}

	return &r, nil
}
