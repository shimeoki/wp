package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) ListTagsProvider(
	ctx app.Ctx,
	fn func(*app.ListTagsWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.ListTagsWorker{
			TagRepo: w.TagRepo,
		})
	})
}
