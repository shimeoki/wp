package config

import "github.com/shimeoki/wp/internal/app"

type DeleteTagWorker struct {
	worker *Worker
}

func (w *DeleteTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.DeleteTagProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.DeleteTagProvider{
			TagRepo: p.TagRepo,
		})
	})
}
