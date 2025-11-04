package domain

import (
	"iter"
	"time"
)

type QueueRepo interface {
	Repo[*Queue]
	FindByWallpaperID(Ctx, ID) (iter.Seq[*Queue], error)
	FindByStatus(Ctx, Status) (iter.Seq[*Queue], error)
}

type Queue struct {
	ID
	Status
	WallpaperID ID
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewQueue(wid ID, s Status, priority int) (*Queue, error) {
	q := &Queue{
		ID:          NewID(),
		Status:      s,
		WallpaperID: wid,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := q.validate(); err != nil {
		return nil, err
	}

	return q, nil
}

func (q *Queue) UpdateStatus(s Status) error {
	q.Status = s
	q.UpdatedAt = time.Now()
	return q.validate()
}

func (q *Queue) UpdatePriority(priority int) error {
	q.Priority = priority
	q.UpdatedAt = time.Now()
	return q.validate()
}

func (q *Queue) validate() error {
	if err := ValidateTimestamps(q.CreatedAt, q.UpdatedAt); err != nil {
		return err
	}

	switch q.Status {
	case QUEUED, USED, SKIPPED:
		return nil
	}

	return NewInvalidStatusError(string(q.Status))
}
