package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) FindWallpaperProvider(
	ctx app.Ctx,
	fn func(*app.FindWallpaperWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.FindWallpaperWorker{
			WallpaperRepo: w.WallpaperRepo,
		})
	})
}
