package app

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shimeoki/wp/internal/v2/domain"
)

type DeleteSourceCommand struct {
	sources domain.SourceRepo
}

func NewDeleteSourceCommand(
	sources domain.SourceRepo,
) *DeleteSourceCommand {
	return &DeleteSourceCommand{
		sources: sources,
	}
}

type DeleteSourceData struct {
	ID string
}

type DeleteSourceResult struct{}

func (cmd *DeleteSourceCommand) Execute(
	ctx Ctx,
	data *DeleteSourceData,
) (*DeleteSourceResult, error) {
	id, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, err
	}

	source, _ := cmd.sources.ByID(ctx, domain.ID(id))
	if source == nil {
		return nil, errors.New("source not found")
	}

	if err := cmd.sources.Delete(ctx, source.ID); err != nil {
		return nil, err
	}

	return &DeleteSourceResult{}, nil
}
