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

func (w *Worker) Handlers() *Handlers {
	return &Handlers{
		CreateWallpaper: app.NewCreateWallpaperHandler(
			&CreateWallpaperWorker{worker: w},
		),

		DeleteWallpaper: app.NewDeleteWallpaperHandler(
			&DeleteWallpaperWorker{worker: w},
		),

		FindWallpaper: app.NewFindWallpaperHandler(
			&FindWallpaperWorker{worker: w},
		),

		ShowWallpaper: app.NewShowWallpaperHandler(
			&ShowWallpaperWorker{worker: w},
		),

		AddTag: app.NewAddTagHandler(
			&AddTagWorker{worker: w},
		),

		CreateTag: app.NewCreateTagHandler(
			&CreateTagWorker{worker: w},
		),

		DeleteTag: app.NewDeleteTagHandler(
			&DeleteTagWorker{worker: w},
		),

		ListTags: app.NewListTagsHandler(
			&ListTagsWorker{worker: w},
		),
	}
}
