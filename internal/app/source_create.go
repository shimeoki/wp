package app

import "github.com/shimeoki/wp/internal/domain"

type CreateSourceProvider interface {
	SourceProvider
}

type CreateSourceWorker Worker[CreateSourceProvider]

type CreateSourceHandler struct {
	worker CreateSourceWorker
}

func NewCreateSourceHandler(w CreateSourceWorker) *CreateSourceHandler {
	return &CreateSourceHandler{worker: w}
}

type CreateSourceCommand struct {
	Name string
	Link *string
}

type CreateSourceResult struct {
	ID string
}

func (h *CreateSourceHandler) Handle(
	ctx Ctx,
	cmd *CreateSourceCommand,
) (*CreateSourceResult, error) {
	var r CreateSourceResult

	if err := h.worker.Work(ctx, func(p CreateSourceProvider) error {
		name, err := domain.ParseName(cmd.Name)
		if err != nil {
			return err
		}

		source, err := domain.NewSource(name, cmd.Link)
		if err != nil {
			return err
		}

		if err := p.SourceRepo().Save(ctx, source); err != nil {
			return err
		}

		r.ID = source.ID.String()

		return nil
	}); err != nil {
		return nil, err
	}

	return &r, nil
}
