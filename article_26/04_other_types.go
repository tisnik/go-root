// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Dvacátá šestá část
//    Problematika nulových hodnot v Go, aneb proč nil != nil
//    https://www.root.cz/clanky/problematika-nulovych-hodnot-v-go-aneb-proc-nil-nil/
//
// Demonstrační příklad číslo 4:
//    Ostatní datové typy

package main

func main() {
	var p *int
	var s []int
	var m map[string]int
	var c chan int
	var f func()
	var i interface{}

	println(p)
	println(s)
	println(m)
	println(c)
	println(f)
	println(i)
}
