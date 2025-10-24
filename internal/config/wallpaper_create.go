package config

import "github.com/shimeoki/wp/internal/app"

type CreateWallpaperWorker struct {
	worker *Worker
}

func (w *CreateWallpaperWorker) Do(
	ctx app.Ctx,
	fn func(*app.CreateWallpaperProvider) error,
) error {
	return w.worker.Do(ctx, func(p *Provider) error {
		return fn(&app.CreateWallpaperProvider{
			Store:         p.Store,
			WallpaperRepo: p.WallpaperRepo,
		})
	})
}
