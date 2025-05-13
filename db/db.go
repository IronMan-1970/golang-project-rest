package db

import (
	"database/sql"
	"fmt"

	 _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	user := "gingonicuser"
    password := "gingonicuser"
    dbName := "eventsdb"
    remoteHost := "132.145.30.190"  // публічна IP-адреса MySQL
    remotePort := 3306

	var err error
	//Data Source=132.145.30.190,22;Network Library=DBMSSOCN;Initial Catalog=eventsdb;User ID=gingonicuser;Password=gingonicuser;
	// Use in-memory SQLite database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, remoteHost, remotePort, dbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		panic("Could not connect to database.")
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		panic("Could not ping database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createUserTableComand :=
		`CREATE TABLE IF NOT EXISTS users(
			id       INT AUTO_INCREMENT PRIMARY KEY,
			email    VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL
			);`
	_, err := DB.Exec(createUserTableComand)
	if err != nil {
		fmt.Printf("Error creating users table: %v\n", err)
		panic("could not create users table")
	}

	createEventTableComand :=
		`CREATE TABLE IF NOT EXISTS events (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    location    TEXT NOT NULL,
    dateTime    DATETIME,
    user_id     INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);`
	_, err = DB.Exec(createEventTableComand)
	if err != nil {
		fmt.Printf("Error creating events table: %v\n", err)
		panic("could not create events table")
	}

	createRegistrationTable :=
	`
	CREATE TABLE IF NOT EXISTS registrations(
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  event_id INTEGER,
	  user_id INTEGER,
	  FORIGEN KEY (event_id) REFERENCES event(id),
	  FORIGEN KEY (user_id) REFERENCES user(id)
	)
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		fmt.Printf("error occured during creation of registration table: %v\n", err)
	}
}
