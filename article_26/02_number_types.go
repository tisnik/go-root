// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Repositář:
//    https://github.com/tisnik/go-root/
//
// Seznam demonstračních příkladů z dvacáté šesté části:
//    https://github.com/tisnik/go-root/blob/master/article_26/README.md
//
// Demonstrační příklad číslo 2:
//    Numerické datové typy

package main

func main() {
	var i1 int8
	var i2 int32
	var u1 uint8
	var u2 uint32

	var f1 float32
	var f2 float64
	var c1 complex64
	var c2 complex128

	println(i1)
	println(i2)
	println(u1)
	println(u2)

	println(f1)
	println(f2)
	println(c1)
	println(c2)
}
