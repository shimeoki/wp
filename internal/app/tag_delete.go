package app

import (
	"errors"

	"github.com/shimeoki/wp/internal/domain"
)

type DeleteTagWorker struct {
	TagRepo domain.TagRepo
}

type DeleteTagProvider Provider[*DeleteTagWorker]

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
	var r DeleteTagResult

	if err := h.provider(ctx, func(w *DeleteTagWorker) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		t, _ := w.TagRepo.FindByName(ctx, name)
		if t == nil {
			return errors.New("tag not found")
		}

		return w.TagRepo.Delete(ctx, t.ID)
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
