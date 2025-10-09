package domain

import (
	"errors"
	"time"
)

type QueueRepo interface {
	Repo[*Queue]
	FindByWallpaperID(Ctx, ID) ([]*Queue, error)
	FindByStatus(Ctx, Status) ([]*Queue, error)
}

type Queue struct {
	ID
	WallpaperID ID

	Status
	Priority int

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewQueue(wid ID, s Status, priority int) *Queue {
	return &Queue{
		ID:          NewID(),
		WallpaperID: wid,

		Status:   s,
		Priority: priority,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (q *Queue) UpdateStatus(s Status) error {
	q.Status = s
	q.UpdatedAt = time.Now()
	return q.Validate()
}

func (q *Queue) UpdatePriority(priority int) error {
	q.Priority = priority
	q.UpdatedAt = time.Now()
	return q.Validate()
}

func (q *Queue) Validate() error {
	if q.CreatedAt.After(q.UpdatedAt) {
		return errors.New("created is after updated")
	}

	switch q.Status {
	case QUEUED:
	case USED:
	case SKIPPED:
		return nil
	}

	return InvalidStatus
}
