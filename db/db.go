package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const connectionString = "postgresql://events_nezf_user:oMWXjqG7u9gtFhajI6S6gUMRGssewjNT@dpg-cvml72je5dus73f7vod0-a.oregon-postgres.render.com/events_nezf"

func InitDB() {
	var err error
	DB, err = sql.Open(
		"postgres", connectionString)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createUserTableCommand := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
`
	_, err := DB.Exec(createUserTableCommand)
	if err != nil {
		panic(fmt.Sprintf("could not create events ", err))
	}

	createEventTableCommand := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP,
        user_id INTEGER
      
    );
`

	_, err = DB.Exec(createEventTableCommand)
	if err != nil {
		panic("could not create events")
	}
}
