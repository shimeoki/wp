package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *TagRepo) Save(ctx domain.Ctx, t *domain.Tag) error {
	if tag, _ := r.FindByID(ctx, t.ID); tag == nil {
		return r.create(ctx, t)
	} else {
		return r.update(ctx, t)
	}
}

func (r *TagRepo) create(ctx domain.Ctx, t *domain.Tag) error {
	_, err := r.db.ExecContext(ctx, tagCreateQuery,
		t.ID.String(),
		t.Name.String(),
		t.CreatedAt,
		t.UpdatedAt,
	)

	return err
}

func (r *TagRepo) update(ctx domain.Ctx, t *domain.Tag) error {
	_, err := r.db.ExecContext(ctx, tagUpdateQuery,
		t.Name.String(),
		t.UpdatedAt,
		t.ID.String(),
	)

	return err
}
