package app

import "github.com/shimeoki/wp/internal/domain"

// Defines the environment that's required for an action to be handled.
// This acts as a base interface for all other providers, which should be used
// to compose a specific provider for an action.
type Provider any

// keep-sorted start block=yes newline_separated=yes skip_lines=1

type AliasProvider interface {
	Provider
	AliasRepo() domain.AliasRepo
}

type SourceProvider interface {
	Provider
	SourceRepo() domain.SourceRepo
}

type StoreProvider interface {
	Provider
	Store() domain.Store
}

type TagProvider interface {
	Provider
	TagRepo() domain.TagRepo
}

type WallpaperProvider interface {
	Provider
	WallpaperRepo() domain.WallpaperRepo
}

// keep-sorted end
