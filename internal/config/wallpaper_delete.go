package config

import "github.com/shimeoki/wp/internal/app"

type DeleteWallpaperWorker struct {
	worker *Worker
}

func (w *DeleteWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.DeleteWallpaperProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.DeleteWallpaperProvider{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
