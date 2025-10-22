package config

import "github.com/shimeoki/wp/internal/app"

type ListTagsWorker struct {
	worker *Worker
}

func (w *ListTagsWorker) Do(
	ctx app.Ctx,
	fn func(*app.ListTagsProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.ListTagsProviders{
			TagRepo: p.TagRepo,
		})
	})
}
