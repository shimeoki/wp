package db

import (
	"database/sql"
	"time"
)

type Queue struct {
	ID
	WallpaperID ID
	Status      *Status
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type QueueCreate struct {
	WallpaperID ID
	StatusID    ID
	Priority    int
}

type QueueUpdate struct {
	ID
	StatusID ID
	Priority int
}

type QueueRepo interface {
	GetAll(Ctx) ([]*Queue, error)
	GetByID(Ctx, ID) (*Queue, error)

	Create(Ctx, *QueueCreate) (ID, error)
	Update(Ctx, *QueueUpdate) error
	Delete(Ctx, ID) error
}

type sqliteQueueRepo struct {
	db *sql.DB
}

func (r *sqliteQueueRepo) GetAll(ctx Ctx) ([]*Queue, error) {
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
	statuses := make(map[ID]*Status)

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

func (r *sqliteQueueRepo) GetByID(ctx Ctx, id ID) (*Queue, error) {
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

func (r *sqliteQueueRepo) Create(ctx Ctx, q *QueueCreate) (ID, error) {
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

	return ID(id), nil
}

func (r *sqliteQueueRepo) Update(ctx Ctx, q *QueueUpdate) error {
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

func (r *sqliteQueueRepo) Delete(ctx Ctx, id ID) error {
	sql := "delete from queue where id = ?"

	_, err := r.db.ExecContext(ctx, sql, id)

	return err
}
