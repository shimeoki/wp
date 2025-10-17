package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteTagHandler struct {
	tags domain.TagRepo
}

func NewDeleteTagHandler(
	tags domain.TagRepo,
) *DeleteTagHandler {
	return &DeleteTagHandler{
		tags: tags,
	}
}

type DeleteTagCommand struct {
	Name string
}

type DeleteTagResult struct{}

func (h *DeleteTagHandler) Handle(
	ctx Ctx,
	cmd *DeleteTagCommand,
) (*DeleteTagResult, error) {
	name, err := domain.ParseName(cmd.Name)
	if err != nil {
		return nil, err
	}

	t, _ := h.tags.FindByName(ctx, name)
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if err := h.tags.Delete(ctx, t.ID); err != nil {
		return nil, err
	}

	return &DeleteTagResult{}, nil
}
