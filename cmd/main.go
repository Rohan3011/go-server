package main

import (
	"fmt"
	"log"

	"github.com/rohan3011/go-server/cmd/api"
	"github.com/rohan3011/go-server/config"
	"github.com/rohan3011/go-server/db"
)

func main() {

	db := db.NewSQLiteDB(config.Env.DBConnStr)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Env.Port), db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
