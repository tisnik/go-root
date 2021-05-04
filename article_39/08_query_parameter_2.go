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

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float32
}

func productsWithPriceBetween(connections *sql.DB, min float32, max float32) ([]Product, error) {
	products := []Product{}

	rows, err := connections.Query("SELECT id, name, description, price FROM products WHERE price between ? and ? ORDER BY id", min, max)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product

		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return products, err
	}
	return products, nil
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

	products, err := productsWithPriceBetween(connections, 0.0, 2.0)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		fmt.Printf("%2d %-20s %-10s %f\n", product.Id, product.Name, product.Description, product.Price)
	}
}
