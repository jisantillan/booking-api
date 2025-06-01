package model

import (
	"booking-api/db"
	"log"
	"time"
)

type Event struct {
	ID          int64     `binding: "required"`
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

func Save(event Event) Event {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	// Use Exec when the query changes db status
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	event.ID = id

	return event
}

// Use Exec when the query dont change db status
func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, datetime, user_id FROM events`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() //"Defer: Cuando termine esta función, cerrá esto automáticamente."

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return events, nil
}
