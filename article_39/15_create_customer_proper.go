// Seriál "Programovací jazyk Go"
//
// Třicátá devátá část
//    Programovací jazyk Go a relační databáze
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-relacni-databaze/

package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Customer struct {
	Id      int `gorm:"Column:ID;PRIMARY_KEY"`
	Name    string
	Surname string
	Address string
	Country string
	Phone   string
}

func main() {
	dbDriver := flag.String("dbdriver", "sqlite3", "database driver specification")
	storageSpecification := flag.String("storage", "./test.db", "storage specification")
	flag.Parse()

	db, err := gorm.Open(*dbDriver, *storageSpecification)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Customer{})

	customer := Customer{Name: "Franta", Surname: "Vomáčka", Address: "Horní dolní", Country: "CR", Phone: "603 123 456"}
	db.Create(&customer)

	var customers []Customer
	db.Find(&customers)
	for _, customer := range customers {
		fmt.Printf("%2d %-10s %-10s %-12s %-12s %s\n", customer.Id, customer.Name, customer.Surname, customer.Address, customer.Country, customer.Phone)
	}
}
