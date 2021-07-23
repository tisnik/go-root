// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	MinWidth         = 0
	TabWidth         = 0
	Padding          = 1
	PaddingCharacter = ' '
	Flags            = 0
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, MinWidth, TabWidth, Padding, PaddingCharacter, Flags)
	fmt.Fprintln(w, "1\t2\t3")
	fmt.Fprintln(w, "4\t5\t6")
	fmt.Fprintln(w, "7\t8\t9")
	w.Flush()

	fmt.Println()

	w = tabwriter.NewWriter(os.Stdout, MinWidth, TabWidth, Padding, PaddingCharacter, Flags)
	fmt.Fprintln(w, "foo\tbar\tbaz")
	fmt.Fprintln(w, "foo\tbar\tbaz")
	fmt.Fprintln(w, "foo\tbar\tbaz")
	w.Flush()
}
