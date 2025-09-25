package api

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/shimeoki/wp/internal/db"
)

type Exporter interface {
	Export(Ctx, io.Writer) error
}

type Importer interface {
	Import(Ctx, io.Reader) error
}

type Version int

var InvalidVersion = errors.New("unexpected version")

type Backup struct {
	Version    `json:"version"`
	Wallpapers `json:"wallpapers"`
}

type JSONer struct {
	repo db.Repo
	txer db.Txer

	tags map[Tag]db.ID
}

func NewJSONer(r db.RepoTxer) *JSONer {
	return &JSONer{repo: r, txer: r}
}

func (j *JSONer) Export(ctx Ctx, out io.Writer) error {
	ws, err := j.repo.Wallpapers().GetAll(ctx)
	if err != nil {
		return err
	}

	bak := &Backup{
		Version:    1,
		Wallpapers: make(map[Hash]*Wallpaper),
	}

	for _, w := range ws {
		hash := Hash(w.Hash)

		if bak.Wallpapers[hash] != nil {
			// FIXME: i really don't like this error
			return errors.New("duplicate hashes in database")
		}

		aliases := make([]Alias, len(w.Aliases))
		for i, a := range w.Aliases {
			aliases[i] = Alias(a.Name)
		}

		tags := make([]Tag, len(w.Tags))
		for i, t := range w.Tags {
			tags[i] = Tag(t.Name)
		}

		sources := make([]Source, len(w.Sources))
		for i, s := range w.Sources {
			sources[i] = Source{Name: Name(s.Name), Link: s.Link}
		}

		bak.Wallpapers[hash] = &Wallpaper{
			Format:  Format(w.Format),
			Aliases: aliases,
			Tags:    tags,
			Sources: sources,
		}
	}

	return json.NewEncoder(out).Encode(bak)
}

func (j *JSONer) importWallpaper(
	ctx Ctx,
	h Hash,
	f Format,
) (ID, error) {
	repo := j.repo.Wallpapers()
	hash := db.Hash(h)

	if w, _ := repo.GetByHash(ctx, hash); w != nil {
		return ID(w.ID), nil
	}

	id, err := repo.Create(
		ctx,
		&db.WallpaperCreate{Hash: hash, Format: string(f)},
	)

	return ID(id), err
}

func (j *JSONer) importAliases(ctx Ctx, aliases []Alias, wid ID) error {
	repo := j.repo.Aliases()

	for _, a := range aliases {
		_, err := repo.Create(
			ctx,
			&db.AliasCreate{WallpaperID: db.ID(wid), Name: db.Name(a)},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (j *JSONer) importSources(ctx Ctx, sources []Source, wid ID) error {
	sourceRepo := j.repo.Sources()
	wallRepo := j.repo.Wallpapers()

	for _, s := range sources {
		id, err := sourceRepo.Create(
			ctx,
			&db.SourceCreate{Name: db.Name(s.Name), Link: s.Link},
		)

		if err != nil {
			return err
		}

		err = wallRepo.AddSource(
			ctx,
			&db.WallpaperSource{WallpaperID: db.ID(wid), SourceID: id},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (j *JSONer) importTags(ctx Ctx, tags []Tag, wid ID) error {
	tagRepo := j.repo.Tags()
	wallRepo := j.repo.Wallpapers()

	for _, n := range tags {
		name := db.Name(n)

		if j.tags[n] == 0 {
			if tag, _ := tagRepo.GetByName(ctx, name); tag != nil {
				j.tags[n] = tag.ID
			}
		}

		if j.tags[n] == 0 {
			id, err := tagRepo.Create(ctx, &db.TagCreate{Name: name})
			if err != nil {
				return err
			}

			j.tags[n] = id
		}

		err := wallRepo.AddTag(
			ctx,
			&db.WallpaperTag{WallpaperID: db.ID(wid), TagID: j.tags[n]},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (j *JSONer) Import(ctx Ctx, in io.Reader) error {
	file := &Backup{}

	if err := json.NewDecoder(in).Decode(file); err != nil {
		return err
	}

	if file.Version != 1 {
		return InvalidVersion
	}

	tx, err := j.txer.WithTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	jsoner := &JSONer{repo: tx, tags: make(map[Tag]db.ID)}

	for hash, w := range file.Wallpapers {
		id, err := jsoner.importWallpaper(ctx, hash, w.Format)
		if err != nil {
			return err
		}

		if err := jsoner.importAliases(ctx, w.Aliases, id); err != nil {
			return err
		}

		if err := jsoner.importTags(ctx, w.Tags, id); err != nil {
			return err
		}

		if err := jsoner.importSources(ctx, w.Sources, id); err != nil {
			return err
		}
	}

	return tx.Commit()
}
