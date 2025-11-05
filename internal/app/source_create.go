package app

import "github.com/shimeoki/wp/internal/domain"

type CreateSourceProvider interface {
	SourceProvider
}

type CreateSourceWorker Worker[CreateSourceProvider]

type CreateSourceHandler struct {
	worker CreateSourceWorker
	logger Logger
}

func NewCreateSourceHandler(
	w CreateSourceWorker,
	l Logger,
) *CreateSourceHandler {
	return &CreateSourceHandler{worker: w, logger: l}
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

	Act(h.logger, ctx, "creating source", cmd)
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
		Fail(h.logger, ctx, "failed to create source", err)
		return nil, err
	}

	return &r, nil
}
