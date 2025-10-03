package domain

import (
	"errors"
	"time"
)

type SourceRepo interface {
	Repo[*Source]
}

type Source struct {
	ID

	Name string
	Link *string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSource(name string, link *string) *Source {
	return &Source{
		ID: NewID(),

		Name: name,
		Link: link,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Source) Rename(name string) error {
	s.Name = name
	s.UpdatedAt = time.Now()
	return s.Validate()
}

func (s *Source) UpdateLink(link *string) error {
	s.Link = link
	s.UpdatedAt = time.Now()
	return s.Validate()
}

func (s *Source) Validate() error {
	if s.Name == "" {
		return errors.New("name is empty")
	}

	if s.CreatedAt.After(s.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
