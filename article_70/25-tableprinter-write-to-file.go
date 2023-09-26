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
	"log"
	"os"

	"github.com/lensesio/tableprinter"
)

type Record struct {
	N int   `header:"n"`
	F int64 `header:"n!"`
}

// Factorial computes factorial for given n that might be positive integer
func Factorial(n int64) int64 {
	switch {
	case n < 0:
		return 1
	case n == 0:
		return 1
	default:
		return n * Factorial(n-1)
	}
}

const MaxN = 17

func main() {
	file, err := os.Create("table3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]Record, 0)

	for n := 0; n <= MaxN; n++ {
		data = append(data, Record{n, Factorial(int64(n))})
	}

	table := tableprinter.New(file)
	table.BorderTop, table.BorderBottom, table.BorderLeft, table.BorderRight = true, true, true, true
	table.CenterSeparator = "│"
	table.ColumnSeparator = "│"
	table.RowSeparator = "─"

	table.Print(data)
}
