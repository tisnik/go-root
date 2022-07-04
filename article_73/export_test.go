// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá třetí část
//    Jazyk Go prakticky: jednotkové testy kódu, který přistupuje k SQL databázím (dokončení)
//    https://www.root.cz/clanky/jazyk-go-prakticky-jednotkove-testy-kodu-ktery-pristupuje-k-sql-databazim-dokonceni/

package main

var (
	DisplayAllRecords                 = displayAllRecords
	ReadAllRecords                    = readAllRecords
	ReadRecordsWithName               = readRecordsWithName
	InsertRecord                      = insertRecord
	DeleteByName                      = deleteByName
	DeleteByNameInTransaction         = deleteByNameInTransaction
	DeleteByNameInTransactionNoCommit = deleteByNameInTransactionNoCommit
)
