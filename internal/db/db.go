package db

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"time"

	"github.com/shimeoki/wp/internal/config"
	_ "modernc.org/sqlite"
)

type Ctx = context.Context

type ID int64
type Hash string
type Name string

type DB interface {
	ExecContext(ctx Ctx, query string, args ...any) (sql.Result, error)
	QueryContext(ctx Ctx, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx Ctx, query string, args ...any) *sql.Row
}

type Repo interface {
	Wallpapers() WallpaperRepo
	Tags() TagRepo
	Sources() SourceRepo
	Aliases() AliasRepo
	Statuses() StatusRepo
	Queues() QueueRepo
}

var InvalidTx = errors.New("invalid transaction")

type TxOptions = sql.TxOptions

type Tx interface {
	Commit() error
	Rollback() error
}

type Txer interface {
	WithTx(Ctx, *TxOptions) (RepoTx, error)
}

type RepoTx interface {
	Repo
	Tx
}

type RepoTxer interface {
	Repo
	Txer
}

//go:embed sqlite.sql
var sqliteScheme string

type SQLiteRepo struct {
	db *sql.DB
	tx *sql.Tx

	wallpapers WallpaperRepo
	tags       TagRepo
	sources    SourceRepo
	aliases    AliasRepo
	statuses   StatusRepo
	queues     QueueRepo
}

func NewSQLiteRepo(config *config.DB) (*SQLiteRepo, error) {
	db, err := sql.Open("sqlite", config.DataSourceName)
	if err != nil {
		return nil, err
	}

	if err := ping(db); err != nil {
		return nil, err
	}

	if _, err := db.Exec(sqliteScheme); err != nil {
		return nil, err
	}

	r := &SQLiteRepo{db: db}
	r.swap(db)

	return r, nil
}

func ping(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}

func (r *SQLiteRepo) WithTx(ctx Ctx, opts *TxOptions) (RepoTx, error) {
	tx, err := r.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	repo := &SQLiteRepo{db: r.db, tx: tx}
	repo.swap(tx)

	return repo, nil
}

func (r *SQLiteRepo) swap(db DB) {
	r.wallpapers = &sqliteWallpaperRepo{db: db}
	r.tags = &sqliteTagRepo{db: db}
	r.sources = &sqliteSourceRepo{db: db}
	r.aliases = &sqliteAliasRepo{db: db}
	r.statuses = &sqliteStatusRepo{db: db}
	r.queues = &sqliteQueueRepo{db: db}
}

func (r *SQLiteRepo) Commit() error {
	if r.tx == nil {
		return InvalidTx
	}

	return r.tx.Commit()
}

func (r *SQLiteRepo) Rollback() error {
	if r.tx == nil {
		return InvalidTx
	}

	return r.tx.Rollback()
}

func (r *SQLiteRepo) Wallpapers() WallpaperRepo {
	return r.wallpapers
}

func (r *SQLiteRepo) Tags() TagRepo {
	return r.tags
}

func (r *SQLiteRepo) Sources() SourceRepo {
	return r.sources
}

func (r *SQLiteRepo) Aliases() AliasRepo {
	return r.aliases
}

func (r *SQLiteRepo) Statuses() StatusRepo {
	return r.statuses
}

func (r *SQLiteRepo) Queues() QueueRepo {
	return r.queues
}
