package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type CreateTagWorker struct {
	TagRepo domain.TagRepo
}

type CreateTagProvider Provider[*CreateTagWorker]

type CreateTagHandler struct {
	provider CreateTagProvider
}

func NewCreateTagHandler(p CreateTagProvider) *CreateTagHandler {
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
	var r CreateTagResult

	if err := h.provider(ctx, func(w *CreateTagWorker) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := w.TagRepo.FindByName(ctx, name)
		if t != nil {
			return errors.New("tag already exists")
		}

		tag, err := domain.NewTag(name)
		if err != nil {
			return err
		}

		return w.TagRepo.Save(ctx, tag)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
