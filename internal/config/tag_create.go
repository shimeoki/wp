package config

import "github.com/shimeoki/wp/internal/app"

type CreateTagProvider struct {
	provider *Provider
}

func (p *CreateTagProvider) Provide(ctx app.Ctx) (app.CreateTagWorker, error) {
	return p.provider.Provide(ctx)
}
