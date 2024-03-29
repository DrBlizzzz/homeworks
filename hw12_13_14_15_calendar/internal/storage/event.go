package storage

import (
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/net/context"
)

type Event struct {
	ID               uuid.UUID     `db:"id"`
	Title            string        `db:"title"`
	TimeStart        time.Time     `db:"start_date"`
	Duration         time.Duration `db:"duration"`
	Description      string        `db:"description"`
	UserID           uuid.UUID     `db:"user_id"`
	NotifyBeforeDays int           `db:"notify_before"`
}

type Storage interface {
	CreateEvent(ctx context.Context, event *Event) error
	GetEventID(ctx context.Context, event *Event) (uuid.UUID, error)
	UpdateEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEventsPerDay(ctx context.Context, day time.Time) ([]Event, error)
	GetEventsPerWeek(ctx context.Context, beginDate time.Time) ([]Event, error)
	GetEventsPerMonth(ctx context.Context, beginDate time.Time) ([]Event, error)
	Close(ctx context.Context) error
	ListForScheduler(ctx context.Context, remindFor time.Duration, period time.Duration) ([]Notification, error)
	ClearEvents(ctx context.Context) error
}

type Notification struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	TimeStart time.Time `db:"start_date" json:"time_start"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}
