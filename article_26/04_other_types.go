// Seriál "Programovací jazyk Go"
//
// Dvacátá šestá část
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
