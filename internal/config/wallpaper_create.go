package config

import "github.com/shimeoki/wp/internal/app"

type CreateWallpaperWorker struct {
	worker *Worker
}

func (w *CreateWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.CreateWallpaperProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.CreateWallpaperProviders{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
