package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) CreateWallpaperProvider(
	ctx app.Ctx,
	fn func(*app.CreateWallpaperWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.CreateWallpaperWorker{
			Store:         w.Store,
			WallpaperRepo: w.WallpaperRepo,
		})
	})
}
