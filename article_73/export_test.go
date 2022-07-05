// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá třetí část
//    Jazyk Go prakticky: jednotkové testy kódu, který přistupuje k SQL databázím (dokončení)
//    https://www.root.cz/clanky/jazyk-go-prakticky-jednotkove-testy-kodu-ktery-pristupuje-k-sql-databazim-dokonceni/
//
// Seznam příkladů ze sedmdesáté třetí části:
//    https://github.com/tisnik/go-root/blob/master/article_73/README.md

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
