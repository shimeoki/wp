package app

import "github.com/shimeoki/wp/internal/domain"

// Defines the environment that's required for an action to be handled.
// This acts as a base interface for all other providers, which should be used
// to compose a specific provider for an action.
type Provider any

type WallpaperProvider interface {
	Provider
	WallpaperRepo() domain.WallpaperRepo
}

type TagProvider interface {
	Provider
	TagRepo() domain.TagRepo
}

type SourceProvider interface {
	Provider
	SourceRepo() domain.SourceRepo
}

type StoreProvider interface {
	Provider
	Store() domain.Store
}

type AliasProvider interface {
	Provider
	AliasRepo() domain.AliasRepo
}
