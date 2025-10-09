package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
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
	t, _ := h.tags.ByName(ctx, cmd.Name)
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if err := h.tags.Delete(ctx, t.ID); err != nil {
		return nil, err
	}

	return &DeleteTagResult{}, nil
}
