// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Třicátá devátá část
//    Programovací jazyk Go a relační databáze
//    https://www.root.cz/clanky/programovaci-jazyk-go-a-relacni-databaze/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů ze třicáté deváté části:
//    https://github.com/tisnik/go-root/blob/master/article_39/README.md

package main

import (
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

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
}
