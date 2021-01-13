// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 3:
//    Celočíselné datové typy bez znaménka

package main

import "fmt"

func main() {
	var a uint8 = 10
	var b uint16 = 1000
	var c uint32 = 10000
	var d uint32 = 1000000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
