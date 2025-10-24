package config

import "github.com/shimeoki/wp/internal/app"

type FindWallpaperWorker struct {
	worker *Worker
}

func (w *FindWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.FindWallpaperProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.FindWallpaperProvider{
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
