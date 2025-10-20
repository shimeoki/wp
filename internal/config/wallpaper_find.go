package config

import "github.com/shimeoki/wp/internal/app"

type FindWallpaperProvider struct {
	provider *Provider
}

func (p *FindWallpaperProvider) Provide(
	ctx app.Ctx,
) (app.FindWallpaperWorker, error) {
	return p.provider.Provide(ctx)
}
