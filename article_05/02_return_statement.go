// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 2:
//    Příkaz return.

package main

func f1() {
	println("f1")
	return
}

func main() {
	println("Hello world!")
	f1()
}
