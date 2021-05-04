// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá devátá část
//    Programovací jazyk Go a relační databáze
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-relacni-databaze/

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	connections, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Can not connect to data storage", err)
	}
	defer connections.Close()
	log.Printf("Connected to database %v", connections)
}
