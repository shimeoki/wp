package domain

import (
	"iter"
	"time"
)

type AliasRepo interface {
	Repo[*Alias]
	FindByWallpaperID(Ctx, ID) (iter.Seq[*Alias], error)
}

type Alias struct {
	ID
	Name
	WallpaperID ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAlias(n Name, wid ID) (*Alias, error) {
	alias := &Alias{
		ID:          NewID(),
		Name:        n,
		WallpaperID: wid,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := alias.validate(); err != nil {
		return nil, err
	}

	return alias, nil
}

func (a *Alias) Rename(n Name) error {
	a.Name = n
	a.UpdatedAt = time.Now()
	return a.validate()
}

func (a *Alias) validate() error {
	return ValidateTimestamps(a.CreatedAt, a.UpdatedAt)
}
