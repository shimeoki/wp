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
			app.WorkerFunc[app.CreateWallpaperProvider](
				func(ctx app.Ctx, j app.Job[app.CreateWallpaperProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		DeleteWallpaper: app.NewDeleteWallpaperHandler(
			app.WorkerFunc[app.DeleteWallpaperProvider](
				func(ctx app.Ctx, j app.Job[app.DeleteWallpaperProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		FindWallpaper: app.NewFindWallpaperHandler(
			app.WorkerFunc[app.FindWallpaperProvider](
				func(ctx app.Ctx, j app.Job[app.FindWallpaperProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		ShowWallpaper: app.NewShowWallpaperHandler(
			app.WorkerFunc[app.ShowWallpaperProvider](
				func(ctx app.Ctx, j app.Job[app.ShowWallpaperProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		AddTag: app.NewAddTagHandler(
			app.WorkerFunc[app.AddTagProvider](
				func(ctx app.Ctx, j app.Job[app.AddTagProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		CreateTag: app.NewCreateTagHandler(
			app.WorkerFunc[app.CreateTagProvider](
				func(ctx app.Ctx, j app.Job[app.CreateTagProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		DeleteTag: app.NewDeleteTagHandler(
			app.WorkerFunc[app.DeleteTagProvider](
				func(ctx app.Ctx, j app.Job[app.DeleteTagProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),

		ListTags: app.NewListTagsHandler(
			app.WorkerFunc[app.ListTagsProvider](
				func(ctx app.Ctx, j app.Job[app.ListTagsProvider]) error {
					return w.Work(ctx, func(p *Provider) error { return j(p) })
				},
			),
		),
	}
}
