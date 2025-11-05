package app

import "github.com/shimeoki/wp/internal/domain"

type DeleteSourceProvider interface {
	SourceProvider
}

type DeleteSourceWorker Worker[DeleteSourceProvider]

type DeleteSourceHandler struct {
	worker DeleteSourceWorker
	logger Logger
}

func NewDeleteSourceHandler(
	w DeleteSourceWorker,
	l Logger,
) *DeleteSourceHandler {
	return &DeleteSourceHandler{worker: w, logger: l}
}

type DeleteSourceCommand struct {
	ID string
}

type DeleteSourceResult struct{}

func (h *DeleteSourceHandler) Handle(
	ctx Ctx,
	cmd *DeleteSourceCommand,
) (*DeleteSourceResult, error) {
	var r DeleteSourceResult

	Act(h.logger, ctx, "deleting source", cmd)
	if err := h.worker.Work(ctx, func(p DeleteSourceProvider) error {
		sources := p.SourceRepo()

		id, err := domain.ParseID(cmd.ID)
		if err != nil {
			return err
		}

		source, _ := sources.FindByID(ctx, id)
		if source == nil {
			return domain.NewNotFoundError("source", "id", id.String())
		}

		return sources.Delete(ctx, source.ID)
	}); err != nil {
		Fail(h.logger, ctx, "failed to delete source", err)
		return nil, err
	}

	return &r, nil
}
