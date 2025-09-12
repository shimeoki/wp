package db

import (
	"context"
	"database/sql"
	"time"
)

type Queue struct {
	ID          int64 // primary key
	WallpaperID int64 // foreign key
	StatusID    int64 // foreign key
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueueRepo interface {
	GetAll(ctx context.Context) ([]*Queue, error)
	GetByID(ctx context.Context, id int64) (*Queue, error)

	Create(ctx context.Context, q *Queue) error
	Update(ctx context.Context, q *Queue) error
	Delete(ctx context.Context, id int64) error
}

type sqliteQueueRepo struct {
	db *sql.DB
}

func (r *sqliteQueueRepo) GetAll(ctx context.Context) ([]*Queue, error) {
	sql := `
		select
			id,
			wallpaper_id,
			status_id,
			priority,
			created_at,
			updated_at
		from queue`

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var qs []*Queue

	for rows.Next() {
		var q Queue

		err := rows.Scan(
			&q.ID,
			&q.WallpaperID,
			&q.StatusID,
			&q.Priority,
			&q.CreatedAt,
			&q.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		qs = append(qs, &q)
	}

	return qs, rows.Err()
}

func (r *sqliteQueueRepo) GetByID(
	ctx context.Context,
	id int64,
) (*Queue, error) {
	sql := `
		select
			id,
			wallpaper_id,
			status_id,
			priority,
			created_at,
			updated_at
		from queue where id = ?`

	var q Queue
	err := r.db.QueryRowContext(ctx, sql, id).
		Scan(
			&q.ID,
			&q.WallpaperID,
			&q.StatusID,
			&q.Priority,
			&q.CreatedAt,
			&q.UpdatedAt,
		)

	return &q, err
}

func (r *sqliteQueueRepo) Create(ctx context.Context, q *Queue) error {
	sql := `
		insert into queue(wallpaper_id, status_id, priority)
		values(?, ?, ?)`

	result, err := r.db.ExecContext(
		ctx,
		sql,
		q.WallpaperID,
		q.StatusID,
		q.Priority,
	)

	if err != nil {
		return err
	}

	q.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteQueueRepo) Update(ctx context.Context, q *Queue) error {
	sql := `
		update queue set
			wallpaper_id = ?,
			status_id = ?,
			priority = ?
		where id = ?`

	_, err := r.db.ExecContext(
		ctx,
		sql,
		q.WallpaperID,
		q.StatusID,
		q.Priority,
		q.ID,
	)

	return err
}

func (r *sqliteQueueRepo) Delete(ctx context.Context, id int64) error {
	sql := "delete from queue where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
