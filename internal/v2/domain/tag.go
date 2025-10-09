package domain

import (
	"errors"
	"time"
)

type TagRepo interface {
	Repo[*Tag]
	FindByName(Ctx, string) (*Tag, error)
}

type Tag struct {
	ID

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTag(name string) *Tag {
	return &Tag{
		ID: NewID(),

		Name: name,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Tag) Rename(name string) error {
	t.Name = name
	t.UpdatedAt = time.Now()
	return t.Validate()
}

func (t *Tag) Validate() error {
	if t.Name == "" {
		return errors.New("name is empty")
	}

	if t.CreatedAt.After(t.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
