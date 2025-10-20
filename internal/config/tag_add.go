package config

import "github.com/shimeoki/wp/internal/app"

type AddTagProvider struct {
	provider *Provider
}

func (p *AddTagProvider) Provide(ctx app.Ctx) (app.AddTagWorker, error) {
	return p.provider.Provide(ctx)
}
