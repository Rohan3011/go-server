package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rohan3011/go-server/config"
	"github.com/rohan3011/go-server/db"
)

func main() {
	db := db.NewPostgresDB(config.Env.DBConnStr)
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
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
		log.Println("Migrate up successful")

	case "down":
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrate down successful")

	default:
		log.Println("invalid command")
	}

}
