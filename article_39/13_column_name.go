// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá devátá část
//    Programovací jazyk Go a relační databáze
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-relacni-databaze/
//
// Demonstrační příklad číslo 13:
//    Specifikace jména sloupce v databázové tabulce.
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_39/README.md
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/html-root/article_39/13_column_name.html
//

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	Id      int `gorm:"Column:ID"`
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

	var customers []Customer
	db.Find(&customers)
	for _, customer := range customers {
		fmt.Printf("%2d %-10s %-10s %-12s %-12s %s\n", customer.Id, customer.Name, customer.Surname, customer.Address, customer.Country, customer.Phone)
	}
}
