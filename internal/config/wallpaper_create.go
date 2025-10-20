package config

import "github.com/shimeoki/wp/internal/app"

type CreateWallpaperProvider struct {
	provider *Provider
}

func (p *CreateWallpaperProvider) Provide(
	ctx app.Ctx,
) (app.CreateWallpaperWorker, error) {
	return p.provider.Provide(ctx)
}
