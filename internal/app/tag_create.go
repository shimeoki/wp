package app

import "github.com/shimeoki/wp/internal/domain"

type CreateTagProvider interface {
	TagProvider
}

type CreateTagWorker Worker[CreateTagProvider]

type CreateTagHandler struct {
	worker CreateTagWorker
	logger Logger
}

func NewCreateTagHandler(w CreateTagWorker, l Logger) *CreateTagHandler {
	return &CreateTagHandler{worker: w, logger: l}
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

	Act(h.logger, ctx, "creating tag", cmd)
	if err := h.worker.Work(ctx, func(p CreateTagProvider) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := p.TagRepo().FindByName(ctx, name)
		if t != nil {
			return domain.NewAlreadyExistsError("tag", "name", name.String())
		}

		tag, err := domain.NewTag(name)
		if err != nil {
			return err
		}

		return p.TagRepo().Save(ctx, tag)
	}); err != nil {
		Fail(h.logger, ctx, "failed to create tag", err)
		return nil, err
	}

	return &r, nil
}
