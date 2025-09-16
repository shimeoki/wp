package db

import (
	"context"
	"encoding/json"
	"errors"
	"io"
)

type jsonWallpaperSource struct {
	Name string  `json:"name"`
	Link *string `json:"link,omitempty"`
}

type jsonWallpaper struct {
	Extension string                `json:"extension"`
	Aliases   []string              `json:"aliases"`
	Tags      []string              `json:"tags"`
	Sources   []jsonWallpaperSource `json:"sources"`
}

type jsonExport struct {
	Version    int                       `json:"version"`
	Wallpapers map[string]*jsonWallpaper `json:"wallpapers"`
}

type JSONExporter struct {
	repo Repo
}

func NewJSONExporter(r Repo) *JSONExporter {
	return &JSONExporter{repo: r}
}

func (e *JSONExporter) Export(ctx context.Context, out io.Writer) error {
	ws, err := e.repo.Wallpapers().GetAll(ctx)
	if err != nil {
		return err
	}

	export := &jsonExport{
		Version:    1,
		Wallpapers: make(map[string]*jsonWallpaper),
	}

	for _, w := range ws {
		if export.Wallpapers[w.Hash] != nil {
			return errors.New("duplicate hashes in database")
		}

		aliases := make([]string, len(w.Aliases))
		for i, a := range w.Aliases {
			aliases[i] = a.Name
		}

		tags := make([]string, len(w.Tags))
		for i, t := range w.Tags {
			tags[i] = t.Name
		}

		sources := make([]jsonWallpaperSource, len(w.Sources))
		for i, s := range w.Sources {
			sources[i] = jsonWallpaperSource{Name: s.Name, Link: s.Link}
		}

		jw := &jsonWallpaper{
			Extension: w.Extension,
			Aliases:   aliases,
			Tags:      tags,
			Sources:   sources,
		}

		export.Wallpapers[w.Hash] = jw
	}

	return json.NewEncoder(out).Encode(export)
}
