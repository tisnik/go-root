package main

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/lib/pq"           // PostgreSQL database driver
	_ "github.com/mattn/go-sqlite3" // SQLite database driver

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Datová struktura s konfigurací připojení k databázi
type StorageConfiguration struct {
	Driver           string `mapstructure:"db_driver" toml:"db_driver"`
	SQLiteDataSource string `mapstructure:"sqlite_datasource" toml:"sqlite_datasource"`
	PGUsername       string `mapstructure:"pg_username" toml:"pg_username"`
	PGPassword       string `mapstructure:"pg_password" toml:"pg_password"`
	PGHost           string `mapstructure:"pg_host" toml:"pg_host"`
	PGPort           int    `mapstructure:"pg_port" toml:"pg_port"`
	PGDBName         string `mapstructure:"pg_db_name" toml:"pg_db_name"`
	PGParams         string `mapstructure:"pg_params" toml:"pg_params"`
}

// Chybové zprávy
const (
	canNotConnectToDataStorageMessage = "Can not connect to data storage"
	connectionToDBNotEstablished      = "Connection to database not established"
	unableToCloseDBRowsHandle         = "Unable to close the DB rows handle"
	databaseOperationFailed           = "Database operation failed"
)

// Inicializace připojení k databázi
func initDatabaseConnection(configuration StorageConfiguration) (*sql.DB, error) {
	driverName := configuration.Driver
	dataSource := ""
	log.Info().Str("driverName", configuration.Driver).Msg("DB connection configuration")

	// inicializace připojení s vybraným driverem
	switch driverName {
	case "sqlite3":
		//driverType := DBDriverSQLite3
		//driver = &sqlite3.SQLiteDriver{}
		dataSource = configuration.SQLiteDataSource
	case "postgres":
		//driverType := DBDriverPostgres
		//driver = &pq.Driver{}
		dataSource = fmt.Sprintf(
			"postgresql://%v:%v@%v:%v/%v?%v",
			configuration.PGUsername,
			configuration.PGPassword,
			configuration.PGHost,
			configuration.PGPort,
			configuration.PGDBName,
			configuration.PGParams,
		)
	default:
		// neznámý driver
		err := fmt.Errorf("driver %v is not supported", driverName)
		log.Err(err).Msg(canNotConnectToDataStorageMessage)
		return nil, err
	}

	// pokus o inicializaci připojení k databázi
	connection, err := sql.Open(driverName, dataSource)

	// test, zda bylo připojení k databázi úspěšné
	if err != nil {
		log.Err(err).Msg(canNotConnectToDataStorageMessage)
		return nil, err
	}

	return connection, nil
}

// Zobrazení všech záznamů v tabulce "persons"
func displayAllRecords(connection *sql.DB) error {
	// dotaz do databáze
	query := "SELECT id, name, surname FROM persons"
	rows, err := connection.Query(query)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		return err
	}

	defer func() {
		// pokud dojde k chybě nebo na konci smyčky, musíme uvolnit prostředky
		if closeErr := rows.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg(unableToCloseDBRowsHandle)
		}
	}()

	// projít všemi vrácenými řádky
	for rows.Next() {
		var (
			id      int
			name    string
			surname string
		)

		// přečtení dat z jednoho vráceného řádku
		if err := rows.Scan(&id, &name, &surname); err != nil {
			return err
		}

		// výpis načteného záznamu
		log.Info().Int("ID", id).
			Str("name", name).
			Str("surname", surname).
			Msg("Record")
	}

	return nil
}

// datová struktura odpovídající struktuře záznamu v databázi
type Record struct {
	Id      int
	Name    string
	Surname string
}

// funkce vracející data přečtená z databázové tabulky
func readAllRecords(connection *sql.DB) ([]Record, error) {
	results := make([]Record, 0)

	// dotaz do databáze
	query := "SELECT id, name, surname FROM persons"
	rows, err := connection.Query(query)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		return results, err
	}

	defer func() {
		// pokud dojde k chybě nebo na konci smyčky, musíme uvolnit prostředky
		if closeErr := rows.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg(unableToCloseDBRowsHandle)
		}
	}()

	// projít všemi vrácenými řádky
	for rows.Next() {
		var record Record

		// přečtení dat z jednoho vráceného řádku
		if err := rows.Scan(&record.Id, &record.Name, &record.Surname); err != nil {
			return results, err
		}

		results = append(results, record)
	}

	return results, nil
}

// funkce vracející data přečtená z databázové tabulky
func readRecordsWithName(connection *sql.DB, name string) ([]Record, error) {
	results := make([]Record, 0)

	// dotaz do databáze
	query := "SELECT id, name, surname FROM persons WHERE name=$1"
	rows, err := connection.Query(query, name)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		return results, err
	}

	defer func() {
		// pokud dojde k chybě nebo na konci smyčky, musíme uvolnit prostředky
		if closeErr := rows.Close(); closeErr != nil {
			log.Error().Err(closeErr).Msg(unableToCloseDBRowsHandle)
		}
	}()

	// projít všemi vrácenými řádky
	for rows.Next() {
		var record Record

		// přečtení dat z jednoho vráceného řádku
		if err := rows.Scan(&record.Id, &record.Name, &record.Surname); err != nil {
			return results, err
		}

		results = append(results, record)
	}

	return results, nil
}

