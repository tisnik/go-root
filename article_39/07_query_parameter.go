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

type Customer struct {
	Id      int
	Name    string
	Surname string
	Address string
	Country string
	Phone   string
}

func customersFromCountry(connections *sql.DB, country string) ([]Customer, error) {
	customers := []Customer{}

	rows, err := connections.Query("SELECT id, name, surname, address, country, phone FROM customers WHERE country=? ORDER BY id", country)
	if err != nil {
		return customers, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer

		if err := rows.Scan(&customer.Id, &customer.Name, &customer.Surname, &customer.Address, &customer.Country, &customer.Phone); err != nil {
			return customers, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return customers, err
	}
	return customers, nil
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

	customers, err := customersFromCountry(connections, "Mexico")
	if err != nil {
		log.Fatal(err)
	}

	for _, customer := range customers {
		fmt.Printf("%2d %-10s %-10s %-12s %-12s %s\n", customer.Id, customer.Name, customer.Surname, customer.Address, customer.Country, customer.Phone)
	}
}
