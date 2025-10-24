package config

import "github.com/shimeoki/wp/internal/app"

type ListTagsWorker struct {
	worker *Worker
}

func (w *ListTagsWorker) Do(
	ctx app.Ctx,
	fn func(*app.ListTagsProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.ListTagsProvider{
			TagRepo: p.TagRepo,
		})
	})
}
