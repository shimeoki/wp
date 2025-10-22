package config

import "github.com/shimeoki/wp/internal/app"

func (p *Provider) ShowWallpaperProvider(
	ctx app.Ctx,
	fn func(*app.ShowWallpaperWorker) error,
) error {
	return p.Provider(ctx, func(w *Worker) error {
		return fn(&app.ShowWallpaperWorker{
			Store:         w.Store,
			WallpaperRepo: w.WallpaperRepo,
		})
	})
}
