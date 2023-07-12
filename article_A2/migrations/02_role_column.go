package main

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users ADD COLUMN role VARCHAR;")
	if err != nil {
		return err
	}
	return nil
}

func Down00002(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER TABLE users DROP COLUMN role;")
	if err != nil {
		return err
	}
	return nil
}
