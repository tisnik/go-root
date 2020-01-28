// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
//
// Demonstrační příklad číslo 5:
//    Příkaz return ve funkci s návratovou hodnotou.

package main

func f2() int {
	println("f2() před příkazem return")
	return 42
	println("f2() po příkazu return")
	return -1
}

func main() {
	println("Hello world!")
	println(f2())
}
