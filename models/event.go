package models

import (
	"errors"
	"go/by/example/restful/api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
    INSERT INTO events(name, description, location, dateTime, user_id)
    VALUES (?, ?, ?, ?, ?)
  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("when setup qerry")
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.UserID,
	)
	if err != nil {
		return errors.New("when run qerry")
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	eventsRow, err := db.DB.Query(query)
	defer eventsRow.Close()
	if err != nil {
		return nil, err
	}

	var eventList []Event

	for eventsRow.Next() {
		var event Event
		err := eventsRow.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)

		if err != nil {
			return nil, err
		}

		eventList = append(eventList, event)
	}
	return eventList, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	getQuery := `
  UPDATE events
  SET name = ?, description = ?, location = ?, dateTime = ?
  WHERE id = ?
  `
	stmt, err := db.DB.Prepare(getQuery)
	if err != nil {
		return errors.New("uncorrect query")
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
func (event Event) Delete() error {
	delQuery := `DELETE FROM events WHERE id = ? `
	stmt, err := db.DB.Prepare(delQuery)
	if err != nil {
		return errors.New("uncorrect query")
	}
	defer stmt.Close()
	_, err = stmt.Exec(&event.ID)
	return err
}
