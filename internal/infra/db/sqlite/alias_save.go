package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *AliasRepo) Save(ctx domain.Ctx, a *domain.Alias) error {
	if alias, _ := r.FindByID(ctx, a.ID); alias == nil {
		return r.create(ctx, a)
	} else {
		return r.update(ctx, a)
	}
}

func (r *AliasRepo) update(ctx domain.Ctx, a *domain.Alias) error {
	_, err := r.db.ExecContext(ctx, aliasUpdateQuery,
		a.Name.String(),
		a.UpdatedAt,
	)

	return err
}

func (r *AliasRepo) create(ctx domain.Ctx, a *domain.Alias) error {
	_, err := r.db.ExecContext(ctx, aliasCreateQuery,
		a.ID.String(),
		a.WallpaperID.String(),
		a.Name.String(),
		a.CreatedAt,
		a.UpdatedAt,
	)

	return err
}
