package domain

import (
	"iter"
	"time"
)

type WallpaperRepo interface {
	Repo[*Wallpaper]
	FindByHash(Ctx, Hash) (*Wallpaper, error)
	FindByTagID(Ctx, ID) (iter.Seq[*Wallpaper], error)
}

type Wallpaper struct {
	ID
	Format
	Hash
	Sources   map[ID]*Source
	Tags      map[ID]*Tag
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewWallpaper(f Format, h Hash) (*Wallpaper, error) {
	w := &Wallpaper{
		ID:        NewID(),
		Format:    f,
		Hash:      h,
		Sources:   make(map[ID]*Source),
		Tags:      make(map[ID]*Tag),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := w.validate(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Wallpaper) UpdateHash(h Hash) error {
	w.Hash = h
	w.UpdatedAt = time.Now()
	return w.validate()
}

func (w *Wallpaper) UpdateFormat(f Format) error {
	w.Format = f
	w.UpdatedAt = time.Now()
	return w.validate()
}

func (w *Wallpaper) AddSource(s *Source) error {
	w.Sources[s.ID] = s
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Wallpaper) RemoveSource(id ID) error {
	if _, ok := w.Sources[id]; !ok {
		return NewInvalidRelationError("source", "id", id.String(), "not found")
	}

	delete(w.Sources, id)
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) AddTag(t *Tag) error {
	w.Tags[t.ID] = t
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Wallpaper) RemoveTag(id ID) error {
	if _, ok := w.Tags[id]; !ok {
		return NewInvalidRelationError("tag", "id", id.String(), "not found")
	}

	delete(w.Tags, id)
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallpaper) validate() error {
	if err := ValidateTimestamps(w.CreatedAt, w.UpdatedAt); err != nil {
		return err
	}

	switch w.Format {
	case JPEG, PNG:
		return nil
	}

	return ErrInvalidFormat
}
