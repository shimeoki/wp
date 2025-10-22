package config

import "github.com/shimeoki/wp/internal/app"

type CreateTagWorker struct {
	worker *Worker
}

func (w *CreateTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.CreateTagProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.CreateTagProviders{
			TagRepo: p.TagRepo,
		})
	})
}
