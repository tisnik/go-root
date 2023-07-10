package main

import (
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const databaseType = "sqlite"

const databaseFile = "./test.db"

const migrationScriptsDirectory = "./"

func main() {
	args := os.Args

	if len(args) <= 1 {
		log.Fatalf("command is expected")
		return
	}

	command := args[1]

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

	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	err = goose.Run(command, db, migrationScriptsDirectory, arguments...)
	if err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
