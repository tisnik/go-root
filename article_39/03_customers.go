// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá devátá část
//    Programovací jazyk Go a relační databáze
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-relacni-databaze/

package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func listOfCustomers(connections *sql.DB) {
	rows, err := connections.Query("SELECT id, name, surname, address, country, phone FROM customers ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var surname string
		var address string
		var country string
		var phone string

		if err := rows.Scan(&id, &name, &surname, &address, &country, &phone); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%2d %-10s %-10s %-12s %-12s %s\n", id, name, surname, address, country, phone)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbDriver := flag.String("dbdriver", "sqlite3", "database driver specification")
	storageSpecification := flag.String("storage", "./test.db", "storage specification")
	flag.Parse()

	connections, err := sql.Open(*dbDriver, *storageSpecification)
	if err != nil {
		log.Fatal("Can not connect to data storage", err)
	}
	defer connections.Close()

	log.Printf("Connected to database %v", connections)
	listOfCustomers(connections)
}
