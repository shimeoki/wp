package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteSourceProvider interface {
	SourceProvider
}

type DeleteSourceWorker Worker[DeleteSourceProvider]

type DeleteSourceHandler struct {
	worker DeleteSourceWorker
}

func NewDeleteSourceHandler(w DeleteSourceWorker) *DeleteSourceHandler {
	return &DeleteSourceHandler{worker: w}
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

	if err := h.worker.Work(ctx, func(p DeleteSourceProvider) error {
		sources := p.SourceRepo()

		id, err := domain.ParseID(cmd.ID)
		if err != nil {
			return err
		}

		source, _ := sources.FindByID(ctx, id)
		if source == nil {
			return errors.New("source not found")
		}

		return sources.Delete(ctx, source.ID)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
