package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// checks the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to postgres database!")

	return db
}
