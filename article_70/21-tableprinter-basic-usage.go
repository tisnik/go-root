// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam příkladů ze sedmdesáté části:
//    https://github.com/tisnik/go-root/blob/master/article_70/README.md

package main

import (
	"os"

	"github.com/lensesio/tableprinter"
)

type Record struct {
	Rank  string `header:"rank"`
	Title string `header:"title"`
	Value int    `header:"value"`
}

func main() {
	data := []Record{
		{"A", "The Good", 500},
		{"B", "The Very very Bad Man", 288},
		{"C", "The Ugly", 120},
		{"D", "The Gopher", 800},
	}

	table := tableprinter.New(os.Stdout)

	table.Print(data)
}
