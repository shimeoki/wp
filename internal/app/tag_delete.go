package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteTagProvider struct {
	TagRepo domain.TagRepo
}

type DeleteTagWorker Worker[*DeleteTagProvider]

type DeleteTagHandler struct {
	worker DeleteTagWorker
}

func NewDeleteTagHandler(w DeleteTagWorker) *DeleteTagHandler {
	return &DeleteTagHandler{worker: w}
}

type DeleteTagCommand struct {
	Name string
}

type DeleteTagResult struct{}

func (h *DeleteTagHandler) Handle(
	ctx Ctx,
	cmd *DeleteTagCommand,
) (*DeleteTagResult, error) {
	var r DeleteTagResult

	if err := h.worker.Do(ctx, func(p *DeleteTagProvider) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := p.TagRepo.FindByName(ctx, name)
		if t == nil {
			return errors.New("tag not found")
		}

		return p.TagRepo.Delete(ctx, t.ID)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
