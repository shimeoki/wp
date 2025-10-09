package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type DeleteTagCommand struct {
	tags domain.TagRepo
}

func NewDeleteTagCommand(
	tags domain.TagRepo,
) *DeleteTagCommand {
	return &DeleteTagCommand{
		tags: tags,
	}
}

type DeleteTagData struct {
	Name string
}

type DeleteTagResult struct{}

func (cmd *DeleteTagCommand) Execute(
	ctx Ctx,
	data *DeleteTagData,
) (*DeleteTagResult, error) {
	t, _ := cmd.tags.ByName(ctx, data.Name)
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if err := cmd.tags.Delete(ctx, t.ID); err != nil {
		return nil, err
	}

	return &DeleteTagResult{}, nil
}
