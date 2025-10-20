package domain

import (
	"errors"
	"time"
)

type TagRepo interface {
	Repo[*Tag]
	FindByName(Ctx, Name) (*Tag, error)
}

type Tag struct {
	ID
	Name
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTag(n Name) (*Tag, error) {
	t := &Tag{
		ID:        NewID(),
		Name:      n,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Tag) Rename(n Name) error {
	t.Name = n
	t.UpdatedAt = time.Now()
	return t.validate()
}

func (t *Tag) validate() error {
	if t.CreatedAt.After(t.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
