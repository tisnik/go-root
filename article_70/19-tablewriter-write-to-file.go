// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/

package main

import (
	"log"
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
	file, err := os.Create("table2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	table := tablewriter.NewWriter(file)
	table.SetHeader([]string{"n", "n!"})

	for n := 0; n <= 20; n++ {
		f := Factorial(int64(n))
		row := []string{strconv.Itoa(n), strconv.FormatInt(f, 10)}
		table.Append(row)
	}

	table.Render()
}
