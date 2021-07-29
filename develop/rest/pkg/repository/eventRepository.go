package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"simplerest/pkg/model"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e *EventRepository) GetEventsForDay(today, tomorrow string) ([]model.Event, error) {
	stmt, err := e.db.Prepare("" +
		"SELECT * " +
		"FROM events " +
		"WHERE date BETWEEN $1 AND $2")
	if err != nil {
		errorText := "Error Prepare in GetEventsForDay"
		fmt.Println(errorText)
		return nil, errors.New(errorText)
	}

	rows, err := stmt.Query(today, tomorrow)
	if err != nil {
		errorText := "Error Query in GetEventsForDay"
		fmt.Println(errorText)
		return nil, errors.New(errorText)
	}

	var events []model.Event

	for rows.Next() {
		var event model.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Date)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		events = append(events, event)
	}
	defer rows.Close()

	return events, nil
}
