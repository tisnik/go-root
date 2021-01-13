// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 19
//    Lokální proměnné

package main

import "fmt"

func main() {
	var i1 int8
	var i2 int32
	var u1 uint8
	var u2 uint32

	var f1 float32
	var f2 float64
	var c1 complex64
	var c2 complex128

	var s string

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(u1)
	fmt.Println(u2)

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println(s)
}
