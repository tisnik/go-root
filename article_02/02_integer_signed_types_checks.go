// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 2:
//    Kontrola celočíselných konstant

package main

import "fmt"

func main() {
	var a int8 = -1000
	var b int16 = -100000
	var c int32 = -10000000000
	var d int32 = -10000000000000000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
