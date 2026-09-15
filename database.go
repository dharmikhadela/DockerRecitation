package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // Imports the driver that allows us to use Postgres with Go
)

var db *sql.DB // Declare a pointer to a sql db

func InitDatabase() {
	dbHostname := "postgres"
	dbName := os.Getenv("LOCAL_DB")
	dbUser := os.Getenv("LOCAL_DB_USER")
	dbPassword := os.Getenv("LOCAL_DB_PASSWORD")

	conn_string := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHostname + ":5432/" + dbName
	println(conn_string)
	local_db, err := sql.Open(
		// Driver
		"pgx",
		// Connection string -- "postgres://username:password@hostname:5432/dbname"
		conn_string,
	)

	if err != nil {
		fmt.Println("Error connecting to db")
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection Successful!")
	}

	db = local_db

	_, err2 := db.Exec(`
	CREATE TABLE IF NOT EXISTS greetings (
		id SERIAL PRIMARY KEY,
    	message TEXT NOT NULL)
	`)

	if err2 != nil {
		fmt.Println("Error creating table")
		fmt.Println(err2)
		return
	}

	_, err3 := db.Exec(`
	INSERT INTO greetings (message) VALUES ('Hello from PostgreSQL!');
	`)

	if err3 != nil {
		fmt.Println("Error inserting in db")
		fmt.Println(err)
		return
	}
}

func CloseDB() {
	db.Close()
}
