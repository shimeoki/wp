package domain

import (
	"errors"
	"time"
)

type AliasRepo interface {
	Repo[*Alias]
	ByWallpaperID(Ctx, ID) ([]*Alias, error)
}

type Alias struct {
	ID
	WallpaperID ID

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAlias(name string, wid ID) *Alias {
	return &Alias{
		ID:          NewID(),
		WallpaperID: wid,

		Name: name,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Alias) Rename(name string) error {
	a.Name = name
	a.UpdatedAt = time.Now()
	return a.Validate()
}

func (a *Alias) Validate() error {
	if a.Name == "" {
		return errors.New("name is empty")
	}

	if a.CreatedAt.After(a.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
