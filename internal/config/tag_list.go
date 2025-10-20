package config

import "github.com/shimeoki/wp/internal/app"

type ListTagsProvider struct {
	provider *Provider
}

func (p *ListTagsProvider) Provide(ctx app.Ctx) (app.ListTagsWorker, error) {
	return p.provider.Provide(ctx)
}
