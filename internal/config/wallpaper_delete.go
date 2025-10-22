package config

import "github.com/shimeoki/wp/internal/app"

type DeleteWallpaperWorker struct {
	worker *Worker
}

func (w *DeleteWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.DeleteWallpaperProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.DeleteWallpaperProviders{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
