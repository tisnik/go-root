// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 11
//    Datové typy reprezentující hodnoty s plovoucí řádovou čárkou

package main

import "fmt"

func main() {
	var a float32 = -1.5
	var b float32 = 1.5
	var c float32 = 1e30
	var d float32 = 1e-30

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e float64 = -1.5
	var f float64 = 1.5
	var g float64 = 1e300
	var h float64 = 1e-300

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
