package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateTagWorker interface {
	Worker
	TagRepo() domain.TagRepo
}

type CreateTagProvider interface {
	Provider[CreateTagWorker]
}

type CreateTagHandler struct {
	provider CreateTagProvider
}

func NewCreateTagHandler(
	p CreateTagProvider,
) *CreateTagHandler {
	return &CreateTagHandler{provider: p}
}

type CreateTagCommand struct {
	Name string
}

type CreateTagResult struct{}

func (h *CreateTagHandler) Handle(
	ctx Ctx,
	cmd *CreateTagCommand,
) (*CreateTagResult, error) {
	worker, err := h.provider.Provide(ctx)
	if err != nil {
		return nil, err
	}

	repo := worker.TagRepo()

	name, err := domain.ParseName(cmd.Name)
	if err != nil {
		return nil, err
	}

	t, _ := repo.FindByName(ctx, name)
	if t != nil {
		return nil, errors.New("tag already exists")
	}

	tag, err := domain.NewTag(name)
	if err != nil {
		return nil, err
	}

	if err := repo.Save(ctx, tag); err != nil {
		return nil, err
	}

	return &CreateTagResult{}, nil
}