// Vložení nového záznamu do tabulky "persons"
func insertRecord(connection *sql.DB, name string, surname string) (int, error) {
	// provedení SQL příkazu se dvěma parametry
	sqlStatement := "INSERT INTO persons (name, surname) VALUES($1, $2);"
	result, err := connection.Exec(sqlStatement, name, surname)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		return 0, err
	}

	// přečíst počet řádků v tabulce, které byly SQL příkazem upraveny
	affected, err := result.RowsAffected()

	// i tato operace může teoreticky skončit s chybou nebo nemusí být podporována
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// Vymazání záznamu nebo záznamů na základě zapsaného jména
func deleteByName(connection *sql.DB, name string) (int, error) {
	// provedení SQL příkazu s jedním parametrem
	sqlStatement := "DELETE FROM persons WHERE name = $1;"
	result, err := connection.Exec(sqlStatement, name)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		return 0, err
	}

	// přečíst počet řádků v tabulce, které byly SQL příkazem upraveny
	affected, err := result.RowsAffected()

	// i tato operace může teoreticky skončit s chybou nebo nemusí být podporována
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// Vymazání záznamu nebo záznamů na základě zapsaného jména
func deleteByNameInTransaction(connection *sql.DB, name string) (int, error) {
	// nová transakce
	tx, err := connection.Begin()

	// provedení SQL příkazu s jedním parametrem
	sqlStatement := "DELETE FROM persons WHERE name = $1;"
	result, err := tx.Exec(sqlStatement, name)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			log.Err(rollbackError).Msgf("error when trying to rollback a transaction")
		}
		return 0, err
	}

	// přečíst počet řádků v tabulce, které byly SQL příkazem upraveny
	affected, err := result.RowsAffected()

	// i tato operace může teoreticky skončit s chybou nebo nemusí být podporována
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			log.Err(rollbackError).Msgf("error when trying to rollback a transaction")
		}
		return 0, err
	}
	commitError := tx.Commit()
	if commitError != nil {
		log.Err(commitError).Msgf("error when trying to commit a transaction")
	}
	return int(affected), nil
}

// Vymazání záznamu nebo záznamů na základě zapsaného jména
func deleteByNameInTransactionNoCommit(connection *sql.DB, name string) (int, error) {
	// nová transakce
	tx, err := connection.Begin()

	// provedení SQL příkazu s jedním parametrem
	sqlStatement := "DELETE FROM persons WHERE name = $1;"
	result, err := tx.Exec(sqlStatement, name)

	// test, zda byl SQL příkaz proveden bez chyby
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			log.Err(rollbackError).Msgf("error when trying to rollback a transaction")
		}
		return 0, err
	}

	// přečíst počet řádků v tabulce, které byly SQL příkazem upraveny
	affected, err := result.RowsAffected()

	// i tato operace může teoreticky skončit s chybou nebo nemusí být podporována
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			log.Err(rollbackError).Msgf("error when trying to rollback a transaction")
		}
		return 0, err
	}
	return int(affected), nil
}

func main() {
	// nastavit logovací systém pro barevný výstup na terminál
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// konfigurace připojení do databáze
	config := StorageConfiguration{
		Driver:     "postgres",
		PGUsername: "postgres",
		PGPassword: "postgres",
		PGHost:     "localhost",
		PGPort:     5432,
		PGDBName:   "testdb",
		PGParams:   "sslmode=disable",
	}

	log.Debug().Msg("Started")

	// inicializace připojení k databázi
	connection, err := initDatabaseConnection(config)
	if err != nil {
		log.Err(err).Msg(connectionToDBNotEstablished)
		return
	}

	log.Info().Msg("Read all records")
	// přečtení všech záznamů
	results, err := readAllRecords(connection)
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}
	// výpis získaných záznamů
	for _, result := range results {
		log.Info().Int("ID", result.Id).
			Str("name", result.Name).
			Str("surname", result.Surname).
			Msg("Record")
	}

	log.Info().Msg("Display all records")
	// přečtení všech záznamů z tabulky "persons"
	err = displayAllRecords(connection)
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}

	// vymazání záznamu či záznamů na základě zapsaného jména
	affected, err := deleteByName(connection, "Eda")
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}
	log.Info().Int("deleted rows", affected).Msg("DELETE")

	// přidání nového záznamu do databáze
	affected, err = insertRecord(connection, "Eda", "Vodopád")
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}
	log.Info().Int("inserted rows", affected).Msg("INSERT")

	// přečtení všech záznamů z tabulky "persons"
	err = displayAllRecords(connection)
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}

	log.Info().Msg("Read records with name Přemysl")
	// přečtení záznamů
	results, err = readRecordsWithName(connection, "Přemysl")
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}
	// výpis získaných záznamů
	for _, result := range results {
		log.Info().Int("ID", result.Id).
			Str("name", result.Name).
			Str("surname", result.Surname).
			Msg("Record")
	}

	// vymazání záznamu či záznamů na základě zapsaného jména
	affected, err = deleteByNameInTransaction(connection, "Přemysl")
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}
	log.Info().Int("deleted rows", affected).Msg("DELETE")

	// přečtení všech záznamů z tabulky "persons"
	err = displayAllRecords(connection)
	if err != nil {
		log.Err(err).Msg(databaseOperationFailed)
		return
	}

	log.Debug().Msg("Finished")
}
