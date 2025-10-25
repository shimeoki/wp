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
				func(ctx app.Ctx, fn func(app.CreateWallpaperProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		DeleteWallpaper: app.NewDeleteWallpaperHandler(
			app.WorkerFunc[app.DeleteWallpaperProvider](
				func(ctx app.Ctx, fn func(app.DeleteWallpaperProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		FindWallpaper: app.NewFindWallpaperHandler(
			app.WorkerFunc[app.FindWallpaperProvider](
				func(ctx app.Ctx, fn func(app.FindWallpaperProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		ShowWallpaper: app.NewShowWallpaperHandler(
			app.WorkerFunc[app.ShowWallpaperProvider](
				func(ctx app.Ctx, fn func(app.ShowWallpaperProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		AddTag: app.NewAddTagHandler(
			app.WorkerFunc[app.AddTagProvider](
				func(ctx app.Ctx, fn func(app.AddTagProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		CreateTag: app.NewCreateTagHandler(
			app.WorkerFunc[app.CreateTagProvider](
				func(ctx app.Ctx, fn func(app.CreateTagProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		DeleteTag: app.NewDeleteTagHandler(
			app.WorkerFunc[app.DeleteTagProvider](
				func(ctx app.Ctx, fn func(app.DeleteTagProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),

		ListTags: app.NewListTagsHandler(
			app.WorkerFunc[app.ListTagsProvider](
				func(ctx app.Ctx, fn func(app.ListTagsProvider) error) error {
					return w.Do(ctx, func(p *Provider) error { return fn(p) })
				},
			),
		),
	}
}
