package main

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	_, err := tx.Exec(`
CREATE TABLE users (
    id      INTEGER NOT NULL,
    name    VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    PRIMARY KEY (id)
);
        `)
	if err != nil {
		return err
	}
	return nil
}

func Down00001(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
