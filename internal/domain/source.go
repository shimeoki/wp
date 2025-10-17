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
	Name
	Link      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSource(n Name, link *string) (*Source, error) {
	s := &Source{
		ID:        NewID(),
		Name:      n,
		Link:      link,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.validate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Source) Rename(n Name) error {
	s.Name = n
	s.UpdatedAt = time.Now()
	return s.validate()
}

func (s *Source) UpdateLink(link *string) error {
	s.Link = link
	s.UpdatedAt = time.Now()
	return s.validate()
}

func (s *Source) validate() error {
	if s.CreatedAt.After(s.UpdatedAt) {
		return errors.New("created is after updated")
	}

	return nil
}
