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
	fmt.Fprintln(w, "1\t1\t1")
	fmt.Fprintln(w, "22\t22\t22")
	fmt.Fprintln(w, "333\t333\t333")
	fmt.Fprintln(w, "4444\t4444\t4444")
	w.Flush()

	fmt.Println()

	w = tabwriter.NewWriter(os.Stdout, MinWidth, TabWidth, Padding, PaddingCharacter, Flags)
	fmt.Fprintln(w, "f\tb\tb")
	fmt.Fprintln(w, "foo\tbar\tbaz")
	fmt.Fprintln(w, "foobar\tbarbaz\tbazfoo")
	w.Flush()
}
