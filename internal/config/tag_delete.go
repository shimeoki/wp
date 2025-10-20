package config

import "github.com/shimeoki/wp/internal/app"

type DeleteTagProvider struct {
	provider *Provider
}

func (p *DeleteTagProvider) Provide(ctx app.Ctx) (app.DeleteTagWorker, error) {
	return p.provider.Provide(ctx)
}
