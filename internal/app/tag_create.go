package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateTagProvider interface {
	TagProvider
}

type CreateTagWorker Worker[CreateTagProvider]

type CreateTagHandler struct {
	worker CreateTagWorker
}

func NewCreateTagHandler(w CreateTagWorker) *CreateTagHandler {
	return &CreateTagHandler{worker: w}
}

type CreateTagCommand struct {
	Name string
}

type CreateTagResult struct{}

func (h *CreateTagHandler) Handle(
	ctx Ctx,
	cmd *CreateTagCommand,
) (*CreateTagResult, error) {
	var r CreateTagResult

	if err := h.worker.Work(ctx, func(p CreateTagProvider) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := p.TagRepo().FindByName(ctx, name)
		if t != nil {
			return errors.New("tag already exists")
		}

		tag, err := domain.NewTag(name)
		if err != nil {
			return err
		}

		return p.TagRepo().Save(ctx, tag)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
