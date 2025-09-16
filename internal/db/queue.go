package db

import (
	"context"
	"database/sql"
	"time"
)

type Queue struct {
	ID          int64
	WallpaperID int64
	Status      *Status
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueueCreate struct {
	WallpaperID int64
	StatusID    int64
	Priority    int
}

type QueueUpdate struct {
	ID       int64
	StatusID int64
	Priority int
}

type QueueRepo interface {
	GetAll(ctx context.Context) ([]*Queue, error)
	GetByID(ctx context.Context, id int64) (*Queue, error)

	Create(ctx context.Context, q *QueueCreate) (int64, error)
	Update(ctx context.Context, q *QueueUpdate) error
	Delete(ctx context.Context, id int64) error
}

type sqliteQueueRepo struct {
	db *sql.DB
}

func (r *sqliteQueueRepo) GetAll(ctx context.Context) ([]*Queue, error) {
	sql := `
		select
			q.id
			, q.wallpaper_id
			, q.priority
			, q.created_at
			, q.updated_at
			, s.id
			, s.name
		from queue as q
		left join status as s on q.status_id = s.id`

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var qs []*Queue
	statuses := make(map[int64]*Status)

	for rows.Next() {
		var s Status
		q := Queue{Status: &s}

		err := rows.Scan(
			&q.ID,
			&q.WallpaperID,
			&q.Priority,
			&q.CreatedAt,
			&q.UpdatedAt,
			&s.ID,
			&s.Name,
		)

		if err != nil {
			return nil, err
		}

		if statuses[q.Status.ID] == nil {
			statuses[q.Status.ID] = q.Status
		}

		q.Status = statuses[q.Status.ID] // reuse same structs
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
			q.id
			, q.wallpaper_id
			, q.priority
			, q.created_at
			, q.updated_at
			, s.id
			, s.name
		from queue as q
		left join status as s on q.status_id = s.id
		where q.id = ?`

	row := r.db.QueryRowContext(ctx, sql, id)
	var s Status
	q := Queue{Status: &s}

	err := row.Scan(
		&q.ID,
		&q.WallpaperID,
		&q.Priority,
		&q.CreatedAt,
		&q.UpdatedAt,
		&s.ID,
		&s.Name,
	)

	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (r *sqliteQueueRepo) Create(
	ctx context.Context,
	q *QueueCreate,
) (int64, error) {
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
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *sqliteQueueRepo) Update(ctx context.Context, q *QueueUpdate) error {
	sql := `update queue set status_id = ?, priority = ? where id = ?`

	_, err := r.db.ExecContext(
		ctx,
		sql,
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
