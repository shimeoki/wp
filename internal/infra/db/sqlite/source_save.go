package sqlite

import "github.com/shimeoki/wp/internal/domain"

func (r *SourceRepo) Save(ctx domain.Ctx, s *domain.Source) error {
	if source, _ := r.FindByID(ctx, s.ID); source == nil {
		return r.create(ctx, s)
	} else {
		return r.update(ctx, s)
	}
}

func (r *SourceRepo) create(ctx domain.Ctx, s *domain.Source) error {
	_, err := r.db.ExecContext(ctx, sourceCreateQuery,
		s.ID.String(),
		s.Name.String(),
		s.Link,
		s.CreatedAt,
		s.UpdatedAt,
	)

	return err
}

func (r *SourceRepo) update(ctx domain.Ctx, s *domain.Source) error {
	_, err := r.db.ExecContext(ctx, sourceUpdateQuery,
		s.Name.String(),
		s.Link,
		s.UpdatedAt,
		s.ID.String(),
	)

	return err
}
