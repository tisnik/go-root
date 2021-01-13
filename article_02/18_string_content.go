// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 18
//    Datový typ řetězec a jeho obsah

package main

import "fmt"

func main() {
	var s string = "Hello\nworld!\nžluťoučký kůň"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%02x ", s[i])
	}
}
