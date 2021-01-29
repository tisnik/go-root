// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
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
