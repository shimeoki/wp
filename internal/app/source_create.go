package app

import "github.com/shimeoki/wp/internal/domain"

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

func (h *CreateSourceHandler) Handle(
	ctx Ctx,
	cmd *CreateSourceCommand,
) (*CreateSourceResult, error) {
	name, err := domain.ParseName(cmd.Name)
	if err != nil {
		return nil, err
	}

	source, err := domain.NewSource(name, cmd.Link)
	if err != nil {
		return nil, err
	}

	if err := h.sources.Save(ctx, source); err != nil {
		return nil, err
	}

	return &CreateSourceResult{ID: source.ID.String()}, nil
}
