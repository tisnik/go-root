package main_test

import (
	"errors"
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

func TestSelectOnError1(t *testing.T) {
	mockedError := errors.New("mocked error")

	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnError(mockedError)

	err = main.DisplayAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	if err != mockedError {
		t.Errorf("different error was returned: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectOnError2(t *testing.T) {
	mockedError := errors.New("mocked error")

	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnError(mockedError)

	results, err := main.ReadAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	if err != mockedError {
		t.Errorf("different error was returned: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(results) != 0 {
		t.Errorf("different number of results read from database: %d instead of 0", len(results))
		return
	}
}

func TestSelectOnError3(t *testing.T) {
	mockedError := errors.New("mocked error")

	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectQuery("SELECT id, name, surname FROM persons WHERE name=.*").WillReturnError(mockedError)

	results, err := main.ReadRecordsWithName(connection, "Eda")

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	if err != mockedError {
		t.Errorf("different error was returned: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(results) != 0 {
		t.Errorf("different number of results read from database: %d instead of 0", len(results))
		return
	}
}

func TestSelectScanError1(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	err = main.DisplayAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectScanError2(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	results, err := main.ReadAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// první řádek by měl být přečten
	if len(results) != 1 {
		t.Errorf("different number of results read from database: %d instead of 1", len(results))
		return
	}
}

func TestSelectScanError3(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	results, err := main.ReadRecordsWithName(connection, "foo")

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// první řádek by měl být přečten
	if len(results) != 1 {
		t.Errorf("different number of results read from database: %d instead of 1", len(results))
		return
	}
}

func TestSelectCloseError1(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.CloseError(errors.New("close error"))
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	err = main.DisplayAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectCloseError2(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.CloseError(errors.New("close error"))
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons").WillReturnRows(rows)

	_, err = main.ReadAllRecords(connection)

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSelectCloseError3(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "surname"})
	rows.AddRow(1, "foo", "bar")
	rows.CloseError(errors.New("close error"))
	rows.AddRow("this is not integer", "x", "y")
	rows.AddRow(3, "a", "b")

	mock.ExpectQuery("SELECT id, name, surname FROM persons .*").WillReturnRows(rows)

	_, err = main.ReadRecordsWithName(connection, "foobar")

	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertion1(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectExec("INSERT INTO persons (name, surname) VALUES($1, $2);").WithArgs("foo", "bar").WillReturnResult(sqlmock.NewResult(1, 1))

	updated, err := main.InsertRecord(connection, "foo", "bar")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if updated != 1 {
		t.Errorf("one row should be updated, but %d rows were updated", updated)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertion2(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectExec("INSERT INTO persons \\(name, surname\\) VALUES\\(\\$1, \\$2\\);").WithArgs("foo", "bar").WillReturnResult(sqlmock.NewResult(1, 1))

	updated, err := main.InsertRecord(connection, "foo", "bar")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if updated != 1 {
		t.Errorf("one row should be updated, but %d rows were updated", updated)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertionOnError(t *testing.T) {
	mockedError := errors.New("mocked error")

	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectExec("INSERT INTO persons \\(name, surname\\) VALUES\\(\\$1, \\$2\\);").WithArgs("foo", "bar").WillReturnError(mockedError)

	_, err = main.InsertRecord(connection, "foo", "bar")
	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	if err != mockedError {
		t.Errorf("different error was returned: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletion1(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectExec("DELETE FROM persons WHERE name = \\$1;").WithArgs("foo").WillReturnResult(sqlmock.NewResult(1, 1))

	updated, err := main.DeleteByName(connection, "foo")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if updated != 1 {
		t.Errorf("one row should be updated, but %d rows were updated", updated)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletionOnError(t *testing.T) {
	mockedError := errors.New("mocked error")

	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectExec("DELETE FROM persons WHERE name = \\$1;").WithArgs("foo").WillReturnError(mockedError)

	_, err = main.DeleteByName(connection, "foo")
	if err == nil {
		t.Fatalf("error was expected while updating stats")
	}

	if err != mockedError {
		t.Errorf("different error was returned: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletionInTransaction(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM persons WHERE name = \\$1;").WithArgs("foo").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updated, err := main.DeleteByNameInTransaction(connection, "foo")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if updated != 1 {
		t.Errorf("one row should be updated, but %d rows were updated", updated)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletionInTransactionNoCommit(t *testing.T) {
	connection, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer connection.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM persons WHERE name = \\$1;").WithArgs("foo").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updated, err := main.DeleteByNameInTransactionNoCommit(connection, "foo")
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if updated != 1 {
		t.Errorf("one row should be updated, but %d rows were updated", updated)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
