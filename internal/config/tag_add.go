package config

import "github.com/shimeoki/wp/internal/app"

type AddTagWorker struct {
	worker *Worker
}

func (w *AddTagWorker) Do(
	ctx app.Ctx,
	fn func(*app.AddTagProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.AddTagProvider{
			TagRepo:       p.TagRepo,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
