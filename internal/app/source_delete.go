package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteSourceHandler struct {
	sources domain.SourceRepo
}

func NewDeleteSourceHandler(
	sources domain.SourceRepo,
) *DeleteSourceHandler {
	return &DeleteSourceHandler{
		sources: sources,
	}
}

type DeleteSourceCommand struct {
	ID string
}

type DeleteSourceResult struct{}

func (h *DeleteSourceHandler) Handle(
	ctx Ctx,
	cmd *DeleteSourceCommand,
) (*DeleteSourceResult, error) {
	id, err := domain.ParseID(cmd.ID)
	if err != nil {
		return nil, err
	}

	source, _ := h.sources.FindByID(ctx, id)
	if source == nil {
		return nil, errors.New("source not found")
	}

	if err := h.sources.Delete(ctx, source.ID); err != nil {
		return nil, err
	}

	return &DeleteSourceResult{}, nil
}
