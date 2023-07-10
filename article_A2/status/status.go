package main

import (
	"log"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const databaseType = "sqlite"

const databaseFile = "./test.db"

const command = "status"

const migrationScriptsDirectory = "./"

func main() {
	db, err := goose.OpenDBWithDriver(databaseType, databaseFile)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}

	err = goose.Run(command, db, migrationScriptsDirectory, arguments...)
	if err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
