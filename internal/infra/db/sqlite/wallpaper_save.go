package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *WallpaperRepo) Save(ctx domain.Ctx, w *domain.Wallpaper) error {
	if wallpaper, _ := r.FindByID(ctx, w.ID); wallpaper == nil {
		return r.create(ctx, w)
	} else {
		return r.update(ctx, w)
	}
}

func (r *WallpaperRepo) create(ctx domain.Ctx, w *domain.Wallpaper) error {
	if _, err := r.db.ExecContext(ctx, wallpaperCreateQuery,
		w.ID.String(),
		w.Format.String(),
		w.Hash.String(),
		w.CreatedAt,
		w.UpdatedAt,
	); err != nil {
		return err
	}

	return r.updateJoins(ctx, w)
}

func (r *WallpaperRepo) update(ctx domain.Ctx, w *domain.Wallpaper) error {
	_, err := r.db.ExecContext(ctx, wallpaperUpdateQuery,
		w.Format.String(),
		w.Hash.String(),
		w.UpdatedAt.String(),
		w.ID.String(),
	)

	return err
}

func (r *WallpaperRepo) updateJoins(
	ctx domain.Ctx,
	w *domain.Wallpaper,
) error {
	now, err := r.FindByID(ctx, w.ID)
	if err != nil {
		return err
	}

	wid := w.ID.String()

	if err := r.updateTags(ctx, wid, now.Tags, w.Tags); err != nil {
		return err
	}

	if err := r.updateSources(ctx, wid, now.Sources, w.Sources); err != nil {
		return err
	}

	return nil
}

func (r *WallpaperRepo) updateTags(
	ctx domain.Ctx,
	wid string,
	current, target map[domain.ID]*domain.Tag,
) error {
	var newTags []domain.ID

	for id := range target {
		if _, ok := current[id]; ok {
			delete(current, id)
		} else {
			newTags = append(newTags, id)
		}
	}

	for _, id := range newTags {
		_, err := r.db.ExecContext(ctx, wallpaperTagCreateQuery,
			wid, id.String())

		if err != nil {
			return err
		}
	}

	for id := range current {
		_, err := r.db.ExecContext(ctx, wallpaperTagDeleteQuery,
			wid, id.String())

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *WallpaperRepo) updateSources(
	ctx domain.Ctx,
	wid string,
	current, target map[domain.ID]*domain.Source,
) error {
	var newSources []domain.ID

	for id := range target {
		if _, ok := current[id]; ok {
			delete(current, id)
		} else {
			newSources = append(newSources, id)
		}
	}

	for _, id := range newSources {
		_, err := r.db.ExecContext(ctx, wallpaperSourceCreateQuery,
			wid, id.String())

		if err != nil {
			return err
		}
	}

	for id := range current {
		_, err := r.db.ExecContext(ctx, wallpaperSourceDeleteQuery,
			wid, id.String())

		if err != nil {
			return err
		}
	}

	return nil
}
