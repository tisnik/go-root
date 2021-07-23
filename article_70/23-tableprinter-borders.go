// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/

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
	table.BorderTop, table.BorderBottom, table.BorderLeft, table.BorderRight = true, true, true, true
	table.CenterSeparator = "│"
	table.ColumnSeparator = "│"
	table.RowSeparator = "─"

	table.Print(data)
}
