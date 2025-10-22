package config

import "github.com/shimeoki/wp/internal/app"

type FindWallpaperWorker struct {
	worker *Worker
}

func (w *FindWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.FindWallpaperProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.FindWallpaperProviders{
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
