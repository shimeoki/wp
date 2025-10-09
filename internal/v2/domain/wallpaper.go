package domain

import (
	"errors"
	"iter"
	"time"
)

type WallpaperRepo interface {
	Repo[*Wallpaper]
	ByHash(Ctx, Hash) (*Wallpaper, error)
	ByTagID(Ctx, ID) (iter.Seq[*Wallpaper], error)
}

type Format string

const (
	JPEG Format = "jpg"
	PNG  Format = "png"
)

var InvalidFormat = errors.New("invalid format")

type Wallpaper struct {
	ID

	Format
	Hash

	Sources map[ID]*Source
	Tags    map[ID]*Tag

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWallpaper(f Format, h Hash) *Wallpaper {
	return &Wallpaper{
		ID: NewID(),

		Format: f,
		Hash:   h,

		Sources: make(map[ID]*Source),
		Tags:    make(map[ID]*Tag),

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (w *Wallpaper) UpdateHash(h Hash) error {
	w.Hash = h
	w.UpdatedAt = time.Now()
	return w.Validate()
}

func (w *Wallpaper) UpdateFormat(f Format) error {
	w.Format = f
	w.UpdatedAt = time.Now()
	return w.Validate()
}

func (w *Wallpaper) AddSource(s *Source) error {
	if err := s.Validate(); err != nil {
		return err
	}

	w.Sources[s.ID] = s
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) RemoveSource(id ID) error {
	if _, ok := w.Sources[id]; !ok {
		return errors.New("source not found")
	}

	delete(w.Sources, id)
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) AddTag(t *Tag) error {
	if err := t.Validate(); err != nil {
		return err
	}

	w.Tags[t.ID] = t
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) RemoveTag(id ID) error {
	if _, ok := w.Tags[id]; !ok {
		return errors.New("tag not found")
	}

	delete(w.Tags, id)
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) Validate() error {
	if w.Hash == "" {
		return errors.New("hash is empty")
	}

	if w.CreatedAt.After(w.UpdatedAt) {
		return errors.New("created is after updated")
	}

	switch w.Format {
	case JPEG:
	case PNG:
		return nil
	}

	return InvalidFormat
}
