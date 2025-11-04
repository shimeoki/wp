package config

import "github.com/shimeoki/wp/internal/app"

type Handlers struct {
	AddSource    *app.AddSourceHandler
	CreateSource *app.CreateSourceHandler
	DeleteSource *app.DeleteSourceHandler
	ListSources  *app.ListSourcesHandler
	RemoveSource *app.RemoveSourceHandler

	AddTag    *app.AddTagHandler
	CreateTag *app.CreateTagHandler
	DeleteTag *app.DeleteTagHandler
	ListTags  *app.ListTagsHandler
	RemoveTag *app.RemoveTagHandler
	RenameTag *app.RenameTagHandler

	CreateWallpaper *app.CreateWallpaperHandler
	DeleteWallpaper *app.DeleteWallpaperHandler
	FindWallpaper   *app.FindWallpaperHandler
	ShowWallpaper   *app.ShowWallpaperHandler

	AddAlias    *app.AddAliasHandler
	RemoveAlias *app.RemoveAliasHandler
	ListAliases *app.ListAliasesHandler
}

// NOTE: it looks very ugly, but i don't think it's possible to DRY.
// if you are going to make a generic wrapper function, then the compiler
// errors in the function body, because it doesn't know will the *Provider
// satisfy the specific provider or not

func NewHandlers(w app.Worker[*Provider]) *Handlers {
	return &Handlers{
		AddSource:    app.NewAddSourceHandler(addSource(w)),
		CreateSource: app.NewCreateSourceHandler(createSource(w)),
		DeleteSource: app.NewDeleteSourceHandler(deleteSource(w)),
		ListSources:  app.NewListSourcesHandler(listSources(w)),
		RemoveSource: app.NewRemoveSourceHandler(removeSource(w)),

		AddTag:    app.NewAddTagHandler(addTag(w)),
		CreateTag: app.NewCreateTagHandler(createTag(w)),
		DeleteTag: app.NewDeleteTagHandler(deleteTag(w)),
		ListTags:  app.NewListTagsHandler(listTags(w)),
		RemoveTag: app.NewRemoveTagHandler(removeTag(w)),
		RenameTag: app.NewRenameTagHandler(renameTag(w)),

		CreateWallpaper: app.NewCreateWallpaperHandler(createWallpaper(w)),
		DeleteWallpaper: app.NewDeleteWallpaperHandler(deleteWallpaper(w)),
		FindWallpaper:   app.NewFindWallpaperHandler(findWallpaper(w)),
		ShowWallpaper:   app.NewShowWallpaperHandler(showWallpaper(w)),

		AddAlias:    app.NewAddAliasHandler(addAlias(w)),
		RemoveAlias: app.NewRemoveAliasHandler(removeAlias(w)),
		ListAliases: app.NewListAliasesHandler(listAliases(w)),
	}
}

func addSource(w app.Worker[*Provider]) app.WorkerFunc[app.AddSourceProvider] {
	return func(ctx app.Ctx, j app.Job[app.AddSourceProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func createSource(
	w app.Worker[*Provider],
) app.WorkerFunc[app.CreateSourceProvider] {
	return func(ctx app.Ctx, j app.Job[app.CreateSourceProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func deleteSource(
	w app.Worker[*Provider],
) app.WorkerFunc[app.DeleteSourceProvider] {
	return func(ctx app.Ctx, j app.Job[app.DeleteSourceProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func listSources(
	w app.Worker[*Provider],
) app.WorkerFunc[app.ListSourcesProvider] {
	return func(ctx app.Ctx, j app.Job[app.ListSourcesProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func removeSource(
	w app.Worker[*Provider],
) app.WorkerFunc[app.RemoveSourceProvider] {
	return func(ctx app.Ctx, j app.Job[app.RemoveSourceProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}
func addTag(w app.Worker[*Provider]) app.WorkerFunc[app.AddTagProvider] {
	return func(ctx app.Ctx, j app.Job[app.AddTagProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func createTag(w app.Worker[*Provider]) app.WorkerFunc[app.CreateTagProvider] {
	return func(ctx app.Ctx, j app.Job[app.CreateTagProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func deleteTag(w app.Worker[*Provider]) app.WorkerFunc[app.DeleteTagProvider] {
	return func(ctx app.Ctx, j app.Job[app.DeleteTagProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func listTags(w app.Worker[*Provider]) app.WorkerFunc[app.ListTagsProvider] {
	return func(ctx app.Ctx, j app.Job[app.ListTagsProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func removeTag(w app.Worker[*Provider]) app.WorkerFunc[app.RemoveTagProvider] {
	return func(ctx app.Ctx, j app.Job[app.RemoveTagProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func renameTag(w app.Worker[*Provider]) app.WorkerFunc[app.RenameTagProvider] {
	return func(ctx app.Ctx, j app.Job[app.RenameTagProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func createWallpaper(
	w app.Worker[*Provider],
) app.WorkerFunc[app.CreateWallpaperProvider] {
	return func(ctx app.Ctx, j app.Job[app.CreateWallpaperProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func deleteWallpaper(
	w app.Worker[*Provider],
) app.WorkerFunc[app.DeleteWallpaperProvider] {
	return func(ctx app.Ctx, j app.Job[app.DeleteWallpaperProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func findWallpaper(
	w app.Worker[*Provider],
) app.WorkerFunc[app.FindWallpaperProvider] {
	return func(ctx app.Ctx, j app.Job[app.FindWallpaperProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func showWallpaper(
	w app.Worker[*Provider],
) app.WorkerFunc[app.ShowWallpaperProvider] {
	return func(ctx app.Ctx, j app.Job[app.ShowWallpaperProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func addAlias(
	w app.Worker[*Provider],
) app.WorkerFunc[app.AddAliasProvider] {
	return func(ctx app.Ctx, j app.Job[app.AddAliasProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func removeAlias(
	w app.Worker[*Provider],
) app.WorkerFunc[app.RemoveAliasProvider] {
	return func(ctx app.Ctx, j app.Job[app.RemoveAliasProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}

func listAliases(
	w app.Worker[*Provider],
) app.WorkerFunc[app.ListAliasesProvider] {
	return func(ctx app.Ctx, j app.Job[app.ListAliasesProvider]) error {
		return w.Work(ctx, func(p *Provider) error { return j(p) })
	}
}
