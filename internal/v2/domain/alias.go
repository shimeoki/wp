package domain

import (
	"errors"
	"time"
)

type AliasRepo interface {
	Repo[*Alias]
	FindByWallpaperID(Ctx, ID) ([]*Alias, error)
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
	if a.CreatedAt.After(a.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
