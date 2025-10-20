package config

import "github.com/shimeoki/wp/internal/app"

type DeleteWallpaperProvider struct {
	provider *Provider
}

func (p *DeleteWallpaperProvider) Provide(
	ctx app.Ctx,
) (app.DeleteWallpaperWorker, error) {
	return p.provider.Provide(ctx)
}
