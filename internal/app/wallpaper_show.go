package app

import (
	"io"

	"github.com/shimeoki/wp/internal/domain"
)

type ShowWallpaperProvider interface {
	StoreProvider
	WallpaperProvider
}

type ShowWallpaperWorker Worker[ShowWallpaperProvider]

type ShowWallpaperHandler struct {
	worker ShowWallpaperWorker
	logger Logger
}

func NewShowWallpaperHandler(
	w ShowWallpaperWorker,
	l Logger,
) *ShowWallpaperHandler {
	return &ShowWallpaperHandler{worker: w, logger: l}
}

type ShowWallpaperQuery struct {
	Hash string
}

type ShowWallpaperResult struct {
	Image  io.ReadCloser
	Format string
}

func (h *ShowWallpaperHandler) Handle(
	ctx Ctx,
	qry *ShowWallpaperQuery,
) (*ShowWallpaperResult, error) {
	var r ShowWallpaperResult

	Act(h.logger, ctx, "showing wallpaper", qry)
	if err := h.worker.Work(ctx, func(p ShowWallpaperProvider) error {
		hash, err := domain.ParseHash(qry.Hash)
		if err != nil {
			return err
		}

		wall, _ := p.WallpaperRepo().FindByHash(ctx, hash)
		if wall == nil {
			return domain.NewNotFoundError("wallpaper", "hash", hash.String())
		}

		img, err := p.Store().Get(ctx, hash)
		if err != nil {
			return err
		}

		r.Image = img
		r.Format = wall.Format.String()

		return nil
	}); err != nil {
		Fail(h.logger, ctx, "failed to show wallpaper", err)
		return nil, err
	}

	return &r, nil
}
