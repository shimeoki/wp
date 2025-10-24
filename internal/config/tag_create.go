package config

import "github.com/shimeoki/wp/internal/app"

type CreateTagWorker struct {
	worker *Worker
}

func (w *CreateTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.CreateTagProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.CreateTagProvider{
			TagRepo: p.TagRepo,
		})
	})
}
