package sqlite

import (
	_ "embed"

	"github.com/shimeoki/wp/internal/domain"
)

var (
	//go:embed sql/source_select.sql
	sourceSelectQuery string

	//go:embed sql/source_count.sql
	sourceCountQuery string

	//go:embed sql/source_delete.sql
	sourceDeleteQuery string

	//go:embed sql/source_create.sql
	sourceCreateQuery string

	//go:embed sql/source_update.sql
	sourceUpdateQuery string
)

type SourceRepo struct {
	db DB
}

func NewSourceRepo(db DB) *SourceRepo {
	return &SourceRepo{db: db}
}

type sourceTable struct {
	ID        integer
	UUID      uid
	Name      text
	Link      *text
	CreatedAt timestamp
	UpdatedAt timestamp
}

func (t *sourceTable) toDomain() *domain.Source {
	return &domain.Source{
		ID:        domain.ID(t.UUID),
		Name:      domain.Name(t.Name),
		Link:      t.Link,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
