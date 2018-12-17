// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 4:
//    Příkaz return ve funkci s návratovou hodnotou.

package main

func f2() int {
	println("f2() před příkazem return")
	return 42
	println("f2() po příkazu return")
}

func main() {
	println("Hello world!")
	println(f2())
}
