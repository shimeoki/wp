package sqlite

import "github.com/shimeoki/wp/internal/domain"

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
