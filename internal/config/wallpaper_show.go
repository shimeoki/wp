package config

import "github.com/shimeoki/wp/internal/app"

type ShowWallpaperProvider struct {
	provider *Provider
}

func (p *ShowWallpaperProvider) Provide(
	ctx app.Ctx,
) (app.ShowWallpaperWorker, error) {
	return p.provider.Provide(ctx)
}
