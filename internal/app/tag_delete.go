package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteTagWorker interface {
	Worker
	TagRepo() domain.TagRepo
}

type DeleteTagProvider interface {
	Provider[DeleteTagWorker]
}

type DeleteTagHandler struct {
	provider DeleteTagProvider
}

func NewDeleteTagHandler(
	p DeleteTagProvider,
) *DeleteTagHandler {
	return &DeleteTagHandler{provider: p}
}

type DeleteTagCommand struct {
	Name string
}

type DeleteTagResult struct{}

func (h *DeleteTagHandler) Handle(
	ctx Ctx,
	cmd *DeleteTagCommand,
) (*DeleteTagResult, error) {
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
	if t == nil {
		return nil, errors.New("tag not found")
	}

	if err := repo.Delete(ctx, t.ID); err != nil {
		return nil, err
	}

	return &DeleteTagResult{}, nil
}
