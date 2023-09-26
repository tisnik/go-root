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
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	MinWidth         = 5
	TabWidth         = 0
	Padding          = 1
	PaddingCharacter = ' '
	Flags            = tabwriter.AlignRight | tabwriter.Debug
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
	w := tabwriter.NewWriter(os.Stdout, MinWidth, TabWidth, Padding, PaddingCharacter, Flags)
	fmt.Fprintln(w, "n\tn!\t")

	for n := 0; n <= 20; n++ {
		fmt.Fprintf(w, "%d\t%d\t\n", n, Factorial(int64(n)))
	}

	w.Flush()
}
