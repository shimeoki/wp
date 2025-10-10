package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type CreateTagHandler struct {
	tags domain.TagRepo
}

func NewCreateTagHandler(
	tags domain.TagRepo,
) *CreateTagHandler {
	return &CreateTagHandler{
		tags: tags,
	}
}

type CreateTagCommand struct {
	Name string
}

type CreateTagResult struct{}

func (h *CreateTagHandler) Handle(
	ctx Ctx,
	cmd *CreateTagCommand,
) (*CreateTagResult, error) {
	t, _ := h.tags.FindByName(ctx, cmd.Name)
	if t != nil {
		return nil, errors.New("tag already exists")
	}

	name, err := domain.ParseName(cmd.Name)
	if err != nil {
		return nil, err
	}

	tag, err := domain.NewTag(name)
	if err != nil {
		return nil, err
	}

	if err := h.tags.Save(ctx, tag); err != nil {
		return nil, err
	}

	return &CreateTagResult{}, nil
}
