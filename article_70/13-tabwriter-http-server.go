// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Sedmdesátá část
//    Vykreslení tabulek do terminálu v jazyce Go
//    https://www.root.cz/clanky/vykresleni-tabulek-do-terminalu-v-jazyce-go/

package main

import (
	"fmt"
	"net/http"
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

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	w := tabwriter.NewWriter(writer, MinWidth, TabWidth, Padding, PaddingCharacter, Flags)
	fmt.Fprintln(w, "n\tn!\t")

	for n := 0; n <= 20; n++ {
		fmt.Fprintf(w, "%d\t", n)
		result := Factorial(int64(n))
		fmt.Fprintf(w, "%d\t", result)
		fmt.Fprintln(w)
	}

	w.Flush()
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8000", nil)
}
