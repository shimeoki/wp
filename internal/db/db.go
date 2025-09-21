package db

import (
	"context"
	"database/sql"
	_ "embed"
	"io"
	"time"

	"github.com/shimeoki/wp/internal/config"
	_ "modernc.org/sqlite"
)

type ID int64
type Hash string

type Exporter interface {
	Export(ctx context.Context, out io.Writer) error
}

type Importer interface {
	Import(ctx context.Context, in io.Reader) error
}

type Repo interface {
	Wallpapers() WallpaperRepo
	Tags() TagRepo
	Sources() SourceRepo
	Aliases() AliasRepo
	Statuses() StatusRepo
	Queues() QueueRepo
}

//go:embed sqlite.sql
var sqliteScheme string

type sqliteRepo struct {
	wallpapers WallpaperRepo
	tags       TagRepo
	sources    SourceRepo
	aliases    AliasRepo
	statuses   StatusRepo
	queues     QueueRepo
}

func NewSQLiteRepo(config *config.DB) (Repo, error) {
	db, err := sql.Open("sqlite", config.DataSourceName)
	if err != nil {
		return nil, err
	}

	err = ping(db)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(sqliteScheme); err != nil {
		return nil, err
	}

	r := &sqliteRepo{
		wallpapers: &sqliteWallpaperRepo{db: db},
		tags:       &sqliteTagRepo{db: db},
		sources:    &sqliteSourceRepo{db: db},
		aliases:    &sqliteAliasRepo{db: db},
		statuses:   &sqliteStatusRepo{db: db},
		queues:     &sqliteQueueRepo{db: db},
	}

	return r, nil
}

func ping(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}

func (r *sqliteRepo) Wallpapers() WallpaperRepo {
	return r.wallpapers
}

func (r *sqliteRepo) Tags() TagRepo {
	return r.tags
}

func (r *sqliteRepo) Sources() SourceRepo {
	return r.sources
}

func (r *sqliteRepo) Aliases() AliasRepo {
	return r.aliases
}

func (r *sqliteRepo) Statuses() StatusRepo {
	return r.statuses
}

func (r *sqliteRepo) Queues() QueueRepo {
	return r.queues
}
