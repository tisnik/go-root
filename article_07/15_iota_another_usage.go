// Seriál "Programovací jazyk Go"
//
// Sedmá část
//    Programovací jazyk Go: dokončení popisu vlastností samotného jazyka
//    https://www.root.cz/clanky/programovaci-jazyk-go-dokonceni-popisu-vlastnosti-samotneho-jazyka/
//
// Demonstrační příklad číslo 15:
//    Konstanty a identifikátor iota.

package main

import "fmt"

const (
	a = iota + 0.5i
	b = iota + 0.5i
	c = 1.0 / iota
	d = 1 << iota
	e = iota ^ 0xff
)

func main() {
	fmt.Printf("%f\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%f\n", c)
	fmt.Printf("%d\n", d)
	fmt.Printf("%d\n", e)
}
