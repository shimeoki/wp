package app

import "github.com/shimeoki/wp/internal/v2/domain"

type CreateSourceCommand struct {
	sources domain.SourceRepo
}

func NewCreateSourceCommand(
	sources domain.SourceRepo,
) *CreateSourceCommand {
	return &CreateSourceCommand{
		sources: sources,
	}
}

type CreateSourceData struct {
	Name string
	Link *string
}

type CreateSourceResult struct {
	ID string
}

func (cmd *CreateSourceCommand) Execute(
	ctx Ctx,
	data *CreateSourceData,
) (*CreateSourceResult, error) {
	source := domain.NewSource(data.Name, data.Link)

	if err := cmd.sources.Save(ctx, source); err != nil {
		return nil, err
	}

	return &CreateSourceResult{ID: source.ID.String()}, nil
}
