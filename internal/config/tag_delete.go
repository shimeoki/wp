package config

import "github.com/shimeoki/wp/internal/app"

type DeleteTagWorker struct {
	worker *Worker
}

func (w *DeleteTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.DeleteTagProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.DeleteTagProviders{
			TagRepo: p.TagRepo,
		})
	})
}
