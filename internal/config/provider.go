package config

import "github.com/shimeoki/wp/internal/domain"

type Provider struct {
	store      domain.Store
	tags       domain.TagRepo
	wallpapers domain.WallpaperRepo
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
