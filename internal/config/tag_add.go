package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) AddTagProvider(
	ctx app.Ctx,
	fn func(*app.AddTagWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.AddTagWorker{
			TagRepo:       w.TagRepo,
			WallpaperRepo: w.WallpaperRepo,
		})
	})
}
