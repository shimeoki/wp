package config

import "github.com/shimeoki/wp/internal/app"

type ShowWallpaperWorker struct {
	worker *Worker
}

func (w *ShowWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.ShowWallpaperProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.ShowWallpaperProvider{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
