package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) CreateTagProvider(
	ctx app.Ctx,
	fn func(*app.CreateTagWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.CreateTagWorker{
			TagRepo: w.TagRepo,
		})
	})
}
