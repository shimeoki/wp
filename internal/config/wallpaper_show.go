package config

import "github.com/shimeoki/wp/internal/app"

type ShowWallpaperWorker struct {
	worker *Worker
}

func (w *ShowWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.ShowWallpaperProviders) error,
) error {
	return w.worker.Do(ctx, func(p *Providers) error {
		return fn(&app.ShowWallpaperProviders{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
