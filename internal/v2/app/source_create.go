package app

import "github.com/shimeoki/wp/internal/v2/domain"

type CreateSourceHandler struct {
	sources domain.SourceRepo
}

func NewCreateSourceHandler(
	sources domain.SourceRepo,
) *CreateSourceHandler {
	return &CreateSourceHandler{
		sources: sources,
	}
}

type CreateSourceCommand struct {
	Name string
	Link *string
}

type CreateSourceResult struct {
	ID string
}

func (h *CreateSourceHandler) Execute(
	ctx Ctx,
	cmd *CreateSourceCommand,
) (*CreateSourceResult, error) {
	source := domain.NewSource(cmd.Name, cmd.Link)

	if err := h.sources.Save(ctx, source); err != nil {
		return nil, err
	}

	return &CreateSourceResult{ID: source.ID.String()}, nil
}
