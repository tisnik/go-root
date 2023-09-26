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
	MinWidth         = 0
	TabWidth         = 0
	PaddingCharacter = ' '
	Flags            = 0
)

func main() {
	for padding := 0; padding < 10; padding++ {
		fmt.Printf("padding = %d\n", padding)
		w := tabwriter.NewWriter(os.Stdout, MinWidth, TabWidth, padding, PaddingCharacter, Flags)
		fmt.Fprintln(w, "1\t1\t1")
		fmt.Fprintln(w, "22\t22\t22")
		fmt.Fprintln(w, "333\t333\t333")
		fmt.Fprintln(w, "4444\t4444\t4444")
		w.Flush()
		fmt.Println()
	}
}
