// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Druhá část:
//    Datové typy v programovacím jazyku Go
//    https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/
//
// Demonstrační příklad číslo 4:
//    Celočíselné konstanty reprezentované v různých číselných soustavách
//
// Dokumentace ve stylu "literate programming":
//    https://tisnik.github.io/go-root/article_02/04_integer_constants.html

package main

import "fmt"

func main() {
	var a uint8 = 10
	var b uint8 = 010
	var c uint8 = 0x10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d int8 = -10
	var e int8 = -010
	var f int8 = -0x10

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
