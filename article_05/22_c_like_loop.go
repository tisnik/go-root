// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 22:
//    Programová smyčka for odvozená od C.

package main

func main() {
	var i int
	for i = 0; i < 10; i++ {
		println(i)
	}
	println()
	println(i)
}
