package sqlite

import (
	_ "embed"

	"github.com/shimeoki/wp/internal/domain"
)

var (
	//go:embed sql/alias_select.sql
	aliasSelectQuery string

	//go:embed sql/alias_count.sql
	aliasCountQuery string
)

type AliasRepo struct {
	db DB
}

func NewAliasRepo(db DB) *AliasRepo {
	return &AliasRepo{db: db}
}

type aliasTable struct {
	ID            integer
	UUID          uid
	WallpaperID   integer
	WallpaperUUID uid
	Name          text
	CreatedAt     timestamp
	UpdatedAt     timestamp
}

func (t *aliasTable) toDomain() *domain.Alias {
	return &domain.Alias{
		ID:          domain.ID(t.UUID),
		Name:        domain.Name(t.Name),
		WallpaperID: domain.ID(t.WallpaperUUID),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
