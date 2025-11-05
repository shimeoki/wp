package app

import "github.com/shimeoki/wp/internal/domain"

type FindWallpaperProvider interface {
	WallpaperProvider
}

type FindWallpaperWorker Worker[FindWallpaperProvider]

type FindWallpaperHandler struct {
	worker FindWallpaperWorker
	logger Logger
}

func NewFindWallpaperHandler(
	w FindWallpaperWorker,
	l Logger,
) *FindWallpaperHandler {
	return &FindWallpaperHandler{worker: w, logger: l}
}

type FindWallpaperQuery struct {
	Hash string
}

type FindWallpaperResult struct {
	Format string
}

func (h *FindWallpaperHandler) Handle(
	ctx Ctx,
	qry *FindWallpaperQuery,
) (*FindWallpaperResult, error) {
	var r FindWallpaperResult

	Act(h.logger, ctx, "finding wallpaper", qry)
	if err := h.worker.Work(ctx, func(p FindWallpaperProvider) error {
		hash, err := domain.ParseHash(qry.Hash)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo().FindByHash(ctx, hash)
		if wall == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		r.Format = wall.Format.String()

		return nil
	}); err != nil {
		Fail(h.logger, ctx, "failed to find wallpaper", err)
		return nil, err
	}

	return &r, nil
}
