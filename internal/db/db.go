package db

import "database/sql"

// todo: import driver

type Repo interface {
	Wallpapers() WallpaperRepo
	Tags() TagRepo
	Sources() SourceRepo
	Aliases() AliasRepo
	Statuses() StatusRepo
	Queues() QueueRepo
}

type sqliteRepo struct {
	wallpapers WallpaperRepo
	tags       TagRepo
	sources    SourceRepo
	aliases    AliasRepo
	statuses   StatusRepo
	queues     QueueRepo
}

func NewSQLiteRepo(dsn string) (Repo, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
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
