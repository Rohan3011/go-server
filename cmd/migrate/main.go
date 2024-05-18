package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/rohan3011/go-server/config"
	"github.com/rohan3011/go-server/db"
)

func main() {
	db := db.NewSQLiteDB(config.Env.DBConnStr)

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"sqlite3",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err == migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Fatal(err)
		}

	default:
		log.Println("invalid command")
	}

}
