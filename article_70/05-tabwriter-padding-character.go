package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	MinWidth         = 0
	TabWidth         = 0
	PaddingCharacter = '.'
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
