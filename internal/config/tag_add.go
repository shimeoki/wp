package config

import "github.com/shimeoki/wp/internal/app"

type AddTagWorker struct {
	worker *Worker
}

func (w *AddTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.AddTagProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.AddTagProviders{
			TagRepo:       p.TagRepo,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
