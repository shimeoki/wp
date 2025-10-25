package app

import "github.com/shimeoki/wp/internal/domain"

type Provider any

type WallpaperProvider interface {
	Provider
	WallpaperRepo() domain.WallpaperRepo
}

type TagProvider interface {
	Provider
	TagRepo() domain.TagRepo
}

type StoreProvider interface {
	Provider
	Store() domain.Store
}
