package repository

import (
	"database/sql"
	"simplerest/pkg/model"
)

type Event interface {
	GetEventsForDay(today, tomorrow string) ([]model.Event, error)
}

type Repository struct {
	Event
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Event: NewEventRepository(db),
	}
}
