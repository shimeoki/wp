package sqlite

import (
	_ "embed"

	"github.com/shimeoki/wp/internal/domain"
)

var (
	//go:embed sql/tag_count.sql
	tagCountQuery string

	//go:embed sql/tag_create.sql
	tagCreateQuery string

	//go:embed sql/tag_delete.sql
	tagDeleteQuery string

	//go:embed sql/tag_select.sql
	tagSelectQuery string

	//go:embed sql/tag_update.sql
	tagUpdateQuery string
)

type TagRepo struct {
	db DB
}

func NewTagRepo(db DB) *TagRepo {
	return &TagRepo{db: db}
}

type tagTable struct {
	ID        integer
	UUID      uid
	Name      text
	CreatedAt timestamp
	UpdatedAt timestamp
}

func (t *tagTable) toDomain() *domain.Tag {
	return &domain.Tag{
		ID:        domain.ID(t.UUID),
		Name:      domain.Name(t.Name),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
