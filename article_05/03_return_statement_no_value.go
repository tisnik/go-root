// Seriál "Programovací jazyk Go"
//
// Pátá část
//    Konstrukce pro řízení běhu programu v jazyce Go
//    https://www.root.cz/clanky/konstrukce-pro-rizeni-behu-programu-v-jazyce-go/
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
