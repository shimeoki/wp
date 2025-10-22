package config

import "github.com/shimeoki/wp/internal/app"

type Handlers struct {
	CreateWallpaper *app.CreateWallpaperHandler
	DeleteWallpaper *app.DeleteWallpaperHandler
	FindWallpaper   *app.FindWallpaperHandler
	ShowWallpaper   *app.ShowWallpaperHandler

	AddTag    *app.AddTagHandler
	CreateTag *app.CreateTagHandler
	DeleteTag *app.DeleteTagHandler
	ListTags  *app.ListTagsHandler
}

func (p *Provider) Handlers() *Handlers {
	return &Handlers{
		CreateWallpaper: app.NewCreateWallpaperHandler(
			p.CreateWallpaperProvider,
		),

		DeleteWallpaper: app.NewDeleteWallpaperHandler(
			p.DeleteWallpaperProvider,
		),

		FindWallpaper: app.NewFindWallpaperHandler(
			p.FindWallpaperProvider,
		),

		ShowWallpaper: app.NewShowWallpaperHandler(
			p.ShowWallpaperProvider,
		),

		AddTag: app.NewAddTagHandler(
			p.AddTagProvider,
		),

		CreateTag: app.NewCreateTagHandler(
			p.CreateTagProvider,
		),

		DeleteTag: app.NewDeleteTagHandler(
			p.DeleteTagProvider,
		),

		ListTags: app.NewListTagsHandler(
			p.ListTagsProvider,
		),
	}
}
