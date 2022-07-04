// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá druhá část
//    Jazyk Go prakticky: jednotkové testy kódu, který přistupuje k SQL databázím
//    https://www.root.cz/clanky/jazyk-go-prakticky-jednotkove-testy-kodu-ktery-pristupuje-k-sql-databazim/

package main_test

import (
	"testing"

	main "db-test"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSelect1(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{})

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	err = main.DisplayAllRecords(connection)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelect2(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.AddRow(2, "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	err = main.DisplayAllRecords(connection)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelect3(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.AddRow(2, "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	results, err := main.ReadAllRecords(connection)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(results) != 3 {
		t.Errorf("different number of results read from database: %d instead of 3", len(results))
		return
	}

	expected := main.Record{1, "foo", "bar"}
	if results[0] != expected {
		t.Errorf("first result is different: %+v versus %+v", results[0], expected)
	}
}

func TestSelect4(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(2, "Eda", "Vodopád")

	mock.ExpectQuery("SELECT id, name, surname FROM persons WHERE name=\\$1").WillReturnRows(rows)

	results, err := main.ReadRecordsWithName(connection, "Eda")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(results) != 1 {
		t.Errorf("different number of results read from database: %d instead of 1", len(results))
		return
	}

	expected := main.Record{2, "Eda", "Vodopád"}
	if results[0] != expected {
		t.Errorf("first result is different: %+v versus %+v", results[0], expected)
	}
}

func TestSelect5(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(2, "Eda", "Vodopád")

	mock.ExpectQuery("SELECT id, name, surname FROM persons WHERE name=.*").WillReturnRows(rows)

	results, err := main.ReadRecordsWithName(connection, "Eda")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(results) != 1 {
		t.Errorf("different number of results read from database: %d instead of 1", len(results))
		return
	}

	expected := main.Record{2, "Eda", "Vodopád"}
	if results[0] != expected {
		t.Errorf("first result is different: %+v versus %+v", results[0], expected)
	}
}
