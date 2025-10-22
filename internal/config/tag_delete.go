package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) DeleteTagProvider(
	ctx app.Ctx,
	fn func(*app.DeleteTagWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.DeleteTagWorker{
			TagRepo: w.TagRepo,
		})
	})
}
