// Seriál "Programovací jazyk Go"
//
// Pátá část
//
// Demonstrační příklad číslo 3:
//    Příkaz return ve funkci bez návratové hodnoty

package main

func f1() {
	println("f1() před příkazem return")
	return
	println("f1() po příkazu return")
}

func main() {
	println("Hello world!")
	f1()
}
