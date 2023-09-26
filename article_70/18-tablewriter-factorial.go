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
	"strconv"

	"github.com/olekukonko/tablewriter"
)

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

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"n", "n!"})

	for n := 0; n <= 20; n++ {
		f := Factorial(int64(n))
		row := []string{strconv.Itoa(n), strconv.FormatInt(f, 10)}
		table.Append(row)
	}

	table.Render()
}
