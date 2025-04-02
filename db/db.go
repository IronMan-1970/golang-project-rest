package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

const connectionString = "postgresql://events_nezf_user:oMWXjqG7u9gtFhajI6S6gUMRGssewjNT@dpg-cvml72je5dus73f7vod0-a.oregon-postgres.render.com/events_nezf"

func InitDB() {
	var err error
	DB, err = sql.Open("postgre", "postgresql://events_nezf_user:oMWXjqG7u9gtFhajI6S6gUMRGssewjNT@dpg-cvml72je5dus73f7vod0-a.oregon-postgres.render.com/events_nezf")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createUserTableComand := `
    CREATE TABLE IF NOT EXISTS users(
      id        INTEGER PRIMARY KEY AUTOINCREMENT,
      email     TEXT NOT NULL UNIQUE,
      password  TEXT NOT NULL 
    )
  `
	_, err := DB.Exec(createUserTableComand)
	if err != nil {
		panic("could not create events")
	}

	createEventTableComand := `
    CREATE TABLE IF NOT EXISTS events (
      id          INTEGER PRIMARY KEY AUTOINCREMENT,
      name        TEXT NOT NULL,
      description TEXT NOT NULL,
      location    TEXT NOT NULL,
      dateTime    DATETIME,
      user_id     INTEGER,
      FOREIGN KEY(user_id) REFERENCES users(id)
  )
`

	_, err = DB.Exec(createEventTableComand)
	if err != nil {
		panic("could not create events")
	}
}
