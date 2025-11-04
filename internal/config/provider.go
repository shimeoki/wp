package config

import "github.com/shimeoki/wp/internal/domain"

type Provider struct {
	store      domain.Store
	tags       domain.TagRepo
	wallpapers domain.WallpaperRepo
	sources    domain.SourceRepo
	aliases    domain.AliasRepo
}

func (p *Provider) Store() domain.Store {
	return p.store
}

func (p *Provider) TagRepo() domain.TagRepo {
	return p.tags
}

func (p *Provider) WallpaperRepo() domain.WallpaperRepo {
	return p.wallpapers
}

func (p *Provider) SourceRepo() domain.SourceRepo {
	return p.sources
}

func (p *Provider) AliasRepo() domain.AliasRepo {
	return p.aliases
}
