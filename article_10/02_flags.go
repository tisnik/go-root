// Seriál "Programovací jazyk Go"
//
// Desátá část
//
// Demonstrační příklad číslo 2:
//    Vylepšené přečtení argumentů předaných na příkazovém řádku.

package main

import (
	"flag"
	"fmt"
)

func main() {
	var width int
	var height int
	var aa bool
	var output string

	flag.IntVar(&width, "width", 0, "image width")
	flag.IntVar(&height, "height", 0, "image height")
	flag.BoolVar(&aa, "aa", false, "enable antialiasing")
	flag.StringVar(&output, "output", "", "output file name")

	flag.Parse()

	fmt.Printf("width: %d\n", width)
	fmt.Printf("height: %d\n", height)
	fmt.Printf("antialiasing: %t\n", aa)
	fmt.Printf("output file name: %s\n", output)
}
