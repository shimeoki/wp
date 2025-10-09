package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/v2/domain"
)

type CreateTagCommand struct {
	tags domain.TagRepo
}

func NewCreateTagCommand(
	tags domain.TagRepo,
) *CreateTagCommand {
	return &CreateTagCommand{
		tags: tags,
	}
}

type CreateTagData struct {
	Name string
}

type CreateTagResult struct{}

func (cmd *CreateTagCommand) Execute(
	ctx Ctx,
	data *CreateTagData,
) (*CreateTagResult, error) {
	t, _ := cmd.tags.ByName(ctx, data.Name)
	if t != nil {
		return nil, errors.New("tag already exists")
	}

	tag := domain.NewTag(data.Name)

	if err := cmd.tags.Save(ctx, tag); err != nil {
		return nil, err
	}

	return &CreateTagResult{}, nil
}
