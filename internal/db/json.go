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

type jsonFile struct {
	Version    int                       `json:"version"`
	Wallpapers map[string]*jsonWallpaper `json:"wallpapers"`
}

type JSONer struct {
	repo Repo
}

func NewJSONer(r Repo) *JSONer {
	return &JSONer{repo: r}
}

func (j *JSONer) Export(ctx context.Context, out io.Writer) error {
	ws, err := j.repo.Wallpapers().GetAll(ctx)
	if err != nil {
		return err
	}

	file := &jsonFile{
		Version:    1,
		Wallpapers: make(map[string]*jsonWallpaper),
	}

	for _, w := range ws {
		if file.Wallpapers[w.Hash] != nil {
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

		file.Wallpapers[w.Hash] = jw
	}

	return json.NewEncoder(out).Encode(file)
}

func (j *JSONer) createWallpaper(
	ctx context.Context,
	hash, extension string,
) (int64, error) {
	ws := j.repo.Wallpapers()

	w, _ := ws.GetByHash(ctx, hash)
	if w != nil {
		return w.ID, nil
	}

	return ws.Create(
		ctx,
		&WallpaperCreate{Hash: hash, Extension: extension},
	)
}

func (j *JSONer) createAliases(
	ctx context.Context,
	aliases []string,
	wid int64,
) error {
	as := j.repo.Aliases()

	for _, alias := range aliases {
		_, err := as.Create(ctx, &AliasCreate{WallpaperID: wid, Name: alias})
		if err != nil {
			return err
		}
	}

	return nil
}

func (j *JSONer) createSources(
	ctx context.Context,
	sources []jsonWallpaperSource,
	wid int64,
) error {
	ss := j.repo.Sources()
	ws := j.repo.Wallpapers()

	for _, source := range sources {
		id, err := ss.Create(
			ctx,
			&SourceCreate{Name: source.Name, Link: source.Link},
		)

		if err != nil {
			return err
		}

		err = ws.AddSource(
			ctx,
			&WallpaperSource{WallpaperID: wid, SourceID: id},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

type jsonTagger struct {
	jsoner *JSONer
	tags   map[string]int64
}

func (t *jsonTagger) createTags(
	ctx context.Context,
	tags []string,
	wid int64,
) error {
	ts := t.jsoner.repo.Tags()
	ws := t.jsoner.repo.Wallpapers()

	for _, name := range tags {
		if t.tags[name] == 0 {
			tag, err := ts.GetByName(ctx, name)
			if err != nil {
				return err
			}

			t.tags[name] = tag.ID
		}

		if t.tags[name] == 0 {
			id, err := ts.Create(ctx, &TagCreate{Name: name})
			if err != nil {
				return err
			}

			t.tags[name] = id
		}

		err := ws.AddTag(
			ctx,
			&WallpaperTag{WallpaperID: wid, TagID: t.tags[name]},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// todo: use a single transaction
func (j *JSONer) Import(ctx context.Context, in io.Reader) error {
	file := &jsonFile{}

	if err := json.NewDecoder(in).Decode(file); err != nil {
		return err
	}

	if file.Version != 1 {
		return errors.New("unexpected version")
	}

	tagger := &jsonTagger{jsoner: j, tags: make(map[string]int64)}

	for hash, jw := range file.Wallpapers {
		id, err := j.createWallpaper(ctx, hash, jw.Extension)
		if err != nil {
			return err
		}

		if err := j.createAliases(ctx, jw.Aliases, id); err != nil {
			return err
		}

		if err := tagger.createTags(ctx, jw.Tags, id); err != nil {
			return err
		}

		if err := j.createSources(ctx, jw.Sources, id); err != nil {
			return err
		}
	}

	return nil
}
