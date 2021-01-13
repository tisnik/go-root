// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 1:
//    Celočíselné datové typy se znaménkem

package main

import "fmt"

func main() {
	var a int8 = -10
	var b int16 = -1000
	var c int32 = -10000
	var d int32 = -1000000

	var r1 rune = 'a'
	var r2 rune = '\x40'
	var r3 rune = '\n'
	var r4 rune = '\u03BB'

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}
