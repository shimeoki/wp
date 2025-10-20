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
			&CreateWallpaperProvider{provider: p},
		),

		DeleteWallpaper: app.NewDeleteWallpaperHandler(
			&DeleteWallpaperProvider{provider: p},
		),

		FindWallpaper: app.NewFindWallpaperHandler(
			&FindWallpaperProvider{provider: p},
		),

		ShowWallpaper: app.NewShowWallpaperHandler(
			&ShowWallpaperProvider{provider: p},
		),

		AddTag: app.NewAddTagHandler(
			&AddTagProvider{provider: p},
		),

		CreateTag: app.NewCreateTagHandler(
			&CreateTagProvider{provider: p},
		),

		DeleteTag: app.NewDeleteTagHandler(
			&DeleteTagProvider{provider: p},
		),

		ListTags: app.NewListTagsHandler(
			&ListTagsProvider{provider: p},
		),
	}
}
